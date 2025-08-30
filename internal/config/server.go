package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

//create server config
type ServerConfig struct {
	Port string
	ReadTimeout time.Duration 
	WriteTimeout time.Duration
	IdleTimeout time.Duration
}

//load server config
func LoadServerConfig() *ServerConfig {
	//load env from .env
	err := godotenv.Load()

	//error handling if .env not found
	if err != nil {
		log.Print(".env not found, system env loaded")
	}

	//get env port from .env
	port := os.Getenv("PORT")

	//if .env not loaded or port env is empty it will be default 
	if port == ""{
		port = "8080"
	}

	readtimeout := 3 * time.Second  //maximum for read response from client
	writetimeout := 5 * time.Second //maximum for send response to client
	idletimeout :=  10* time.Minute //time of waiting for inactive connection

	//return config
	return &ServerConfig{
		Port : port,
		ReadTimeout: readtimeout,
		WriteTimeout: writetimeout,
		IdleTimeout: idletimeout,
	}
}