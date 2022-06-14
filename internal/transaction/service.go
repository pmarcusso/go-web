package transaction

import (
	"log"
	"time"
)

type Service interface {
	GetOne(id int) (Transaction, error)
	GetAll() ([]Transaction, error)
	Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error)
	Update(id, codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error)
	UpdateIssuer(id int, issuer string) (Transaction, error)
	UpdateReceiver(id int, receiver string) (Transaction, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetOne(id int) (Transaction, error) {
	oneTransaction, err := s.repository.GetOne(id)

	if err != nil {
		log.Println(err.Error())
		return Transaction{}, err
	}

	return oneTransaction, nil
}

func (s *service) GetAll() ([]Transaction, error) {
	transactionsList, err := s.repository.GetAll()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return transactionsList, nil
}

func (s *service) Update(id, codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error) {
	updatedTransaction, err := s.repository.Update(id, codTransaction, currencyType, issuer, receiver, dateTransaction)

	if err != nil {
		log.Println(err.Error())
		return updatedTransaction, err
	}

	return updatedTransaction, err
}

func (s *service) UpdateIssuer(id int, issuer string) (Transaction, error) {
	updateIssuer, err := s.repository.UpdateIssuer(id, issuer)

	if err != nil {
		log.Println(err.Error())
		return updateIssuer, err
	}

	return updateIssuer, err
}
func (s *service) UpdateReceiver(id int, receiver string) (Transaction, error) {
	updateReceiver, err := s.repository.UpdateReceiver(id, receiver)

	if err != nil {
		log.Println(err.Error())
		return updateReceiver, err
	}

	return updateReceiver, err
}

func (s *service) Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (Transaction, error) {

	newTransaction, err := s.repository.Store(codTransaction, currencyType, issuer, receiver, dateTransaction)

	if err != nil {
		log.Println(err)
		return Transaction{}, err
	}

	return newTransaction, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}
	return err
}
