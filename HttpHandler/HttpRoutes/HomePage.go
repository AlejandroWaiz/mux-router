package httproutes

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Alejandro Waiz's example. Go programmer.")
}
