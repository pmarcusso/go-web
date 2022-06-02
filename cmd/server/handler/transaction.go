package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pmarcusso/go-web/internal/transaction"
	"net/http"
	"os"
	"strconv"
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

func (t *Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		token := c.GetHeader("token")

		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "você não tem permissão para fazer a solicitação solicitada.",
			})
			return
		}

		idConvertido, err := strconv.Atoi(id)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "id is not a number",
			})
			fmt.Println(err)
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if err := validateFields(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}

		updatedTransaction, err := t.service.Update(idConvertido, req.CodTransaction, req.CurrencyType, req.Issuer, req.Receiver, req.DateTransaction)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, updatedTransaction)
	}
}

func (t *Transaction) UpdateIssuerReceiver() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		token := c.GetHeader("token")

		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "você não tem permissão para fazer a solicitação solicitada.",
			})
			return
		}

		idConvertido, err := strconv.Atoi(id)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "id is not a number",
			})
			fmt.Println(err)
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if req.Receiver != "" {
			updatedReceiver, err := t.service.UpdateReceiver(idConvertido, req.Receiver)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, updatedReceiver)
			return
		}

		if req.Issuer != "" {
			updatedIssuer, err := t.service.UpdateIssuer(idConvertido, req.Issuer)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, updatedIssuer)
			return

		}

	}

}

func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "você não tem permissão para continuar.",
			})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := validateFields(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}

		trans, err := t.service.Store(req.CodTransaction, req.CurrencyType, req.Issuer, req.Receiver, req.DateTransaction)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, trans)
	}
}

func (t *Transaction) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		token := c.GetHeader("token")

		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "você não tem permissão para fazer a solicitação solicitada.",
			})
			return
		}

		idConvertido, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "id is not a number",
			})
			fmt.Println(err)
			return
		}

		oneTransaction, err := t.service.GetOne(idConvertido)

		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"error": "transação não encontrada",
			})
			return
		}

		c.JSON(http.StatusOK, oneTransaction)
		return
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

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "você não tem permissão para fazer a solicitação solicitada.",
			})
			return
		}

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "id is not a number",
			})
			fmt.Println(err)
			return
		}

		err = t.service.Delete(idConvertido)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("o produto %d for removido", idConvertido)})
	}
}

//TODO CRIAR STRUCT DE ERROR
func validateFields(req request) error {

	if req.CodTransaction == 0 {
		return errors.New("O campo [codTransaction] está vazio ou nulo")
	}

	if req.CurrencyType == "" {
		return errors.New("O campo [currency] está vazio ou nulo")
	}

	if req.Issuer == "" {
		return errors.New("O campo [issuer] está vazio ou nulo")
	}

	if req.Receiver == "" {
		return errors.New("O campo [receiver] está vazio ou nulo")
	}

	if req.DateTransaction.String() == "" {
		return errors.New("O campo [dateTransaction] está vazio ou nulo")
	}

	return nil
}
