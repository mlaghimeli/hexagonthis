package core

type RequestPlatesAndDriversService interface {
	GetAllRegisteredPlatesAndDrivers() []PlateAndDriver
	RegisterPlateAndDriver(PlateAndDriver) error
}
