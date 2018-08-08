package main

import (
	"context"
	"fmt"
	"log"
	"os"
    "math/rand"

    "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sonm-io/core/blockchain"
    "github.com/sonm-io/core/proto"
    "github.com/sonm-io/core/util"
)

func getBlockNumber() {
	// client, err := blockchain.NewClient("https://sidechain-dev.sonm.com")
	client, err := blockchain.NewClient("http://172.18.196.12:8545")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	result, _ := client.GetLastBlock(context.TODO())
	fmt.Println(result)
}

func placeOrder() {
	// hexKey := "a5dd45e0810ca83e21f1063e6bf055bd13544398f280701cbfda1346bcf3ae64"
    hexKey := "f72394ce4e4869821edb545c44a0d407267fbff890aa7368d53c7ff3ec309e6c"
	prv, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Fatalln(err)
		return
	}

    fmt.Println("setup api")
	api, err := blockchain.NewAPI(context.Background())
	if err != nil {
		log.Fatalln(err)
		return
	}

    fmt.Println("get balance")
	balance, err := api.SidechainToken().BalanceOf(context.Background(), crypto.PubkeyToAddress(prv.PublicKey))
	if err != nil {
		log.Fatalln(err)
		return
	}

    log.Println("Balance: ", balance)

	allowance, _ := api.SidechainToken().AllowanceOf(context.Background(), crypto.PubkeyToAddress(prv.PublicKey), api.ContractRegistry().MarketAddress())

	if err != nil {
		log.Fatalln(err)
		return
	}

    log.Println("Allowance: ", allowance)

	// price := sonm.NewBigIntFromInt(1)
    price := sonm.NewBigIntFromInt(0)

    log.Println("Price: ", price.Unwrap())

	// if balance.Cmp(price.Unwrap()) < 0 {
	// 	log.Fatalln("lack of balance")
	// 	return
	// }

	// if allowance.Cmp(price.Unwrap()) < 0 {
	// 	log.Fatalln("lack of allowance")
	// 	return
	// }

    // OrderType_ANY OrderType = 0
    // OrderType_BID OrderType = 1
    // OrderType_ASK OrderType = 2
	order := &sonm.Order{
		OrderType:      2, //sonm.OrderType_ASK,
		CounterpartyID: sonm.NewEthAddress(common.HexToAddress("0x0")),
		Duration:       3600 - 50 + (rand.Uint64() % 100),
		Price:          price,
		Netflags:       sonm.NetFlagsFromBoolSlice([]bool{true, true, true}),
		IdentityLevel:  1,
		Blacklist:      "0x0",
		Tag:            []byte("00000"),
		Benchmarks: &sonm.Benchmarks{
			Values: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
	}

	res, err := api.Market().PlaceOrder(
		context.Background(),
		prv,
		order,
	)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	log.Println(res.Id.Unwrap().String())
	ordId, err := util.ParseBigInt(res.Id.Unwrap().String())
	if err != nil {
		log.Fatalln("Cannot cast")
		return
	}
    log.Println("OrderId: ", ordId)
    // log.Println("Canceling")
	// err = api.Market().CancelOrder(context.Background(), prv, ordId)

	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }
	// log.Println("canceled")
}

func main() {
    getBlockNumber()
	placeOrder()
}
