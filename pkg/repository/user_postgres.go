package repository

import (
	"fmt"
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

var _ User = (*UserPostgres)(nil)

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Exist(id int) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE id=$1", usersTable)
	row := r.db.QueryRow(query, id)
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (r *UserPostgres) Get(id int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, balance FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *UserPostgres) Refill(id int, amount float64) error {
	tx, err := r.db.Begin()

	query := fmt.Sprintf("UPDATE %s SET balance=balance+$1 WHERE id=$2", usersTable)
	_, err = tx.Exec(query, amount, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s (user_from ,user_to, amount, description) VALUES ($1, $2, $3, $4)", historyTable)
	_, err = tx.Exec(query, id, id, amount, refillDescription)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (r *UserPostgres) ExecuteTransfer(transfer models.Transfer) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("UPDATE %s SET balance=balance+$1 WHERE id=$2", usersTable)
	_, err = tx.Exec(query, transfer.Amount, transfer.UserTo)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("UPDATE %s SET balance=balance-$1 WHERE id=$2", usersTable)
	_, err = tx.Exec(query, transfer.Amount, transfer.UserFrom)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s (user_from ,user_to, amount, description) VALUES ($1, $2, $3, $4)", historyTable)
	_, err = tx.Exec(query, transfer.UserFrom, transfer.UserTo, transfer.Amount, transferDescription)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}
