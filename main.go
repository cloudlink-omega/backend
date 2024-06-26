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

	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		panic(err)
	}

	authlessMode, err := strconv.ParseBool(os.Getenv("AUTHLESS_MODE"))
	if err != nil {
		panic(err)
	}

	useInMemoryClientMgr, err := strconv.ParseBool(os.Getenv("USE_IN_MEMORY_CLIENT_MGR"))
	if err != nil {
		panic(err)
	}

	enableEmail, err := strconv.ParseBool(os.Getenv("ENABLE_EMAIL"))
	if err != nil {
		panic(err)
	}

	emailPort, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	if err != nil {
		panic(err)
	}

	// Initialize data manager
	mgr := dm.New(

		/*
			SERVER_NICKNAME: Specifies a global nickname for the server; You should NOT change this after public deployment!

			The desired format for a server nickname is as follows:

			[ISO 3166-1 alpha-3 country code]-[A one-word name]-[numerical instance number]

			e.g. "USA-Omega-1".
		*/
		os.Getenv("SERVER_NICKNAME"),

		/*
			SQL_DRIVER: Specifies the SQL driver to use for the database. Default is "mysql".
			You will need to modify this file to match the SQL driver you want to use.
		*/
		os.Getenv("SQL_DRIVER"),

		// Change this to SQL-driver specific connection string.
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DATABASE"),
		),

		/*
			AUTHLESS_MODE: Specifies if the server should be in Authless Mode.

			By default, this should be set to false, and you should connect to a SQL database for the server.
			However, some implementations may act as a standalone signaling server. In this case, set this to true
			to completely ignore authentication.

			This is intended for standalone or development environments. Note that using authless mode renders your
			server vulnerable to connection spamming or spoofed user accounts.
		*/
		authlessMode,

		/*
			USE_IN_MEMORY_CLIENT_MGR: Specifies if the server should use a KeyDB server for client management.

			By default, this should be set to true. However, if you are using Authless Mode, you should set this to false.
			Note that using the built-in client manager will severely harm performance.

			This is intended for standalone or development environments.
		*/
		useInMemoryClientMgr,

		// Specify a boolean value if you want to enable email sending on the server.
		enableEmail,

		// Change this to desired outgoing email server port (e.g. 587)
		emailPort,

		// Change this to desired outgoing email server address (e.g. smtp.gmail.com)
		os.Getenv("EMAIL_SERVER"),

		// Specify your email username here (e.g. someone@example.com)
		os.Getenv("EMAIL_USERNAME"),

		// Specify your email password here (Use an app password if you have multifactor enabled)
		os.Getenv("EMAIL_PASSWORD"),

		// Specify the server's public hostname here. (Used for magic links)
		os.Getenv("SERVER_PUBLIC_HOSTNAME"),
	)

	// Run the server
	api.RunServer(
		os.Getenv("API_HOST"),
		apiPort,
		mgr,
	)
}
