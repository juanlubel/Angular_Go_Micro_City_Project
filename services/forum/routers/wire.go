//+build wireinject

package router

import (
	"Go_Gingonic_Server/topics"
	"Go_Gingonic_Server/utils"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initTopicsAPI(db *gorm.DB, jwt *utils.JWT) topics.TopicAPI {
	wire.Build(topics.ProvideTopicAPI, topics.ProvideTopicService, topics.ProvideTopicRepostiory)

	return topics.TopicAPI{}
}

/* El propósito de este fichero es proporcionar información sobre qué proveedores usar para construir un Evento y,
por lo tanto, lo excluiremos de nuestro binario final con la restricción de compilación en la parte superior del archivo.

Se nos generara un fichero llamado wire_gen.go, ese sera el Evento.

*/
