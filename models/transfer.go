package models

type Transfer struct {
	UserFrom int     `json:"user_from"`
	UserTo   int     `json:"user_to"`
	Amount   float64 `json:"amount"`
}
