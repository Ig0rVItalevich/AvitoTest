package models

type Transfer struct {
	UserFrom int     `json:"user_from"`
	UserTo   int     `json:"user_to"`
	Amount   float64 `json:"amount"`
}

func (t *Transfer) Validate() bool {
	if t.Amount <= 0 || t.UserTo <= 0 || t.UserFrom <= 0 {
		return false
	}

	return true
}
