package accounts

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql
)

// AccountsRepository : obtain the context of the database
type AccountsRepository struct {
	DB *gorm.DB
}

// ProvideAccountsRepostiory : provides the repository to use the db context
func ProvideAccountsRepostiory(DB *gorm.DB) AccountsRepository {
	return AccountsRepository{DB: DB}
}

//FindAll : obtain from db all the data of the indicated gorm model
func (p *AccountsRepository) FindAll() []BankAccount {
	var bankAccount []BankAccount
	p.DB.Find(&bankAccount)

	return bankAccount
}

func (p *AccountsRepository) FindByBankAndName(owner string, bank string) BankAccount {
	var bankAccount BankAccount
	println(owner)
	p.DB.Where("Owner = ? AND Bank=?", owner, bank).First(&bankAccount)

	return bankAccount
}

//FindByOwner :  obtain from db the data that contains the param name
func (p *AccountsRepository) FindByOwner(owner string) BankAccount {
	var bankAccount BankAccount
	println(owner)
	p.DB.Where("Owner = ?", owner).First(&bankAccount)

	return bankAccount
}

//Save : stores the changes occurred when creating or updating entries in the database
func (p *AccountsRepository) Save(bankAccount BankAccount) BankAccount {
	p.DB.Save(&bankAccount)

	return bankAccount
}

// Delete : Deletes from the data base
func (p *AccountsRepository) Delete(bankAccount BankAccount) {
	p.DB.Delete(&bankAccount)
}
