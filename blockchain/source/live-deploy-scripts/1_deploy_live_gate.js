let GateMultisig = artifacts.require('./MultiSigWallet.sol');
let GateKeeperLive = artifacts.require('./SimpleGatekeeperWithLimitLive.sol');
let SNMMaster = artifacts.require('./SNM.sol')

let MSOwners = [
    '0xdaec8F2cDf27aD3DF5438E5244aE206c5FcF7fCd',
    '0xd9a43e16e78c86cf7b525c305f8e72723e0fab5e',
    '0x72cb2a9AD34aa126fC02b7d32413725A1B478888',
    '0x1f50Be5cbFBFBF3aBD889e17cb77D31dA2Bd7227',
    '0xe062C67207F7E478a93EF9BEA39535d8EfFAE3cE',
    '0x5fa359a9137cc5ac2a85d701ce2991cab5dcd538',
    '0x7aa5237e0f999a9853a9cc8c56093220142ce48e',
    '0xd43f262536e916a4a807d27080092f190e25d774',
    '0xdd8422eed7fe5f85ea8058d273d3f5c17ef41d1c',
];

let MSRequired = 1;
// let freezingTime = 60 * 15;
let freezingTime = 0;
let SNMMasterchainAddress = '0x983f6d60db79ea8ca4eb9968c6aff8cfa04b3c63';
// let actualGasPrice = 3000000000;
let actualGasPrice = 0;

module.exports = function (deployer, network) {
    deployer.then(async () => { // eslint-disable-line promise/catch-or-return
        if (network === 'master') {
            // 0) deploy `SNM Master`
            await deployer.deploy(SNMMaster, { gasPrice: actualGasPrice });
            let token = await SNMMaster.deployed();

            // 1) deploy `GatekeeperLive` multisig
            await deployer.deploy(GateMultisig, MSOwners, MSRequired, { gasPrice: actualGasPrice });
            let multisig = await GateMultisig.deployed();

            // 2) deploy Live Gatekeeper
            await deployer.deploy(GateKeeperLive, token.address, freezingTime, { gasPrice: actualGasPrice });
            let gk = await GateKeeperLive.deployed();

            // 2.1) add keeper with 100k limit for testing
            await gk.ChangeKeeperLimit('0xAfA5a3b6675024af5C6D56959eF366d6b1FBa0d4', 100000 * 1e18, { gasPrice: actualGasPrice }); // eslint-disable-line max-len
            await gk.ChangeKeeperLimit('0xfa578b05fbd9e1e7c1e69d5add1113240d641bc2', 100000 * 1e18, { gasPrice: actualGasPrice }); // eslint-disable-line max-len
            await gk.ChangeKeeperLimit('0x56c8b9ab7a9594f2d60427fcedbff6ab63c43281', 100000 * 1e18, { gasPrice: actualGasPrice }); // eslint-disable-line max-len

            // 3) transfer Live Gatekeeper ownership to `GatekeeperLive` multisig
            await gk.transferOwnership(multisig.address, { gasPrice: actualGasPrice });
        }
    });
};
