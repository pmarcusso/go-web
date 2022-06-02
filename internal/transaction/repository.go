package transaction

import (
	"errors"
	"github.com/pmarcusso/go-web/pkg/store"
	"log"
	"time"
)

type Repository interface {
	GetOne(id int) (Transaction, error)
	GetAll() ([]Transaction, error)
	Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error)
	Update(id, codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error)
	UpdateIssuer(id int, issuer string) (Transaction, error)
	UpdateReceiver(id int, receiver string) (Transaction, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Transaction, error) {

	r.db.Read(&transactions)

	//if len(transactions) == 0 {
	//	transactions = make([]Transaction, 0)
	//}

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

	err := r.db.Read(&transactions)
	if err != nil {
		return Transaction{}, err
	}
	newTrasaction := Transaction{CodTransaction: codTransaction, CurrencyType: currencyType, Issuer: issuer, Receiver: receiver, DateTransaction: dateTransaction}
	newTrasaction = r.generateId(&newTrasaction)
	transactions = append(transactions, newTrasaction)

	if err := r.db.Write(transactions); err != nil {
		return Transaction{}, err
	}

	return newTrasaction, nil
}

func (r *repository) Update(id, codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error) {
	err := r.db.Read(&transactions)
	if err != nil {
		return Transaction{}, err
	}

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

	for i := range transactions {
		if transactions[i].Id == transaction.Id {
			transactions[i] = transaction
		}
	}

	if err := r.db.Write(transactions); err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func (r *repository) UpdateIssuer(id int, issuer string) (Transaction, error) {
	err := r.db.Read(&transactions)
	if err != nil {
		return Transaction{}, err
	}

	issuerTransaction, err := r.GetOne(id)

	if err != nil {
		log.Println(err.Error())
		return issuerTransaction, err
	}

	for i := range transactions {
		if transactions[i].Id == issuerTransaction.Id {
			transactions[i].Issuer = issuer
			issuerTransaction = transactions[i]
		}
	}

	if err := r.db.Write(transactions); err != nil {
		return Transaction{}, err
	}

	return issuerTransaction, nil
}

func (r *repository) UpdateReceiver(id int, receiver string) (Transaction, error) {
	err := r.db.Read(&transactions)
	if err != nil {
		return Transaction{}, err
	}

	receiverTransaction, err := r.GetOne(id)
	if err != nil {
		log.Println(err.Error())
		return receiverTransaction, err
	}

	receiverTransaction.Receiver = receiver

	for i := range transactions {
		if transactions[i].Id == receiverTransaction.Id {
			transactions[i].Receiver = receiverTransaction.Receiver
			receiverTransaction = transactions[i]
		}
	}

	if err := r.db.Write(transactions); err != nil {
		return Transaction{}, err
	}

	return receiverTransaction, nil
}

func (r *repository) Delete(id int) error {
	err := r.db.Read(&transactions)
	if err != nil {
		return err
	}
	transaction, err := r.GetOne(id)
	var index int

	for i := range transactions {
		if transactions[i].Id == transaction.Id {
			index = i
		}
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}

	transactions = append(transactions[:index], transactions[index+1:]...)
	if err := r.db.Write(transactions); err != nil {
		return err
	}
	return nil
}

func (r *repository) generateId(transaction *Transaction) Transaction {

	if err := r.db.Read(&transactions); err != nil {
		return Transaction{}
	}

	transLen := len(transactions)

	if transLen == 0 {
		transaction.Id = 1
		return *transaction
	}

	lastTransaction := transactions[transLen-1]

	transaction.Id = lastTransaction.Id + 1

	return *transaction
}
