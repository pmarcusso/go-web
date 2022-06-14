package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/pmarcusso/go-web/internal/transaction"
	"github.com/pmarcusso/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "../../../transactions.json")
	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	p := NewTransaction(service)
	r := gin.Default()

	pr := r.Group("/transactions")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetTransaction_OK(t *testing.T) {

	r := createServer()
	req, responseRecorder := createRequestTest(http.MethodGet, "/transactions/", "")

	r.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestTransaction_Store(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/transactions/", `{
  "id": 1,
  "codTransaction": 4534534,
  "currency": "EUR",
  "issuer": "Londres",
  "receiver": "Portugal",
  "dateTransaction": "0001-01-01T00:00:00Z"
 }`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
