package models

type BalanceInfo struct {
	ID           int            `json:"id,omitempty"`
	UserID       int            `json:"user_id,omitempty"`
	Transactions []*Transaction `json:"transactions,omitempty"`
}

type Transaction struct {
	ID        int `json:"id,omitempty"`
	UserID    int `json:"user_id,omitempty"`
	ForeignID int `json:"foreign_id,omitempty"`
	Amount    int `json:"amount,omitempty"`
}
