package admin

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql
)

// AdminRepository : obtain the context of the database
type Repository struct {
	DB *gorm.DB
}

// ProvideProductRepostiory : provides the repository to use the db context
func ProvideAdminRepostiory(DB *gorm.DB) Repository {
	return Repository{DB: DB}
}


//FindAll : obtain from db all the data of the indicated gorm model
func (p *Repository) FindAll() []Admin {
	var admin []Admin
	p.DB.Find(&admin)

	return admin
}

//FindByID :  obtain from db the data that contains the param id
func (p *Repository) FindByID(id uint) Admin {
	var admin Admin
	p.DB.First(&admin)

	return admin
}

//FindByName : used for login Admin User
func (p *Repository) FindByName(condition interface{}) Admin {
	var admin Admin
	p.DB.Where(condition).First(&admin)
	//err := db.Where(condition).First(&model).Error
	return admin
}

//Save : stores the changes occurred when creating or updating entries in the database
func (p *Repository) Save(admin Admin) Admin {
	p.DB.Save(&admin)

	return admin
}

// Delete : Deletes from the data base
func (p *Repository) Delete(admin Admin) {
	p.DB.Delete(&admin)
}
