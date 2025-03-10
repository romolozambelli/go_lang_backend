package main

import (
	"backend/src/config"
	"backend/src/router"
	"fmt"
	"log"
	"net/http"
)

// const defaultMessage = "Hello Default Message!"
// const newWelcomeMessage = "Hello, welcome to this OpenFeature-enabled website!"

func main() {

	config.LoadVariables()

	// err := openfeature.SetProviderAndWait(flagd.NewProvider())
	// if err != nil {
	// 	// If a provider initialization error occurs, log it and exit
	// 	log.Fatalf("Failed to set the OpenFeature provider: %v", err)
	// }

	// Initialize OpenFeature client
	//client := openfeature.NewClient("GoStartApp")

	// Evaluate welcome-message feature flag
	// welcomeMessage, _ := client.BooleanValue(
	// 	context.Background(), "welcome-message", false, openfeature.EvaluationContext{},
	// )

	// if welcomeMessage {
	// 	fmt.Println(newWelcomeMessage)
	// } else {
	// 	fmt.Println(defaultMessage)
	// }

	r := router.Generate()

	fmt.Printf("Application Started. Listenning HTTP requests\n")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PortDB), r))
}
