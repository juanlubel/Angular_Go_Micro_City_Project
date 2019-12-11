package bank

//ToBank : mapper single objectDTO to gorm struct
func ToBank(bankDTO BankDTO) Bank {
	return Bank{BankName: bankDTO.BankName}
}

/*ToBankDTO : mapper to gorm struct to single object DTO */
func ToBankDTO(bank Bank) BankDTO {
	return BankDTO{ID: bank.ID, BankName: bank.BankName}
}

//ToBankDTOs : Mapper to gorm stuct to array of DTO objects
func ToBankDTOs(banks []Bank) []BankDTO {
	bankdtos := make([]BankDTO, len(banks))

	for i, itm := range banks {
		bankdtos[i] = ToBankDTO(itm)
	}

	return bankdtos
}
