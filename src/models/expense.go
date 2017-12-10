package model

import (
	"github.com/jinzhu/gorm"
)

type Expense struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:CreatedBy"`
	CreatedBy uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	Place     string  `gorm:"not null"`
}
