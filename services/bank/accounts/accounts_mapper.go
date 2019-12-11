package accounts

//ToAccount : mapper single objectDTO to gorm struct
func ToAccount(fullAccountDTO FullAccountDTO) BankAccount {
	return BankAccount{
		AccountNumber: fullAccountDTO.AccountNumber,
		Balance:       fullAccountDTO.Balance,
		Bank:          fullAccountDTO.Bank,
		Owner:         fullAccountDTO.Owner,
	}
}

/*ToFullAccountDTO : mapper to gorm struct to single object DTO */
func ToFullAccountDTO(bankAccount BankAccount) FullAccountDTO {
	return FullAccountDTO{
		AccountNumber: bankAccount.AccountNumber,
		Balance:       bankAccount.Balance,
		Bank:          bankAccount.Bank,
		Owner:         bankAccount.Owner,
	}
}

/*ToLoggedUserDTO : mapper to gorm struct to single object DTO */
func ToLoggedUserDTO(bankAccount BankAccount, token string) LoggedUserDTO {
	return LoggedUserDTO{
		Owner: bankAccount.Owner,
		Token: token,
	}
}

//ToFullAccountDTOs : Mapper to gorm stuct to array of DTO objects
func ToFullAccountDTOs(bankAccount []BankAccount) []FullAccountDTO {
	bankAccountsDtos := make([]FullAccountDTO, len(bankAccount))

	for i, itm := range bankAccount {
		bankAccountsDtos[i] = ToFullAccountDTO(itm)
	}

	return bankAccountsDtos
}
