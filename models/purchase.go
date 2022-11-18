package models

type Purchase struct {
	UserId    int     `json:"user_id"`
	OrderId   int     `json:"order_id"`
	ProductId int     `json:"product_id"`
	Amount    float64 `json:"amount"`
}
