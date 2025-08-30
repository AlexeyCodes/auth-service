package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

//Run migrations   pool of connections with db to create only one connection, migrations dir
func RunMigrations(pool *pgxpool.Pool, migrationsDir string) { 
	
	//find all files with .sql in dir
	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))

	//error handling if failed to read migrations
	if err != nil {
		log.Fatalf("Faied to read migrations: %v", err)
	}

	//go throught all files in dir with .sql
	for _, file := range files {
		//read from file
		content, err := os.ReadFile(file)

		//erorr hadling if failed to read file
		if err != nil {
			log.Fatalf("Failed to read file %s: %v", file, err)
		}

		//create context
		ctx := context.Background()

		//make sql requesrt from file
		_, err = pool.Exec(ctx, string(content))
		
		//error hndling if unable to read migration from file 
		if err != nil {
			log.Fatalf("Fail to launch migration %s: %v", file, err)
		}

		//log successfully aplied migraation from file
		fmt.Printf("Migration %s applied successfully\n", file)
	}
}
