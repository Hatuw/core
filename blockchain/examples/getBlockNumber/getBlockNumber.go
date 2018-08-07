package main

import (
	"context"
	"log"
	"os"
	"fmt"

	"github.com/sonm-io/core/blockchain"
)

func main() {
	// client, err := blockchain.NewClient("https://sidechain-dev.sonm.com")
	client, err := blockchain.NewClient("http://172.18.196.12")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	client.GetLastBlock(context.TODO())
	fmt.Println(context.TODO())
	fmt.Println(client)
}
