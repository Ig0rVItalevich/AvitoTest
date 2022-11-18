package repository

import (
	"fmt"
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/jmoiron/sqlx"
	"math"
)

type ReportPostgres struct {
	db *sqlx.DB
}

var _ Report = (*ReportPostgres)(nil)

func NewReportPostgres(db *sqlx.DB) *ReportPostgres {
	return &ReportPostgres{db: db}
}

func (r *ReportPostgres) GetUserReport(input models.InputUserReport) ([]models.HistoryRow, error) {
	var history []models.HistoryRow

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_from=$1 OR user_to=$2", historyTable)
	row := r.db.QueryRow(query, input.UserId, input.UserId)
	var count int
	if err := row.Scan(&count); err != nil {
		return history, err
	}

	pagesCount := int(math.Ceil(float64(count) / float64(input.RecordsOnPage)))
	if pagesCount == 0 {
		return history, nil
	}

	offset := input.RecordsOnPage * (input.Page - 1)
	if offset >= count {
		offset = (pagesCount - 1) * input.RecordsOnPage
	}

	query = fmt.Sprintf("SELECT * FROM %s WHERE user_from=$1 OR user_to=$2 ORDER BY %s LIMIT %d OFFSET %d", historyTable, input.OrderBy, input.RecordsOnPage, offset)
	err := r.db.Select(&history, query, input.UserId, input.UserId)

	return history, err
}

func (r *ReportPostgres) GetRevenueReport(year int, month int) ([]models.RevenueReportRow, error) {
	var revenueReport []models.RevenueReportRow
	query := fmt.Sprintf(`SELECT product_id, sum(amount) AS sum FROM %s 
                            	WHERE extract(year FROM accepted_date)=$1 AND extract(month FROM accepted_date)=$2 AND description=$3
								GROUP BY product_id`, historyTable)
	err := r.db.Select(&revenueReport, query, year, month, purchaseDescription)

	return revenueReport, err
}
