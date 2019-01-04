package models

import (
	"github.com/jinzhu/gorm"
)

//User model declaration
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"type:varchar(100);unique_index"`

	Gender   string `json:"reason"`
	Password string `json:"password"`
	Num      int    `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
	Expense  []Expense
}
