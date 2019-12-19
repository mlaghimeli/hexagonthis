package repository

import (
	"errors"
	"github.com/mercadolibre/hexagonthis/core"
)

type InMemoryDB struct {
	maxLimit int
	platesAndDrivers []core.PlateAndDriver
}

func (db *InMemoryDB) GetAllPlates() []core.PlateAndDriver {
	return db.platesAndDrivers
}

func (db *InMemoryDB) SavePlate(pd core.PlateAndDriver) error {
	if len(db.platesAndDrivers) == db.maxLimit {
		return errors.New("Fail to save plate to, memory repo full")
	}
	db.platesAndDrivers = append(db.platesAndDrivers, pd)
	return nil
}

func InMemoryDBFactory(maxLimit int) *InMemoryDB {
	return &InMemoryDB{maxLimit, make([]core.PlateAndDriver, 0)}
}