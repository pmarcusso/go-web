package repository

import (
	"database/sql"
	"errors"
	"github.com/pmarcusso/go-web/internal/transaction/domain"
	"time"
)

type Repository interface {
	GetOne(id int) (domain.Transaction, error)
	GetAll() ([]domain.Transaction, error)
	Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (domain.Transaction, error)
	Update(id, codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (domain.Transaction, error)
	UpdateIssuer(id int, issuer string) (domain.Transaction, error)
	UpdateReceiver(id int, receiver string) (domain.Transaction, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Transaction, error) {

	transactions := []domain.Transaction{}

	rows, err := r.db.Query(sqlGetAll)

	if err != nil {
		return transactions, err
	}

	defer rows.Close() // impedir vazamento de memoria

	for rows.Next() {
		var transaction domain.Transaction

		err := rows.Scan(&transaction.Id, &transaction.CodTransaction, &transaction.CurrencyType, &transaction.Issuer,
			&transaction.Receiver, &transaction.DateTransaction)

		if err != nil {
			return transactions, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *repository) GetOne(id int) (domain.Transaction, error) {

	var transaction domain.Transaction

	transactions, err := r.GetAll()

	if err != nil {
		return transaction, err
	}

	for _, transaction := range transactions {
		if id == transaction.Id {
			return transaction, nil
		}
	}

	return transaction, errors.New("id não encontrado")
}

func (r *repository) Store(codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (domain.Transaction, error) {

	transaction := domain.Transaction{CodTransaction: codTransaction, CurrencyType: currencyType, Issuer: issuer, Receiver: receiver, DateTransaction: time.Now()}

	stmt, err := r.db.Prepare(sqlStore)

	if err != nil {
		return transaction, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(&transaction.CodTransaction, &transaction.CurrencyType, &transaction.Issuer,
		&transaction.Receiver, &transaction.DateTransaction)

	if err != nil {
		return transaction, err
	}

	lastID, err := res.LastInsertId()

	transaction.Id = int(lastID)

	return transaction, nil
}

func (r *repository) Update(id, codTransaction int, currencyType, issuer, receiver string, dateTransaction time.Time) (domain.Transaction, error) {

	transaction, err := r.GetOne(id)

	if err != nil {
		return transaction, err
	}

	transaction.CodTransaction = codTransaction
	transaction.CurrencyType = currencyType
	transaction.Issuer = issuer
	transaction.Receiver = receiver
	transaction.DateTransaction = dateTransaction

	stmt, err := r.db.Prepare(sqlUpdate)

	if err != nil {
		return transaction, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&transaction.CodTransaction, &transaction.CurrencyType, &transaction.Issuer, &transaction.Receiver, &transaction.Id)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) UpdateIssuer(id int, issuer string) (domain.Transaction, error) {

	issuerTransaction, err := r.GetOne(id)

	if err != nil {
		return issuerTransaction, err
	}

	issuerTransaction.Issuer = issuer

	stmt, err := r.db.Prepare(sqlUpdateIssuer)

	if err != nil {
		return issuerTransaction, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&issuerTransaction.Issuer, &issuerTransaction.Id)

	if err != nil {
		return issuerTransaction, err
	}

	return issuerTransaction, nil
}

func (r *repository) UpdateReceiver(id int, receiver string) (domain.Transaction, error) {
	receiverTransaction, err := r.GetOne(id)

	if err != nil {
		return receiverTransaction, err
	}

	receiverTransaction.Receiver = receiver

	stmt, err := r.db.Prepare(sqlUpdateReceiver)

	if err != nil {
		return receiverTransaction, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&receiverTransaction.Receiver, &receiverTransaction.Id)

	return receiverTransaction, nil
}

func (r *repository) Delete(id int) error {

	transaction, err := r.GetOne(id)

	if err != nil {
		return err
	}

	stmt, err := r.db.Prepare(slqDelete)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&transaction.Id)

	if err != nil {
		return err
	}

	return nil
	//err := r.db.Read(&transactions)
	//if err != nil {
	//	return err
	//}
	//transaction, err := r.GetOne(id)
	//var index int
	//
	//for i := range transactions {
	//	if transactions[i].Id == transaction.Id {
	//		index = i
	//	}
	//}
	//
	//if err != nil {
	//	log.Println(err.Error())
	//	return err
	//}
	//
	//transactions = append(transactions[:index], transactions[index+1:]...)
	//if err := r.db.Write(&transactions); err != nil {
	//	return err
	//}
	//return nil
}

//func (r *repository) generateId(transaction *Transaction) Transaction {
//
//	if err := r.db.Read(&transactions); err != nil {
//		return Transaction{}
//	}
//
//	transLen := len(transactions)
//
//	if transLen == 0 {
//		transaction.Id = 1
//		return *transaction
//	}
//
//	lastTransaction := transactions[transLen-1]
//
//	transaction.Id = lastTransaction.Id + 1
//
//	return *transaction
//}
