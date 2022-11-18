package service

import (
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/Ig0rVItalevich/avito-test/pkg/repository"
)

type PurchaseService struct {
	repos repository.Purchase
}

var _ Purchase = (*PurchaseService)(nil)

func NewPurchaseService(repos repository.Purchase) *PurchaseService {
	return &PurchaseService{repos: repos}
}

func (s *PurchaseService) Reserve(purchase models.Purchase) error {
	_, err := s.repos.GetReservedTransaction(purchase)
	if err == nil {
		return ErrReservedPurchaseAlreadyExists
	}

	return s.repos.Reserve(purchase)
}

func (s *PurchaseService) Accept(purchase models.Purchase) error {
	transaction, err := s.repos.GetReservedTransaction(purchase)
	if err != nil || transaction.Amount != purchase.Amount {
		return ErrReservedPurchaseDoesNotExist
	}

	return s.repos.Accept(purchase)
}

func (s *PurchaseService) Cancel(purchase models.Purchase) error {
	transaction, err := s.repos.GetReservedTransaction(purchase)
	if err != nil || transaction.Amount != purchase.Amount {
		return ErrReservedPurchaseDoesNotExist
	}

	return s.repos.Cancel(purchase)
}
