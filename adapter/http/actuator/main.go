package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody bla bla
type HealthBody struct {
	Status string `json:"status"`
}

// Health  ddddeff
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// HealthBody blabla
	var status = HealthBody{Status: "alive"}

	json.NewEncoder(w).Encode(status)

}
