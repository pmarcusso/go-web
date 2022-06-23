package repository

const (
	sqlGetAll         = "SELECT * FROM transactions"
	sqlStore          = "INSERT INTO transactions (cod_transaction, currency_type, issuer, receiver, date_transaction) VALUES (?, ?, ?, ?, ?)"
	sqlFindOne        = "SELECT * FROM transactions WHERE id = ?"
	sqlUpdate         = "UPDATE transactions SET cod_transaction=?, currency_type=?, issuer=?, receiver=? WHERE id=?"
	sqlUpdateIssuer   = "UPDATE transactions SET issuer=? WHERE id=?"
	sqlUpdateReceiver = "UPDATE transactions SET receiver=? WHERE id=?"
	slqDelete         = "DELETE FROM transactions WHERE id=?"
)
