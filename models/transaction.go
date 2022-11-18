package models

type Transaction struct {
	Id          int     `json:"id" db:"id"`
	UserFrom    int     `json:"user_id" db:"user_id"`
	Amount      float64 `json:"amount" db:"amount"`
	OrderId     int     `json:"order_id" db:"order_id"`
	ProductId   int     `json:"product_id" db:"product_id"`
	Status      string  `json:"status" db:"status"`
	CreatedDate string  `json:"created_date" db:"created_date"`
}
