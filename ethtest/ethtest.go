package ethtest

import(
    "context"
    "fmt"
    "log"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
)

type BlockJson struct {
    BlockHash     string    `json:"blockHash"`
    BlockNumber   uint64    `json:"blockNumber"`
    BlockTime     uint64    `json:"blockTime"`
    ParentHash    string    `json:"parentHash"`
    TransHash     []string  `json:"transactions"`

}

func GetEthtransactionHash(keyid string) *types.Transaction {
    client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545/")
    if err != nil {
	fmt.Println("this error is::::",err)
    }
    txHashTku := common.HexToHash(keyid)
    txTku, _, err := client.TransactionByHash(context.Background(), txHashTku)

    if err != nil {
	log.Fatal("txhash error::",err)
    }
    fmt.Println("【Transaction Hash】",txHashTku.Hex())
    fmt.Println("【The transaction tx data string】\n", txTku)

    return txTku
}


func GetEthblocksId(keyid string) *types.Block {
    client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545/")
    if err != nil {
	fmt.Println("this error is::::",err)
    }
    blockHash := common.HexToHash(keyid)
    blockresult, err := client.BlockByHash(context.Background(), blockHash)

    if err != nil {
	log.Fatal("blockhash error::",err)
    }
    var transResult  []string
    for i := 0; i < len(blockresult.Transactions()); i++{
        transResult = append(transResult, blockresult.Transactions()[i].Hash().Hex())
    }
    fmt.Println(len(blockresult.Transactions())) // 7
    result := BlockJson{
	BlockHash: blockresult.Hash().Hex(),
	BlockNumber: blockresult.Number().Uint64(),
	BlockTime: blockresult.Time(),
	ParentHash: blockresult.ParentHash().Hex(),
	TransHash: transResult,
    }
    fmt.Println("result!!!!",result)
    return blockresult
}
