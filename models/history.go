package models

type HistoryRow struct {
	Id           int     `json:"id" db:"id"`
	UserFrom     int     `json:"user_from" db:"user_from"`
	UserTo       int     `json:"user_to" db:"user_to"`
	Amount       float64 `json:"amount" db:"amount"`
	OrderId      int     `json:"order-id" db:"order_id"`
	ProductId    int     `json:"product-id" db:"product_id"`
	Description  string  `json:"description" db:"description"`
	AcceptedDate string  `json:"accepted_date" db:"accepted_date"`
}
