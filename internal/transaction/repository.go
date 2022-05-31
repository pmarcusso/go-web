package transaction

import "time"

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error)
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

func (r *repository) Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error) {
	newTrasaction := Transaction{CodTransaction: codTransaction, CurrencyType: currencyType, Issuer: issuer, Receiver: receiver, DateTransaction: dateTransaction}

	newTrasaction = generateId(&newTrasaction)

	transactions = append(transactions, newTrasaction)

	return newTrasaction, nil
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
