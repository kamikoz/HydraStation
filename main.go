package main

import (
	"HYDRA/hlogger"
	"fmt"
	"net/http"
)

var logger = hlogger.GetInstance()

func main() {
	logger.Println("Starting Hydra web service")

	http.HandleFunc("/",sroot)
	http.ListenAndServe(":8080",nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger.Println("Received an HTTP request on root url")
	fmt.Fprint(w,"Welcome to Hydra Station")
}
