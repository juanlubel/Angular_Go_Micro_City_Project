//+build wireinject

package router

import (
	"Go_Gingonic_Server/accounts"
	banks "Go_Gingonic_Server/bank"
	"Go_Gingonic_Server/utils"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initBankAPI(db *gorm.DB) banks.BankAPI {
	wire.Build(banks.ProvideBankAPI, banks.ProvideBankService, banks.ProvideBankRepostiory)

	return banks.BankAPI{}
}

func initAccountsAPI(db *gorm.DB, jwt *utils.JWT) accounts.AccountAPI {
	wire.Build(accounts.ProvideAccountAPI, accounts.ProvideAccountService, accounts.ProvideAccountsRepostiory)

	return accounts.AccountAPI{}
}

/* El propósito de este fichero es proporcionar información sobre qué proveedores usar para construir un Evento y,
por lo tanto, lo excluiremos de nuestro binario final con la restricción de compilación en la parte superior del archivo.

Se nos generara un fichero llamado wire_gen.go, ese sera el Evento.

*/
