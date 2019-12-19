package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mercadolibre/hexagonthis/core"
	"net/http"
)

type HttpRepository struct {
}

func (db *HttpRepository) GetAllPlates() []core.PlateAndDriver {
	resp, _ := http.Get("http://localhost:3000")
	defer resp.Body.Close()
	var plateDriver []core.PlateAndDriver
	_ = json.NewDecoder(resp.Body).Decode(&plateDriver)
	return plateDriver
}

func (db *HttpRepository) SavePlate(pd core.PlateAndDriver) error {
	bytesRepresentation, _ := json.Marshal(pd)
	r, err := http.Post("http://localhost:3000/save", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return err
	}
	if r.Status != "200" {
		return errors.New("Fail to save plate to remote repo")
	}
	return nil
}

func HttpRepositoryFactory() *HttpRepository {
	return &HttpRepository{}
}