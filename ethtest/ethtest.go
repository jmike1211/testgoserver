package ethtest

import(
    "context"
    "fmt"
    "log"
    "strconv"
    "math/big"
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

type LimitBlocksJson struct {
    BlockHash     string    `json:"blockHash"`
    BlockNumber   uint64    `json:"blockNumber"`
    BlockTime     uint64    `json:"blockTime"`
    ParentHash    string    `json:"parentHash"`
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


func GetEthblocksId(keyid string) BlockJson {
    client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545/")
    if err != nil {
	fmt.Println("this error is::::",err)
    }
    blockNum, _ := strconv.Atoi(keyid)
    bigblockNum := big.NewInt(int64(blockNum))
    blockresult, err := client.BlockByNumber(context.Background(), bigblockNum)

    if err != nil {
	log.Fatal("blockhash error::",err)
    }
    var transResult  []string
    for i := 0; i < len(blockresult.Transactions()); i++{
        transResult = append(transResult, blockresult.Transactions()[i].Hash().Hex())
    }
    result := BlockJson{
	BlockHash: blockresult.Hash().Hex(),
	BlockNumber: blockresult.Number().Uint64(),
	BlockTime: blockresult.Time(),
	ParentHash: blockresult.ParentHash().Hex(),
	TransHash: transResult,
    }
    return result
}

func GetEthblocksLimit(keyid string) []LimitBlocksJson {
    client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545/")
    if err != nil {
	fmt.Println("this error is::::",err)
    }
    blockNum, _ := strconv.Atoi(keyid)
    bigblockNum := uint64(blockNum)
    blockresult, err := client.BlockByNumber(context.Background(),nil)
    if err != nil {
	log.Fatal("blockhash error::",err)
    }
    nowBlockhigh := blockresult.Number().Uint64()
    var blockResult  []LimitBlocksJson
    for i := nowBlockhigh; i > (nowBlockhigh - bigblockNum); i--{
	limitBlockNum := big.NewInt(int64(i))
	limitBlockresult, _ := client.BlockByNumber(context.Background(),limitBlockNum)
	blockResult = append(blockResult, LimitBlocksJson{
            BlockHash: limitBlockresult.Hash().Hex(),
            BlockNumber: limitBlockresult.Number().Uint64(),
            BlockTime: limitBlockresult.Time(),
            ParentHash: limitBlockresult.ParentHash().Hex(),
        })
    }
    return blockResult
}
