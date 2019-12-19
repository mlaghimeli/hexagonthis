package core

type service struct {
	plateRepository PlatesAndDriversRepository
}

func (s *service) GetAllRegisteredPlatesAndDrivers() []PlateAndDriver {
	// all necessary business logic like validations and so on
	return s.plateRepository.GetAllPlates()
}

func (s *service) RegisterPlateAndDriver(pd PlateAndDriver) error {
	// all necessary business logic like validations and so on
	return s.plateRepository.SavePlate(pd)
}

func PlatesAndDriversServiceFactory(pr PlatesAndDriversRepository) RequestPlatesAndDriversService {
	return &service{pr}
}
