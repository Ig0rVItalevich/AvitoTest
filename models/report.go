package models

type InputUserReport struct {
	UserId        int    `json:"user_id"`
	OrderBy       string `json:"order_by"`
	RecordsOnPage int    `json:"records_on_page"`
	Page          int    `json:"page"`
}

func (i *InputUserReport) Validate() bool {
	if i.OrderBy != "sum" && i.OrderBy != "date" {
		return false
	}

	if i.RecordsOnPage <= 0 || i.Page <= 0 {
		return false
	}

	return true
}

type RevenueReportRow struct {
	ProductId int     `json:"product_id" db:"product_id"`
	Sum       float64 `json:"sum" db:"sum"`
}
