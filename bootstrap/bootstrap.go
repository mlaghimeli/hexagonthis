package bootstrap

import (
	"github.com/mercadolibre/hexagonthis/core"
	"github.com/mercadolibre/hexagonthis/repository"
	"github.com/mercadolibre/hexagonthis/rest"
)

func SetupConfigurableDependencies() {

	//var platesRepository core.PlatesAndDriversRepository = repository.InMemoryDBFactory(3)
	var platesRepository core.PlatesAndDriversRepository = repository.ExternalApplicationRepositoryFactory()
	hexagon := core.PlatesAndDriversServiceFactory(platesRepository)
	restAdapter := rest.RestControllerFactory(hexagon)
	restAdapter.Bind()

}