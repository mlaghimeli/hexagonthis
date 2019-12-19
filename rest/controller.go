package rest

import (
	"encoding/json"
	"net/http"

	"github.com/mercadolibre/hexagonthis/core"

	"github.com/gorilla/mux"
)

func (rc *RestController) Bind() {
	router := mux.NewRouter()
	router.HandleFunc("/platesAndDrivers", rc.GetAll).Methods("GET")
	router.HandleFunc("/platesAndDrivers", rc.SavePlate).Methods("POST")
	http.ListenAndServe(":8000", router)
}

func (rc *RestController) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rc.requestPlatesAndDriversService.GetAllRegisteredPlatesAndDrivers())
}

func (rc *RestController) SavePlate(w http.ResponseWriter, r *http.Request) {
	var plateAndDriver core.PlateAndDriver
	_ = json.NewDecoder(r.Body).Decode(&plateAndDriver)
	err := rc.requestPlatesAndDriversService.RegisterPlateAndDriver(plateAndDriver)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("not ok"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

type RestController struct {
	requestPlatesAndDriversService core.RequestPlatesAndDriversService
}

func RestControllerFactory(rps core.RequestPlatesAndDriversService) *RestController {
	return &RestController{rps}
}