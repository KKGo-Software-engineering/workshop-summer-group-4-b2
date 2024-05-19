package transaction

import (
	"database/sql"
	"fmt"
	"strings"
)

type Repository interface {
	GetAll(filter Filter, paginate Pagination) ([]Transaction, error)
	Create(request CreateTransactionRequest) (CreateTransactionResponse, error)
	GetExpenses(spenderId int) ([]Transaction, error)
	GetSummary(spenderId int, txnTypes []string) ([]GetTransactionResponse, error)
	UpdateExpense(transaction Transaction) error
	DeleteExpense(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

func (r repository) GetAll(filter Filter, paginate Pagination) ([]Transaction, error) {
	expenses := []Transaction{}
	query := "SELECT id, date, amount, category, image_url, note, spender_id FROM transaction"
	conditions := []string{}
	args := []interface{}{}

	if filter.Date != nil {
		conditions = append(conditions, fmt.Sprintf("date = $%d", len(args)+1))
		args = append(args, filter.Date)
	}
	if filter.Amount != 0 {
		conditions = append(conditions, fmt.Sprintf("amount = $%d", len(args)+1))
		args = append(args, filter.Amount)
	}
	if filter.Category != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", len(args)+1))
		args = append(args, filter.Category)
	}

	// Add WHERE clause if there are conditions
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	offset := (paginate.Page - 1) * paginate.ItemPerPage
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, paginate.ItemPerPage, offset)

	fmt.Println(query)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("hit")
			return []Transaction{}, nil
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		expense := Transaction{}
		err = rows.Scan(&expense.ID, &expense.Date, &expense.Amount, &expense.Category, &expense.ImageUrl, &expense.Note, &expense.SpenderId)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (r repository) Create(request CreateTransactionRequest) (CreateTransactionResponse, error) {
	return CreateTransactionResponse{}, nil
}

func (r repository) GetExpenses(spenderId int) ([]Transaction, error) {
	return nil, nil
}

func (r repository) GetSummary(spenderId int, txnTypes []string) ([]GetTransactionResponse, error) {
	return nil, nil
}

func (r repository) UpdateExpense(transaction Transaction) error {
	return nil
}

func (r repository) DeleteExpense(id int) error {
	return nil
}
