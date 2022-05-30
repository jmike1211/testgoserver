# testgoserver
go run ethapi.go

## using API
/blocks?limit=n

/blocks/:id

/transaction/:txHash

##eth rpc
go-ethereum/ethclient
client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545/")

