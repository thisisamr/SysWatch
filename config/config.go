package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Host string
	Port string
)

var Green = "\033[32m"
var Cyan = "\033[36m"
var Reset = "\033[0m"

func InitConfig() {
	// Load environment variables from the .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults or system environment variables.")
	}

	// Set Host
	H, hostSet := os.LookupEnv("HOST")
	if !hostSet {
		H = "localhost"
		log.Println("Host not set, defaulting to 'localhost'.")
	}
	Host = H

	// Set Port
	P, portSet := os.LookupEnv("PORT")
	if !portSet {
		P = "3000"
		log.Println("Port not set, defaulting to '3000'.")
	}
	Port = P
	log.Println(Cyan + `
 ▗▄▄▖▗▖  ▗▖▗▄▄▖▗▖ ▗▖ ▗▄▖▗▄▄▄▖▗▄▄▖▗▖ ▗▖    
▐▌    ▝▚▞▘▐▌   ▐▌ ▐▌▐▌ ▐▌ █ ▐▌   ▐▌ ▐▌    
 ▝▀▚▖  ▐▌  ▝▀▚▖▐▌ ▐▌▐▛▀▜▌ █ ▐▌   ▐▛▀▜▌    
▗▄▄▞▘  ▐▌ ▗▄▄▞▘▐▙█▟▌▐▌ ▐▌ █ ▝▚▄▄▖▐▌ ▐▌    
			` + Reset)
}
