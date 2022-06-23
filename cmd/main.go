package main

import (
	"database/sql"
	"github.com/pmarcusso/go-web/internal/shared/middleware"
	"github.com/pmarcusso/go-web/internal/transaction/controller"
	"github.com/pmarcusso/go-web/internal/transaction/repository"
	"github.com/pmarcusso/go-web/internal/transaction/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/pmarcusso/go-web/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	if err := godotenv.Load("./.env"); err != nil {
		log.Println(err)
		log.Fatal("erro ao carregar o arquivo .env")
	}

	r := gin.Default()

	dataSource := "root:root@tcp(localhost:3306)/bootcamp?parseTime=true"
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(conn)
	newService := service.NewService(repo)
	newTransaction := controller.NewTransaction(newService)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	transactionGroup := r.Group("/transactions")
	{
		transactionGroup.Use(middleware.TokenAuthMiddleware())

		transactionGroup.GET("/", newTransaction.GetAll())
		transactionGroup.GET("/:id", newTransaction.GetOne())
		transactionGroup.POST("/", newTransaction.Store())
		transactionGroup.PUT("/:id", newTransaction.Update())
		transactionGroup.PATCH("/:id", newTransaction.UpdateIssuerReceiver())
		transactionGroup.DELETE("/:id", newTransaction.Delete())
	}

	r.GET("/query", GetQueryParameterValueHandler)
	r.GET("/greetings", greetingsHandler)

	err = r.Run()
	if err != nil {
		return
	}
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
