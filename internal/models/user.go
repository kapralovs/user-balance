package models

type User struct {
	ID      int `json:"id,omitempty"`
	Balance int `json:"balance,omitempty"`
}
