package topics

import (
	"github.com/jinzhu/gorm"
)

//Topic type: Modelo usado por Gorm para el tratamiento de la db.
type Topic struct {
	gorm.Model
	Author     string `gorm:"type:varchar(50);PRIMARY_KEY"`
	TopicTitle string
	Refer      uint
	Comments   []Comment `gorm:"foreignkey:TopicTittle;association_foreignkey:TopicTitle"`
}

//Comment type: Modelo usado por Gorm para el tratamiento de la db.
type Comment struct {
	gorm.Model
	Author      string
	Body        string
	TopicTittle string
}
