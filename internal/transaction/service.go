package transaction

import (
	"log"
	"time"
)

type Service interface {
	GetAll() ([]Transaction, error)
	Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Transaction, error) {
	transactionsList, err := s.repository.GetAll()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return transactionsList, nil
}

func (s *service) Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error) {

	newTransaction, err := s.repository.Store(codTransaction, currencyType, issuer, receiver, dateTransaction)

	if err != nil {
		log.Println(err)
		return Transaction{}, err
	}

	return newTransaction, nil
}
