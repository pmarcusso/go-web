package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pmarcusso/go-web/internal/transaction"
	"github.com/pmarcusso/go-web/pkg/web"
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

		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, "id is not a number"))
			fmt.Println(err)
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		}

		if err := validateFields(req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		updatedTransaction, err := t.service.Update(idConvertido, req.CodTransaction, req.CurrencyType, req.Issuer, req.Receiver, req.DateTransaction)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedTransaction, ""))
	}
}

func (t *Transaction) UpdateIssuerReceiver() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		token := c.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		idConvertido, err := strconv.Atoi(id)
		if err != nil {

			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, "id is not a number"))
			fmt.Println(err)
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		}

		if req.Receiver != "" {
			updatedReceiver, err := t.service.UpdateReceiver(idConvertido, req.Receiver)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedReceiver, ""))
			return
		}

		if req.Issuer != "" {
			updatedIssuer, err := t.service.UpdateIssuer(idConvertido, req.Issuer)
			if err != nil {
				c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
				return
			}
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedIssuer, ""))
			return
		}
	}
}

// Store StoreTransactions godoc
// @Summary Store transactions
// @Tags Transactions
// @Description store transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Transaction to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 401 {object} web.Response
// @Router /transactions [post]
func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if err := validateFields(req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		trans, err := t.service.Store(req.CodTransaction, req.CurrencyType, req.Issuer, req.Receiver, req.DateTransaction)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		c.JSON(http.StatusCreated, trans)
	}
}

func (t *Transaction) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// token := c.GetHeader("token")

		// if token != os.Getenv("TOKEN") {
		// 	c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
		// 	return
		// }

		idConvertido, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, "id is not a number"))
			fmt.Println(err)
			return
		}

		oneTransaction, err := t.service.GetOne(idConvertido)

		if err != nil {
			c.IndentedJSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "transação não encontrada"))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(http.StatusNotFound, oneTransaction, ""))
		return
	}
}

// GetAll ListTransactions godoc
// @Summary List products
// @Tags Transactions
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} request
// @Router /transactions [get]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		t, err := t.service.GetAll()

		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "não há transações"))
		}

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, &t, ""))
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, "id is not a number"))
			fmt.Println(err)
			return
		}

		err = t.service.Delete(idConvertido)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusNotFound, fmt.Sprintf("o produto %d for removido", idConvertido), ""))
	}
}

func validateFields(req request) error {

	if req.CodTransaction == 0 {
		return errors.New("o campo [codTransaction] está vazio ou nulo")
	}

	if req.CurrencyType == "" {
		return errors.New("o campo [currency] está vazio ou nulo")
	}

	if req.Issuer == "" {
		return errors.New("o campo [issuer] está vazio ou nulo")
	}

	if req.Receiver == "" {
		return errors.New("o campo [receiver] está vazio ou nulo")
	}

	if req.DateTransaction.String() == "" {
		return errors.New("o campo [dateTransaction] está vazio ou nulo")
	}

	return nil
}
