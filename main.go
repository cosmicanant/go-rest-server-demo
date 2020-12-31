package main

import (
	"net/http"

	"github.com/cosmicanant/go-sample-app/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(`:3000`, nil)
}
