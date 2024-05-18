package get

import "time"

type Filter struct {
	Date     *time.Time `json:"date"`
	Amount   float32    `json:"amount"`
	Category string     `json:"category"`
}

type Paginate struct {
	ItemPerPage int `json:"itemPerPage"`
	Page        int `json:"page"`
}

type Expense struct {
	ID        int        `json:"id"`
	Date      *time.Time `json:"date"`
	Amount    float32    `json:"amount"`
	Category  string     `json:"category"`
	ImageUrl  string     `json:"imageUrl"`
	Note      string     `json:"note"`
	SpenderId string     `json:"spenderId"`
}
