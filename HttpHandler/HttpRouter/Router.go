package httprouter

import (
	"fmt"

	"github.com/gorilla/mux"
)

func CreateRouter() (string, *mux.Router) {

	port := fmt.Sprintf(":%s", "3000")
	router := mux.NewRouter().StrictSlash(true)

	fmt.Println("Starting router...")

	return port, router

}
