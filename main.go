package main

import (
	"backend/src/config"
	"backend/src/router"
	"fmt"
	"log"
	"net/http"
)

// Init function used only to create the secret key
// func init() {
// 	key := make([]byte, 64)

// 	if _, erro := rand.Read(key); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString((key))

// 	fmt.Println(stringBase64)

// }

func main() {
	fmt.Printf("Social Backend Starting\n")

	config.LoadVariables()

	fmt.Printf("Environment Variables Loaded with Success\n")

	r := router.Generate()

	fmt.Printf("DB Connected to Port : %d \n", config.PortDB)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PortDB), r))

}
