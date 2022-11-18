package service

import (
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

func (s *UserService) Exist(id int) (bool, error) {
	flag, err := s.repos.Exist(id)
	if err != nil {
		return false, err
	}

	return flag, nil
}

func (s *UserService) Get(id int) (models.User, error) {
	flag, err := s.Exist(id)
	if err != nil {
		return models.User{}, err
	}

	if !flag {
		return models.User{}, ErrUserDoesNotExist
	}

	return s.repos.Get(id)
}

func (s *UserService) Refill(id int, amount float64) error {
	flag, err := s.Exist(id)
	if err != nil {
		return err
	}

	if !flag {
		return ErrUserDoesNotExist
	}

	return s.repos.Refill(id, amount)
}

func (s *UserService) ExecuteTransfer(transfer models.Transfer) error {
	flag, err := s.Exist(transfer.UserTo)
	if err != nil {
		return err
	}

	if !flag {
		return ErrUserDoesNotExist
	}

	flag, err = s.Exist(transfer.UserFrom)
	if err != nil {
		return err
	}

	if !flag {
		return ErrUserDoesNotExist
	}

	user, err := s.Get(transfer.UserFrom)
	if err != nil {
		return err
	}
	if user.Balance-transfer.Amount < 0 {
		return ErrNotEnoughMoney
	}

	return s.repos.ExecuteTransfer(transfer)
}
