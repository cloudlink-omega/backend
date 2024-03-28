package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	api "github.com/cloudlink-omega/backend/pkg/api"
	dm "github.com/cloudlink-omega/backend/pkg/data"
	godotenv "github.com/joho/godotenv"

	/*
		Before you run the server, you will need to manually specify which SQL driver you want to use.
		CL Omega is natively developed using "mysql" and "sqlite".
	*/

	_ "github.com/go-sql-driver/mysql"
	// _ "modernc.org/sqlite"
)

func main() {
	/*
		// Initialize data manager
		mgr := dm.New(
			"sqlite",
			"file:./test.db?_pragma=foreign_keys(1)", // Use SQLite for testing/development purposes only
		)*/

	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	// Get API port
	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		panic(err)
	}

	// Initialize data manager
	mgr := dm.New(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DATABASE"),
		),
	)

	// Run the server
	api.RunServer(
		os.Getenv("API_HOST"),
		apiPort,
		mgr,
	)
}
