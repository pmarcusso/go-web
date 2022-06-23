package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/pmarcusso/go-web/cmd/server/handler"
	"github.com/pmarcusso/go-web/docs"
	"github.com/pmarcusso/go-web/internal/transaction"
	"github.com/pmarcusso/go-web/pkg/web"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set token environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.NewResponse(http.StatusUnauthorized, nil, "token vazio"))
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.NewResponse(http.StatusUnauthorized, nil, "token inválido"),
			)
			return
		}

		c.Next()
	}
}

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

	//db := store.New(store.FileType, "../transactions.json")

	dataSource := "root:root@tcp(localhost:3306)/bootcamp?parseTime=true"
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	repo := transaction.NewRepository(conn)
	service := transaction.NewService(repo)
	controller := handler.NewTransaction(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	transactionGroup := r.Group("/transactions")
	{
		transactionGroup.Use(TokenAuthMiddleware())

		transactionGroup.GET("/", controller.GetAll())
		transactionGroup.GET("/:id", controller.GetOne())
		transactionGroup.POST("/", controller.Store())
		transactionGroup.PUT("/:id", controller.Update())
		transactionGroup.PATCH("/:id", controller.UpdateIssuerReceiver())
		transactionGroup.DELETE("/:id", controller.Delete())
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
