package transaction

import (
	"errors"
	"log"
	"time"
)

type Repository interface {
	GetOne(id int) (Transaction, error)
	GetAll() ([]Transaction, error)
	Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error)
	Update(id, codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Transaction, error) {
	if len(transactions) == 0 {
		transactions = make([]Transaction, 0)
	}

	return transactions, nil
}

func (r *repository) GetOne(id int) (Transaction, error) {
	for _, transaction := range transactions {
		if id == transaction.Id {
			return transaction, nil
		}
	}
	err := errors.New("id não encontrado")

	return Transaction{}, err
}

func (r *repository) Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error) {
	newTrasaction := Transaction{CodTransaction: codTransaction, CurrencyType: currencyType, Issuer: issuer, Receiver: receiver, DateTransaction: dateTransaction}

	newTrasaction = generateId(&newTrasaction)

	transactions = append(transactions, newTrasaction)

	return newTrasaction, nil
}

func (r *repository) Update(id, codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error) {
	transaction, err := r.GetOne(id)

	if err != nil {
		log.Println(err.Error())
		return transaction, err
	}

	transaction.CodTransaction = codTransaction
	transaction.CurrencyType = currencyType
	transaction.Issuer = issuer
	transaction.Receiver = receiver
	transaction.DateTransaction = dateTransaction

	for i, _ := range transactions {
		if transactions[i].Id == transaction.Id {
			transactions[i] = transaction
		}
	}

	return transaction, nil
}

func generateId(transaction *Transaction) Transaction {

	transLen := len(transactions)

	if transLen == 0 {
		transaction.Id = 1
		return *transaction
	}

	lastTransaction := transactions[transLen-1]

	transaction.Id = lastTransaction.Id + 1

	return *transaction
}
