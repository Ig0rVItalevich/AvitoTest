package service

import (
	"errors"
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/Ig0rVItalevich/avito-test/pkg/repository"
)

type UserService struct {
	repos repository.User
}

var _ User = (*UserService)(nil)

func NewUserService(repos repository.User) *UserService {
	return &UserService{repos: repos}
}

func (s *UserService) Get(id int) (models.User, error) {
	return s.repos.Get(id)
}

func (s *UserService) Refill(id int, amount float64) error {
	return s.repos.Refill(id, amount)
}

func (s *UserService) ExecuteTransfer(transfer models.Transfer) error {
	user, err := s.Get(transfer.UserFrom)
	if err != nil {
		return err
	}
	if user.Balance-transfer.Amount < 0 {
		return errors.New("not enough money")
	}

	return s.repos.ExecuteTransfer(transfer)
}
