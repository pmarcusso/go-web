package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pmarcusso/go-web/internal/transaction"
	"net/http"
	"time"
)

type request struct {
	CodTransaction  int       `json:"codTransaction"`
	CurrencyType    string    `json:"currency"`
	Issuer          string    `json:"issuer"`
	Receiver        string    `json:"receiver"`
	DateTransaction time.Time `json:"dateTransaction"`
}

type Transaction struct {
	service transaction.Service
}

func NewTransaction(t transaction.Service) *Transaction {
	return &Transaction{service: t}
}

func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "você não tem permissão para fazer a solicitação solicitada.",
			})
			return
		}

		var req request

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		trans, err := t.service.Store(req.CodTransaction, req.CurrencyType, req.Issuer, req.Receiver, req.DateTransaction)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusCreated, trans)

	}
}

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "você não tem permissão para fazer a solicitação solicitada.",
			})
			return
		}

		t, err := t.service.GetAll()

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"data": &t,
		})
	}
}
