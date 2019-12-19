package core

type PlatesAndDriversRepository interface {
	GetAllPlates() []PlateAndDriver
	SavePlate(PlateAndDriver) error
}
