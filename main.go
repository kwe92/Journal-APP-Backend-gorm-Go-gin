package main

import (
	"fmt"
	"journal_api/router"
	"journal_api/utility"
)

// TODO: Review function / method comments

const address = ":8080"

func main() {
	utility.LoadEnv()

	utility.LoadDatabase()

	router := router.SetupRouter()

	router.Run(address)

	fmt.Println("Server running on port 8000")

}

// *gorm.DB.AutoMigrate(pointer_to_a_struct_that_will_be_a_table)

//   - automatic schema migration for a given struct
//   - will create the table and column names if they don't exist

// Loading Environment Variables

//   - environment variables must be set or loaded in order to be used

// godotenv.Load(.env_file)

//   - used to load your environment variables
//   - should be called at the top of main

// Gin Web Framework

//   - Build API's simply and quickly
//   - easy request parsing and validation
//   - provides the ability to group endpoints
//   - provides the ability to add middleware to groups of endpoints
