package httphandler

import (
	"log"
	"net/http"

	httprouter "github.com/AlejandroWaiz/mux-router/HttpHandler/HttpRouter"
	httproutes "github.com/AlejandroWaiz/mux-router/HttpHandler/HttpRoutes"
)

func StartRouter() {

	port, router := httprouter.CreateRouter()

	router.HandleFunc("/home", httproutes.HomePage)
	router.HandleFunc("/hc", httproutes.HealthCheck)

	log.Fatal(http.ListenAndServe(port, router))

}
