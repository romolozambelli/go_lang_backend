package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// String to connect to MySql
	StringConnectionDb = ""

	// Port of the MySql connection
	PortDB = 0
)

// Initialize environment variables
func LoadVariables() {

	fmt.Println("Loading environment variables ...")

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	PortDB, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		PortDB = 9000
	}

	StringConnectionDb = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

}
