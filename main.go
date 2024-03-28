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

	/*
		Authless Mode

		This mode completely removes the need for a SQL database,
		and the server will accept any login request and generate
		randomized auth tokens.

		To use this mode, set AUTHLESS_MODE=true in the .env file.

		This is intended for standalone or development environments.
		By default, this should be set to false.
	*/
	authlessMode, err := strconv.ParseBool(os.Getenv("AUTHLESS_MODE"))
	if err != nil {
		panic(err)
	}

	/*
		Use in-memory client manager

		This mode disables the use of a KeyDB server for client management.
		Be aware that this mode is not memory-efficient, and can cause performance issues.

		To use an external client management system (e.g. KeyDB), set USE_IN_MEMORY_CLIENT_MGR=false in the .env file.

		This is intended for standalone or development environments.
		By default, this is set to true (Note: Server does not have any code to use KeyDB at this time!)
	*/
	useInMemoryClientMgr, err := strconv.ParseBool(os.Getenv("USE_IN_MEMORY_CLIENT_MGR"))
	if err != nil {
		panic(err)
	}

	// Initialize data manager
	mgr := dm.New(
		"mysql", // Change this to desired SQL driver
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DATABASE"),
		), // Change this to SQL-driver specific connection string
		authlessMode,
		useInMemoryClientMgr,
	)

	// Run the server
	api.RunServer(
		os.Getenv("API_HOST"),
		apiPort,
		mgr,
	)
}
