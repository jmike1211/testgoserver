package main

import(
    "github.com/gin-gonic/gin"
    "net/http"
    "testgoserver/ethtest"
)

type TestData struct {
    Hello string `json:"hello"`
}

func blocks(c *gin.Context) {
	limitBlocks := c.Query("limit")
        c.JSON(http.StatusOK, limitBlocks)
}

func blocksId(c *gin.Context) {
	blockId := c.Param("id")
	jsonresult := ethtest.GetEthblocksId(blockId)
	c.JSON(http.StatusOK, gin.H{"return": jsonresult})
}

func transaction(c *gin.Context) {
	transHash := c.Param("txHash")
	jsonresult := ethtest.GetEthtransactionHash(transHash)
	c.JSON(http.StatusOK, gin.H{"return": jsonresult})
}

func main() {
        server := gin.Default()
        server.GET("/blocks", blocks)
	server.GET("/blocks/:id", blocksId)
	server.GET("/transaction/:txHash", transaction)
        server.Run(":8888")
}

