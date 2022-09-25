package models

type User struct {
	ID         int    `json:"id,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Balance    int    `json:"balance,omitempty"`
}
