package accounts

import (
	"github.com/jinzhu/gorm"
)

//BankAccount type: Modelo usado por Gorm para el tratamiento de la db.
type BankAccount struct {
	gorm.Model
	AccountNumber string `gorm:"type:varchar(10);PRIMARY_KEY"`
	Owner         string
	Balance       int
	Bank          string `gorm:"foreignkey:BankName"`
}
