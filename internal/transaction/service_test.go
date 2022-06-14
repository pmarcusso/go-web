package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pmarcusso/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_GetAll(t *testing.T) {
	t.Run("should return a valid produc list", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")

		input := []Transaction{
			{
				Id:             1,
				CodTransaction: 4534534,
				CurrencyType:   "EUR",
				Issuer:         "Londres",
				Receiver:       "Portugal",
			}, {
				Id:             2,
				CodTransaction: 4534534,
				CurrencyType:   "EUR",
				Issuer:         "Amsterdam",
				Receiver:       "Portugal",
			},
		}

		dataJson, _ := json.Marshal(input)

		dbMock := &store.Mock{
			Data: dataJson,
			Err:  nil,
		}

		fileStore.AddMock(dbMock)

		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)
		resp, _ := myService.GetAll()

		fmt.Println("Resp:", resp)
		fmt.Println("Input:", input)

		assert.Equal(t, input, resp)
	})

	t.Run("test GetAll errors", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")

		expectedError := errors.New("error for GetAll")

		dbMock := &store.Mock{
			Data: []byte{},
			Err:  expectedError,
		}

		fileStore.AddMock(dbMock)

		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)

		_, err := myService.GetAll()

		assert.Equal(t, expectedError, err)
	})
}

func TestService_Store(t *testing.T) {
	t.Run("test store a transaction", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")
		input := []Transaction{
			{
				CodTransaction: 4534534,
				CurrencyType:   "EUR",
				Issuer:         "Londres",
				Receiver:       "Portugal",
			},
		}

		dbMock := &store.Mock{
			Data: []byte("[]"),
		}

		fileStore.AddMock(dbMock)
		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)
		response, _ := myService.Store(input[0].CodTransaction, input[0].CurrencyType, input[0].Issuer, input[0].Receiver, time.Now())

		assert.Equal(t, input[0].CodTransaction, response.CodTransaction)
		assert.Equal(t, input[0].CurrencyType, response.CurrencyType)
		assert.Equal(t, input[0].Issuer, response.Issuer)
		assert.Equal(t, input[0].Receiver, response.Receiver)
		assert.Equal(t, 1, response.Id)
	})

	t.Run("test error store a transaction", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")
		input := []Transaction{
			{
				CodTransaction: 4534534,
				CurrencyType:   "EUR",
				Issuer:         "Londres",
				Receiver:       "Portugal",
			},
		}

		expectedError := errors.New("error store service")
		dbMock := &store.Mock{
			Err: expectedError,
		}

		fileStore.AddMock(dbMock)
		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)
		result, resultError := myService.Store(input[0].CodTransaction, input[0].CurrencyType, input[0].Issuer, input[0].Receiver, time.Now())

		assert.Equal(t, expectedError, resultError)
		assert.Equal(t, Transaction{}, result)
	})
}

func TestService_Update(t *testing.T) {
	t.Run("test update a transaction", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")
		input := []Transaction{
			{
				Id:             1,
				CodTransaction: 4534534,
				CurrencyType:   "YEN",
				Issuer:         "Japan",
				Receiver:       "Coreia do Norte",
			},
		}
		dataJson, _ := json.Marshal(input)

		dbMock := store.Mock{
			Data: dataJson,
		}

		fileStore.AddMock(&dbMock)

		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)

		response, _ := myService.Update(input[0].Id, 123, "BR", "Brazil", "Argentina", time.Now())

		assert.Equalf(t, 123, response.CodTransaction, "CodTransaction - Should be Equal")
		assert.Equalf(t, "BR", response.CurrencyType, "CurrencyType - Should be Equal")
		assert.Equalf(t, "Brazil", response.Issuer, "Issuer - Should be Equal")
		assert.Equalf(t, "Argentina", response.Receiver, "Receiver - Should be Equal")
		assert.Equal(t, 1, response.Id)
	})

	t.Run("test error store a transaction", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")
		expectedError := errors.New("test error update")

		input := []Transaction{
			{
				Id:             2,
				CodTransaction: 4534534,
				CurrencyType:   "EUR",
				Issuer:         "Londres",
				Receiver:       "Portugal",
			},
		}

		dbMock := store.Mock{
			Err: expectedError,
		}

		fileStore.AddMock(&dbMock)

		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)

		_, err := myService.Update(input[0].Id, input[0].CodTransaction, input[0].CurrencyType, input[0].Issuer, input[0].Receiver, time.Now())
		assert.Equal(t, expectedError, err)
	})
}

func TestService_UpdateIssuer(t *testing.T) {
	t.Run("should update field issuer", func(t *testing.T) {

		fileStore := store.New(store.FileType, "")

		input := []Transaction{
			{
				Id:             2,
				CodTransaction: 4534534,
				CurrencyType:   "EUR",
				Issuer:         "Londres",
				Receiver:       "Portugal",
			},
		}

		dataJson, _ := json.Marshal(input)

		dbMock := store.Mock{Data: dataJson}
		fileStore.AddMock(&dbMock)

		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)

		result, _ := myService.UpdateIssuer(2, "Brasil")
		assert.Equal(t, "Brasil", result.Issuer)
	})

	t.Run("should return error when id does not exists", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")

		input := []Transaction{
			{
				Id:             2,
				CodTransaction: 4534534,
				CurrencyType:   "EUR",
				Issuer:         "Londres",
				Receiver:       "Portugal",
			},
		}

		dataJson, _ := json.Marshal(input)

		dbMock := store.Mock{Data: dataJson}
		fileStore.AddMock(&dbMock)

		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)
		result, err := myService.UpdateIssuer(3, "Brasil")
		assert.Error(t, err)
		assert.NotEqual(t, input[0].Issuer, result.Issuer)
	})
}

func TestService_UpdateReceiver(t *testing.T) {
	t.Run("should update receiver field", func(t *testing.T) {

		fileStore := store.New(store.FileType, "")

		input := []Transaction{
			{
				Id:             2,
				CodTransaction: 4534534,
				CurrencyType:   "EUR",
				Issuer:         "Londres",
				Receiver:       "Portugal",
			},
		}

		dataJson, _ := json.Marshal(input)

		dbMock := store.Mock{Data: dataJson}
		fileStore.AddMock(&dbMock)

		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)
		result, _ := myService.UpdateReceiver(2, "Brasil")

		assert.Equal(t, "Brasil", result.Receiver)

	})
}

func TestService_Delete(t *testing.T) {
	t.Run("test service delete", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")

		input := []Transaction{
			{
				Id:             2,
				CodTransaction: 4534534,
				CurrencyType:   "EUR",
				Issuer:         "Londres",
				Receiver:       "Portugal",
			},
		}

		dataJson, _ := json.Marshal(input)

		dbMock := store.Mock{Data: dataJson}

		fileStore.AddMock(&dbMock)

		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)
		resp := myService.Delete(2)

		assert.Nil(t, resp)
	})

	t.Run("test service delete", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")

		input := []Transaction{
			{
				Id:             2,
				CodTransaction: 4534534,
				CurrencyType:   "EUR",
				Issuer:         "Londres",
				Receiver:       "Portugal",
			},
		}

		dataJson, _ := json.Marshal(input)

		dbMock := store.Mock{Data: dataJson}

		fileStore.AddMock(&dbMock)

		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)
		resp := myService.Delete(3)

		assert.Error(t, resp)
		assert.Equal(t, "id não encontrado", resp.Error())
	})
}
