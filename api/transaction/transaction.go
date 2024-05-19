package transaction

import "time"

type Filter struct {
	Date     *time.Time `json:"date"`
	Amount   float32    `json:"amount"`
	Category string     `json:"category"`
}

type Pagination struct {
	ItemPerPage int `json:"itemPerPage"`
	Page        int `json:"page"`
}

type Transaction struct {
	ID        int        `json:"id"`
	Date      *time.Time `json:"date"`
	Amount    float32    `json:"amount"`
	Category  string     `json:"category"`
	ImageUrl  string     `json:"image_url"`
	Note      string     `json:"note"`
	SpenderId string     `json:"spender_id"`
}

type CreateTransactionRequest struct {
	Date      *time.Time `json:"date"`
	Amount    float32    `json:"amount"`
	Category  string     `json:"category"`
	ImageUrl  string     `json:"image_url"`
	Note      string     `json:"note"`
	SpenderId string     `json:"spender_id"`
	TxnType   string     `json:"transaction_type"`
}

type CreateTransactionResponse struct {
	ID int `json:"id"`
}

type SummaryResponse struct {
	TotalAmountSpend     float32 `json:"total_amount_spend"`
	AvgAmountSpendPerDay float32 `json:"average_amount_spend_per_day"`
	Total                int     `json:"total"`
}

type BalanceResponse struct {
	TotalAmountEarned float32 `json:"total_amount_earned"`
	TotalAmountSpend  float32 `json:"total_amount_spend"`
	TotalAmountSaved  float32 `json:"total_amount_saved"`
}

type GetTransactionResponse struct {
	ID        int        `json:"id"`
	Date      *time.Time `json:"date"`
	Amount    float32    `json:"amount"`
	Category  string     `json:"category"`
	ImageUrl  string     `json:"image_url"`
	Note      string     `json:"note"`
	SpenderId string     `json:"spender_id"`
	TxnType   string     `json:"transaction_type"`
}
