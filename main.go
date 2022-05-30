package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Transacao struct {
	Id            int       `json:"id"`
	CodTransacao  int       `json:"codTransaction"`
	Moeda         string    `json:"currency"`
	Emissor       string    `json:"issuer"`
	Receptor      string    `json:"receiver"`
	DataTransacao time.Time `json:"dateTransaction"`
}

func main() {

	r := gin.Default()
	//r.GET("/greetings", greetings)
	r.GET("/transactions", getAll)
	r.Run()
}

func getAll(c *gin.Context) {

	//transacaos := make([]Transacao, 2)
	transacaos := []Transacao{
		{Id: 1, CodTransacao: 2, Moeda: "BRL", Emissor: "Brazil", Receptor: "Argentina", DataTransacao: time.Now()},
		{Id: 2, CodTransacao: 2, Moeda: "CHL", Emissor: "Chile", Receptor: "Colombia", DataTransacao: time.Now()},
	}
	c.IndentedJSON(200, transacaos)
}
