package service

import (
	"github.com/pmarcusso/go-web/internal/transaction/domain"
	repository2 "github.com/pmarcusso/go-web/internal/transaction/repository"
	"log"
	"time"
)

type Service interface {
	GetOne(id int) (domain.Transaction, error)
	GetAll() ([]domain.Transaction, error)
	Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (domain.Transaction, error)
	Update(id, codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (domain.Transaction, error)
	UpdateIssuer(id int, issuer string) (domain.Transaction, error)
	UpdateReceiver(id int, receiver string) (domain.Transaction, error)
	Delete(id int) error
}

type service struct {
	repository repository2.Repository
}

func NewService(r repository2.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetOne(id int) (domain.Transaction, error) {
	oneTransaction, err := s.repository.GetOne(id)

	if err != nil {
		log.Println(err.Error())
		return domain.Transaction{}, err
	}

	return oneTransaction, nil
}

func (s *service) GetAll() ([]domain.Transaction, error) {
	transactionsList, err := s.repository.GetAll()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return transactionsList, nil
}

func (s *service) Update(id, codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (domain.Transaction, error) {
	updatedTransaction, err := s.repository.Update(id, codTransaction, currencyType, issuer, receiver, dateTransaction)

	if err != nil {
		log.Println(err.Error())
		return updatedTransaction, err
	}

	return updatedTransaction, err
}

func (s *service) UpdateIssuer(id int, issuer string) (domain.Transaction, error) {
	updateIssuer, err := s.repository.UpdateIssuer(id, issuer)

	if err != nil {
		log.Println(err.Error())
		return updateIssuer, err
	}

	return updateIssuer, err
}
func (s *service) UpdateReceiver(id int, receiver string) (domain.Transaction, error) {
	updateReceiver, err := s.repository.UpdateReceiver(id, receiver)

	if err != nil {
		log.Println(err.Error())
		return updateReceiver, err
	}

	return updateReceiver, err
}

func (s *service) Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (domain.Transaction, error) {

	newTransaction, err := s.repository.Store(codTransaction, currencyType, issuer, receiver, dateTransaction)

	if err != nil {
		log.Println(err)
		return domain.Transaction{}, err
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
