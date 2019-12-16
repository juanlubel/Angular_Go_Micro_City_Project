package admin

import "github.com/jinzhu/gorm"

//Admin type: Modelo usado por Gorm para el tratamiento de la db.
type Admin struct {
	gorm.Model
	Name  string
	Pass string
}
