package repository

import (
	"fmt"
	"github.com/Ig0rVItalevich/avito-test/models"
	"github.com/jmoiron/sqlx"
)

type PurchasePostgres struct {
	db *sqlx.DB
}

var _ Purchase = (*PurchasePostgres)(nil)

func NewPurchasePostgres(db *sqlx.DB) *PurchasePostgres {
	return &PurchasePostgres{db: db}
}

func (r *PurchasePostgres) GetReservedTransaction(purchase models.Purchase) (models.Transaction, error) {
	var transaction models.Transaction
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND order_id=$2 AND product_id=$3 AND status=$4", transactionsTable)
	err := r.db.Get(&transaction, query, purchase.UserId, purchase.OrderId, purchase.ProductId, reservedStatus)

	return transaction, err
}

func (r *PurchasePostgres) Reserve(purchase models.Purchase) error {
	tx, err := r.db.Begin()
	query := fmt.Sprintf("UPDATE %s SET balance=balance-$1 WHERE id=$2", usersTable)
	_, err = tx.Exec(query, purchase.Amount, purchase.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("UPDATE %s SET balance=balance+$1 WHERE id=$2", usersTable)
	_, err = tx.Exec(query, purchase.Amount, companyAccount)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id, amount, order_id, product_id) VALUES($1, $2, $3, $4)", transactionsTable)
	_, err = tx.Exec(query, purchase.UserId, purchase.Amount, purchase.OrderId, purchase.ProductId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (r *PurchasePostgres) Accept(purchase models.Purchase) error {
	tx, err := r.db.Begin()
	query := fmt.Sprintf("UPDATE %s SET status=$1 WHERE user_id=$2 AND order_id=$3 AND product_id=$4 AND status=$5", transactionsTable)
	_, err = tx.Exec(query, acceptedStatus, purchase.UserId, purchase.OrderId, purchase.ProductId, reservedStatus)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s (user_from, user_to, amount, order_id, product_id, description) VALUES ($1, $2, $3, $4, $5, $6)", historyTable)
	_, err = tx.Exec(query, purchase.UserId, companyAccount, purchase.Amount, purchase.OrderId, purchase.ProductId, purchaseDescription)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (r *PurchasePostgres) Cancel(purchase models.Purchase) error {
	tx, err := r.db.Begin()
	query := fmt.Sprintf("UPDATE %s SET status=$1 WHERE user_id=$2 AND order_id=$3 AND product_id=$4 AND status=$5", transactionsTable)
	_, err = tx.Exec(query, canceledStatus, purchase.UserId, purchase.OrderId, purchase.ProductId, reservedStatus)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("UPDATE %s SET balance=balance-$1 WHERE id=$2", usersTable)
	_, err = tx.Exec(query, purchase.Amount, companyAccount)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("UPDATE %s SET balance=balance+$1 WHERE id=$2", usersTable)
	_, err = tx.Exec(query, purchase.Amount, purchase.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}
