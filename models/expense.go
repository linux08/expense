package models

import "github.com/jinzhu/gorm"

//Expense model declaration
type Expense struct {
	gorm.Model

	Name   string
	Reason string `json:"reason"`
	Vat    string `json:"vat"`
	Num    int    `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
	UserID uint   //  `gorm:"index"`

	// User   *User  `json:"user_id"`
}
