package models

type Expense struct {
	ID     string `json: "id"`
	Name   string `json: "name"`
	Reason string `json: "reason"`
	Vat    string `json: "vat"`
	User   *User  `json: "id"`
}

type User struct {
	ID       string `json: "id"`
	Name     string `json: "name"`
	Gender   string `json: "reason"`
	email    string `json: "vat"`
	Password string `json: "password"`
}
