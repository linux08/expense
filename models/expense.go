package models

import "github.com/jinzhu/gorm"

//Expense model declaration
type Expense struct {
	gorm.Model
	Name string

	Reason string `json:"reason"`
	Vat    string `json:"vat"`
	UserID uint
	User   *User `json:"user_id"`
}
