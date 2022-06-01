package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pmarcusso/go-web/cmd/server/handler"
	"github.com/pmarcusso/go-web/internal/transaction"
	"net/http"
	"time"
)

type Transaction struct {
	Id              int       `form:"id" json:"id"`
	CodTransaction  int       `form:"codTransaction" json:"codTransaction" binding:"required"`
	CurrencyType    string    `form:"currency" json:"currency" binding:"required"`
	Issuer          string    `form:"issuer" json:"issuer" binding:"required"`
	Receiver        string    `form:"receiver" json:"receiver" binding:"required"`
	DateTransaction time.Time `form:"dateTransaction" json:"dateTransaction" time_format:"2006-01-02"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "é necessário inserir um valor válido"
	}
	return "Erro desconhecido"
}

func main() {

	repo := transaction.NewRepository()
	service := transaction.NewService(repo)
	controller := handler.NewTransaction(service)

	r := gin.Default()

	transaction := r.Group("/transactions")
	{
		transaction.GET("/", controller.GetAll())
		transaction.GET("/:id", controller.GetOne())
		transaction.POST("/", controller.Store())
		transaction.PUT("/:id", controller.Update())
	}

	r.GET("/query", GetQueryParameterValueHandler)
	r.GET("/greetings", greetingsHandler)
	r.Run()
}

func GetQueryParameterValueHandler(c *gin.Context) {
	queryValue := c.Query("currency")

	c.JSON(http.StatusOK, gin.H{
		"currency": queryValue,
	})
}

func greetingsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Meu nome é Paulo Henrique",
	})
}
