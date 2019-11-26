//+build wireinject

package main

import (
	"Go_Gingonic_Server/bank"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initBankAPI(db *gorm.DB) bank.BankAPI {
	wire.Build(bank.ProvideBankAPI, bank.ProvideBankService, bank.ProvideBankRepostiory)

	return bank.BankAPI{}
}

/* El propósito de este fichero es proporcionar información sobre qué proveedores usar para construir un Evento y,
por lo tanto, lo excluiremos de nuestro binario final con la restricción de compilación en la parte superior del archivo.

Se nos generara un fichero llamado wire_gen.go, ese sera el Evento.

*/
