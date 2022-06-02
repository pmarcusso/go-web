package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/pmarcusso/go-web/cmd/server/handler"
	"github.com/pmarcusso/go-web/internal/transaction"
	"github.com/pmarcusso/go-web/pkg/store"
	"log"
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

	if err := godotenv.Load(); err != nil {
		log.Fatal("erro ao carregar o arquivo .env")
	}

	db := store.New(store.FileType, "./transactions.json")
	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	controller := handler.NewTransaction(service)

	r := gin.Default()

	transactionGroup := r.Group("/transactions")
	{
		transactionGroup.GET("/", controller.GetAll())
		transactionGroup.GET("/:id", controller.GetOne())
		transactionGroup.POST("/", controller.Store())
		transactionGroup.PUT("/:id", controller.Update())
		transactionGroup.PATCH("/:id", controller.UpdateIssuerReceiver())
		transactionGroup.DELETE("/:id", controller.Delete())
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
