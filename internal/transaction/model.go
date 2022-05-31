package transaction

import "time"

type Transaction struct {
	Id              int       `form:"id" json:"id"`
	CodTransaction  int       `form:"codTransaction" json:"codTransaction" binding:"required"`
	CurrencyType    string    `form:"currency" json:"currency" binding:"required"`
	Issuer          string    `form:"issuer" json:"issuer" binding:"required"`
	Receiver        string    `form:"receiver" json:"receiver" binding:"required"`
	DateTransaction time.Time `form:"dateTransaction" json:"dateTransaction" time_format:"2006-01-02"`
}

var transactions []Transaction
