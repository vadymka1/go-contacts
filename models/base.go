package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB // db

func init() {
	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)

	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)

	// conn, err := gorm.Open("postgres", "host=localhost port=5432 user=electropk dbname=go_crud password=password")

	if err != nil {
		fmt.Print(err)
	}

	db = conn

	db.Debug().AutoMigrate(&Account{}, &Contact{})

}

/*
	Get db connection
**/
func GetDb() *gorm.DB {
	return db
}
