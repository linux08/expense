package models

//User model declaration
type User struct {
	ID       string `json:"user_id"`
	Name     string `json:"name"`
	Gender   string `json:"reason"`
	Email    string `json:"vat"`
	Password string `json:"password"`
}
