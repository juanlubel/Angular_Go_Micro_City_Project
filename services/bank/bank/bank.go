package bank

import "github.com/jinzhu/gorm"

//Bank type: Modelo usado por Gorm para el tratamiento de la db.
type Bank struct {
	gorm.Model
	BankName string
}
