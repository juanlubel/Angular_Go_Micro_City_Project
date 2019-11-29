package accounts

//AccountService : Is in charge of communicate the api with the repository
type AccountService struct {
	AccountsRepository AccountsRepository
}

func ProvideAccountService(p AccountsRepository) AccountService {
	return AccountService{AccountsRepository: p}
}

func (p *AccountService) FindAll() []BankAccount {
	return p.AccountsRepository.FindAll()
}

func (p *AccountService) FindByBankAndName(owner string, bank string) BankAccount {
	return p.AccountsRepository.FindByBankAndName(owner, bank)
}
func (p *AccountService) FindByOwner(owner string) BankAccount {
	return p.AccountsRepository.FindByOwner(owner)
}

func (p *AccountService) Save(bankAccount BankAccount) BankAccount {
	p.AccountsRepository.Save(bankAccount)

	return bankAccount
}

func (p *AccountService) Delete(bankAccount BankAccount) {
	p.AccountsRepository.Delete(bankAccount)
}
