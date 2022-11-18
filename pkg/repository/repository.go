package repository

import (
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Exist(id int) (bool, error)
	Get(id int) (models.User, error)
	Refill(id int, amount float64) error
	ExecuteTransfer(transfer models.Transfer) error
}

type Purchase interface {
	GetReservedTransaction(purchase models.Purchase) (models.Transaction, error)
	Reserve(purchase models.Purchase) error
	Accept(purchase models.Purchase) error
	Cancel(purchase models.Purchase) error
}

type Report interface {
	GetUserReport(input models.InputUserReport) ([]models.HistoryRow, error)
	GetRevenueReport(year int, month int) ([]models.RevenueReportRow, error)
}

type Repository struct {
	User
	Purchase
	Report
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:     NewUserPostgres(db),
		Purchase: NewPurchasePostgres(db),
		Report:   NewReportPostgres(db),
	}
}
