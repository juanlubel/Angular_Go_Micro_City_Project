//+build wireinject

package main

import (
	"Go_Gingonic_Server/admin"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitAdminAPI(db *gorm.DB) admin.Api {
	wire.Build(admin.ProvideAdminRepostiory, admin.ProvideAdminService, admin.ProvideAdminApi)

	return admin.Api{}
}

/* El propósito de este fichero es proporcionar
información sobre qué proveedores usar para construir un Evento y,
por lo tanto, lo excluiremos de nuestro binario final con la restricción de compilación en la parte superior del archivo.

Se nos generara un fichero llamado wire_gen.go, ese sera el Evento.

*/
