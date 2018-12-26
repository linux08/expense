package models

//Expense model declaration
type Expense struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Reason string `json:"reason"`
	Vat    string `json:"vat"`
	User   *User  `json:"user_id"`
}
