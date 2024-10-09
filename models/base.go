package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func initEnv() {
	// Load values from .env into the system.
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value.
	// Outputs a panic if the value is missing.
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panic("Value " + v + "does not exist")
	}
	return value
}

func initDB() {
	username := getEnvValue("db_user")
	password := getEnvValue("db_pass")
	dbName := getEnvValue("db_name")
	dbHost := getEnvValue("db_host")
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	log.Println(dbUri)

	if conn, err := gorm.Open("postgres", dbUri); err != nil {
		log.Println(err)
	} else {
		db = conn
		db.Debug().AutoMigrate(&Account{}, &Contact{})
	}
}

// Returns the DB object descriptor.
func GetDB() *gorm.DB {
	return db
}
