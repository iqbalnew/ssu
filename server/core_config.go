package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Listen address is an array of IP addresses and port combinations.
	// Listen address is an array so that this service can listen to many interfaces at once.
	// You can use this value for example: []string{"192.168.1.12:80", "25.49.25.73:80"} to listen to
	// listen to interfaces with IP address of 192.168.1.12 and 25.49.25.73, both on port 80.
	ListenAddress string `config:"LISTEN_ADDRESS"`

	CorsAllowedHeaders []string `config:"CORS_ALLOWED_HEADERS"`
	CorsAllowedMethods []string `config:"CORS_ALLOWED_METHODS"`
	CorsAllowedOrigins []string `config:"CORS_ALLOWED_ORIGINS"`

	JWTSecret string `config:"JWT_SECRET"`
	Dsn       string `config:"DB_DSN"`
}

var config *Config

func initConfig() {
	// Todo: add env checker

	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		log.Fatalf("Error loading .env file")
	}

	config = &Config{
		ListenAddress:      fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
		CorsAllowedHeaders: []string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Content-Disposition", "Origin", "X-Requested-With", "X-Forwarded-For"},
		CorsAllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "PUT"},
		CorsAllowedOrigins: []string{},
		JWTSecret:          os.Getenv("JWT_SECRET"),
		Dsn:                os.Getenv("DB_DSN"),
	}

}

func (c *Config) AsString() string {
	data, _ := json.Marshal(c)
	return string(data)
}
