package service

import (
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/Ig0rVItalevich/avito-test/pkg/repository"
)

type User interface {
	Get(id int) (models.User, error)
	Refill(id int, amount float64) error
	ExecuteTransfer(transfer models.Transfer) error
}

type Purchase interface {
	Reserve(purchase models.Purchase) error
	Accept(purchase models.Purchase) error
	Cancel(purchase models.Purchase) error
}

type Report interface {
	GetUserReport(input models.InputUserReport) ([]models.HistoryRow, error)
	GetRevenueReport(year int, month int) (string, error)
}

type Service struct {
	User
	Purchase
	Report
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:     NewUserService(repos),
		Purchase: NewPurchaseService(repos),
		Report:   NewReportService(repos),
	}
}
