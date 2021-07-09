package httproutes

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Healthcheck okey!")
	w.WriteHeader(http.StatusOK)
}
