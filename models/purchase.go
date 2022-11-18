package models

type Purchase struct {
	UserId    int     `json:"user_id"`
	OrderId   int     `json:"order_id"`
	ProductId int     `json:"product_id"`
	Amount    float64 `json:"amount"`
}

func (p *Purchase) Validate() bool {
	if p.UserId <= 0 || p.OrderId <= 0 || p.ProductId <= 0 || p.Amount <= 0 {
		return false
	}

	return true
}
