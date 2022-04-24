package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Listen address is an array of IP addresses and port combinations.
	// Listen address is an array so that this service can listen to many interfaces at once.
	// You can use this value for example: []string{"192.168.1.12:80", "25.49.25.73:80"} to listen to
	// listen to interfaces with IP address of 192.168.1.12 and 25.49.25.73, both on port 80.
	Env string `json:"env"`

	ListenAddress string `config:"LISTEN_ADDRESS"`

	CorsAllowedHeaders []string `config:"CORS_ALLOWED_HEADERS"`
	CorsAllowedMethods []string `config:"CORS_ALLOWED_METHODS"`
	CorsAllowedOrigins []string `config:"CORS_ALLOWED_ORIGINS"`

	JWTSecret   string `config:"JWT_SECRET"`
	JWTDuration string `config:"JWT_DURATION"`
	Dsn         string `config:"DB_DSN"`

	AmqpURI string `config:"AMQP_URI`

	MongoURI string `config:"MONGO_URI"`

	LoggerPort string `config:"LOGGER_PORT"`
	LoggerHost string `config:"LOGGER_HOST"`
	LoggerTag  string `config:"LOGGER_TAG"`

	AnnouncementService string `config:"ANNOUNCEMENT_SERVICE"`

	// AccountService      string `config:"ACCOUNT_SERVICE"`
	// CompanyService      string `config:"COMPANY_SERVICE"`
}

var config *Config

func initConfig() {
	// Todo: add env checker

	godotenv.Load(".env")

	config = &Config{
		Env:                 getEnv("ENV", "DEV"),
		ListenAddress:       fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
		CorsAllowedHeaders:  []string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Content-Disposition", "Origin", "Content-Length", "Authorization", "ResponseType", "X-Requested-With", "X-Forwarded-For"},
		CorsAllowedMethods:  []string{"GET", "POST", "PATCH", "DELETE", "PUT"},
		CorsAllowedOrigins:  []string{"*"},
		JWTSecret:           getEnv("JWT_SECRET", "secret"),
		JWTDuration:         getEnv("JWT_DURATION", "48h"),
		Dsn:                 getEnv("DB_DSN", ""),
		AmqpURI:             getEnv("AMQP_URI", ""),
		AnnouncementService: getEnv("ANNOUNCEMENT_SERVICE", ":9091"),
		MongoURI:            getEnv("MONGO_URI", "mongodb://root:12345@mongo:27017/"),
		LoggerPort:          getEnv("LOGGER_PORT", ""),
		LoggerHost:          getEnv("LOGGER_HOST", ""),
		LoggerTag:           getEnv("LOGGER_TAG", ""),
		// AccountService:      os.Getenv("ACCOUNT_SERVICE"),
		// CompanyService:      os.Getenv("COMPANY_SERVICE"),
	}

}

func (c *Config) AsString() string {
	data, _ := json.Marshal(c)
	return string(data)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func logOutput() func() {
	logfile := `../logfile`
	// open file read/write | create if not exist | clear file at open if exists
	f, _ := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	// save existing stdout | MultiWriter writes to saved stdout and file
	out := os.Stdout
	mw := io.MultiWriter(out, f)

	// get pipe reader and writer | writes to pipe writer come out pipe reader
	r, w, _ := os.Pipe()

	// replace stdout,stderr with pipe writer | all writes to stdout, stderr will go through pipe instead (fmt.print, log)
	os.Stdout = w
	os.Stderr = w

	// writes with log.Print should also write to mw
	log.SetOutput(mw)

	//create channel to control exit | will block until all copies are finished
	exit := make(chan bool)

	go func() {
		fmt.Println("=================> start log")
		// copy all reads from pipe to multiwriter, which writes to stdout and file
		_, _ = io.Copy(mw, r)
		// when r or w is closed copy will finish and true will be sent to channel
		exit <- true
	}()

	// function to be deferred in main until program exits
	return func() {
		// close writer then block on exit channel | this will let mw finish writing before the program exits
		_ = w.Close()
		<-exit
		// close file after all writes have finished
		_ = f.Close()
	}

}
