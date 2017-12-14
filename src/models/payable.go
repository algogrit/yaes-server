package model

import "github.com/jinzhu/gorm"

type Payable struct {
	gorm.Model
	User    User    `json:"-"`
	Expense Expense `json:"-"`

	UserID    uint `gorm:"not null"`
	ExpenseID uint `gorm:"not null"`

	RatioInPercentage float64 `gorm:"not null"`
	AmountOwed        float64 `gorm:"not null"`
	Status            string  `gorm:"not null"`
}
