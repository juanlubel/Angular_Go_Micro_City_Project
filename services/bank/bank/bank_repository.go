package bank

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql
)

// BankRepository : obtain the context of the database
type BankRepository struct {
	DB *gorm.DB
}

// ProvideBankRepostiory : provides the repository to use the db context
func ProvideBankRepostiory(DB *gorm.DB) BankRepository {
	return BankRepository{DB: DB}
}

//FindAll : obtain from db all the data of the indicated gorm model
func (p *BankRepository) FindAll() []Bank {
	var banks []Bank
	p.DB.Find(&banks)

	return banks
}

//FindByID :  obtain from db the data that contains the param id
func (p *BankRepository) FindByID(id uint) Bank {
	var bank Bank
	p.DB.First(&bank, id)

	return bank
}

//Save : stores the changes occurred when creating or updating entries in the database
func (p *BankRepository) Save(bank Bank) Bank {
	p.DB.Save(&bank)

	return bank
}

// Delete : Deletes from the data base
func (p *BankRepository) Delete(bank Bank) {
	p.DB.Delete(&bank)
}
