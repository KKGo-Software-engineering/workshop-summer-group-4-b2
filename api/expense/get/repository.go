package get

import (
	"database/sql"
)

type Repository interface {
	GetAll(filter Filter, paginate Paginate) ([]Expense, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

func (r repository) GetAll(filter Filter, paginate Paginate) ([]Expense, error) {
	expenses := []Expense{}
	query := "SELECT id, date, amount, category, image_url, note, spender_id FROM transaction"

	rows, err := r.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []Expense{}, nil
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		expense := Expense{}
		err = rows.Scan(&expense.ID, &expense.Date, &expense.Amount, &expense.Category, &expense.ImageUrl, &expense.Note, &expense.SpenderId)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}
