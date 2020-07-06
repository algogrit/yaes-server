package entities

import "github.com/jinzhu/gorm"

// Payable represents the amount owed by a User to a UserID for an ExpenseID
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
