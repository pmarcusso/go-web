package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pmarcusso/go-web/pkg/web"
	"log"
	"net/http"
	"os"
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
