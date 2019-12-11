package bank

//BankService : Is in charge of communicate the api with the repository
type BankService struct {
	BankRepository BankRepository
}

func ProvideBankService(p BankRepository) BankService {
	return BankService{BankRepository: p}
}

func (p *BankService) FindAll() []Bank {
	return p.BankRepository.FindAll()
}

func (p *BankService) FindByID(id uint) Bank {
	return p.BankRepository.FindByID(id)
}

func (p *BankService) Save(bank Bank) Bank {
	p.BankRepository.Save(bank)

	return bank
}

func (p *BankService) Delete(bank Bank) {
	p.BankRepository.Delete(bank)
}
