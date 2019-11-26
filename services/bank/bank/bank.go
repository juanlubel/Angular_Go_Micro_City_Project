package bank

import "github.com/jinzhu/gorm"

//Bank type: Modelo usado por Gorm para el tratamiento de la db.
type Bank struct {
	gorm.Model
	BankName string
}

/* //BankAccount type: Modelo usado por Gorm para el tratamiento de la db.
type BankAccount struct {
	gorm.Model
	AccountNumber string `gorm:"type:varchar(10);PRIMARY_KEY"`
	Owner         string `gorm:"type:varchar(50);`
	Balance       int
	Bank          Bank `gorm:"foreignkey:BankName"`
} */
