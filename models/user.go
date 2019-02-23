package models

import (
	"github.com/jinzhu/gorm"
)

//User model declaration
type User struct {
	gorm.Model

	Name       string
	Email      string `gorm:"type:varchar(100);unique_index"`
	Gender     string `json:"Gender"`
	Password   string `json:"Password"`
	Num        int    `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
	ExpenseIDs []Expense
	Expense    Expense `gorm:"auto_preload"`
}
