package main

import (
	"backend/src/config"
	"backend/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Social Backend Starting")

	config.LoadVariables()

	fmt.Println("Environment Variables Loaded with Success")

	r := router.Generate()

	fmt.Printf("DB Connected to Port : %d", config.PortDB)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PortDB), r))

}
