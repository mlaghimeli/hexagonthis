package test

import (
	"errors"
	"github.com/mercadolibre/hexagonthis/repository"
	"testing"

	"github.com/mercadolibre/hexagonthis/core"

	"github.com/stretchr/testify/assert"
)

func TestShould_give_all_car_platesAndDrivers_when_asked_with_support_of_PlatesAndDriversRepository(t *testing.T) {
	var pr core.PlatesAndDriversRepository = &PlatesRepositoryMock{}
	r := core.PlatesAndDriversServiceFactory(pr)
	plates := r.GetAllRegisteredPlatesAndDrivers()
	assert.NotNil(t, plates)
	assert.NotEmpty(t, plates)
	assert.EqualValues(t, plates, []core.PlateAndDriver{{"MML-0626", "Murilo", "99999999"}, {"MJL-6864", "Gisele", "33333333"}})
}

func TestShould_succesfully_save_plateAndDriver_when_asked_with_support_of_PlatesAndDriversRepository(t *testing.T) {
	var pr core.PlatesAndDriversRepository = &PlatesRepositoryMock{}
	r := core.PlatesAndDriversServiceFactory(pr)
	err := r.RegisterPlateAndDriver(core.PlateAndDriver{})
	assert.Nil(t, err)
}

func TestShould_give_error_saving_plateAndDriver_when_repo_full_with_support_of_PlatesAndDriversRepository(t *testing.T) {
	var pr core.PlatesAndDriversRepository = &PlatesRepositoryMock{errors.New("Repo full")}
	r := core.PlatesAndDriversServiceFactory(pr)
	err := r.RegisterPlateAndDriver(core.PlateAndDriver{})
	assert.Error(t, err)
	assert.EqualError(t, err,"Repo full")
}


func TestShould_save_and_retrieve_plateAndDriver_with_support_of_PlatesAndDriversInMemoryDB(t *testing.T) {
	var pr core.PlatesAndDriversRepository = repository.InMemoryDBFactory(2)
	r := core.PlatesAndDriversServiceFactory(pr)
	err := r.RegisterPlateAndDriver(core.PlateAndDriver{})
	assert.Error(t, err)
	assert.EqualError(t, err,"Repo full")
}

type PlatesRepositoryMock struct {
	err error
}

func (p *PlatesRepositoryMock) GetAllPlates() []core.PlateAndDriver {
	return []core.PlateAndDriver{{"MML-0626", "Murilo", "99999999"}, {"MJL-6864", "Gisele", "33333333"}}
}

func (p *PlatesRepositoryMock) SavePlate(driver core.PlateAndDriver) error{
	return p.err
}