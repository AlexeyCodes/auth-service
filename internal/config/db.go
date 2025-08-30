package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

//variable responsible for pool of connections
var DB *pgxpool.Pool


//load db
func LoadDB(){

	//load .env file
	err := godotenv.Load()

	//error handling if .env not found
	if err != nil {
		log.Print(".env not found, system env loaded")
	}
	
	//load env variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	//create connection to db
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, name)

	//create new local pool
	pool, err := pgxpool.New(context.Background(), connection)

	//eror handling if unable to connect to db
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	//create context
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	//guarantee it would be completed if something goes wrong
	defer cancel()

	//ping db
	err = pool.Ping(ctx)

	//error handling if cannot ping db
	if err != nil {
		log.Fatalf("Cannot ping db: %v", err)
	}

	//Successfull conect
	log.Println("Connected successfully")
	
	//appropiate result of locsl pool to global db var
	DB = pool
}