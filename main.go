package main

import (
	"backend/src/config"
	"backend/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Social Backend Starting\n")

	config.LoadVariables()

	fmt.Printf("Environment Variables Loaded with Success\n")

	r := router.Generate()

	fmt.Printf("DB Connected to Port : %d \n", config.PortDB)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PortDB), r))

}
