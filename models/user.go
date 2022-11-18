package models

type User struct {
	Id      int     `json:"id" db:"id"`
	Balance float64 `json:"balance" db:"balance"`
}

type RefillBalance struct {
	UserId int     `json:"user_id"`
	Amount float64 `json:"amount"`
}
