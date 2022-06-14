package transaction

import (
	"encoding/json"
	"github.com/pmarcusso/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_GetAll(t *testing.T) {
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
				Issuer:         "Londres",
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

		resp, _ := myRepo.GetAll()

		assert.Equal(t, input, resp)
	})

	t.Run("should create an empty transaction when len is zero", func(t *testing.T) {

		input := []Transaction{}
		dataJson, _ := json.Marshal(input)

		fileStore := store.New(store.FileType, "")
		dbMock := store.Mock{Data: dataJson}

		fileStore.AddMock(&dbMock)

		myRepo := NewRepository(fileStore)
		myService := NewService(myRepo)

		result, _ := myService.GetAll()

		assert.Equal(t, make([]Transaction, 0), result)

	})
}

func TestRepositoryDelete(t *testing.T) {
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
			Issuer:         "Londres",
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
	response := myRepo.Delete(1)
	assert.Nil(t, response)
}
