package entities

import "github.com/jinzhu/gorm"

// Expense represents the shared paid by a User
type Expense struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:CreatedBy" json:"-"`
	CreatedBy uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	Place     string  `gorm:"not null"`
	Payables  []Payable
}
