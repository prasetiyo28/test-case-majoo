package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type Person struct {
	gorm.Model
	First_Name string
	Last_Name  string
}

func DBinit() *gorm.DB {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlHost := os.Getenv("MYSQL_HOSTNAME")
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	mysqlConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase,
	)

	db, err := gorm.Open("mysql", mysqlConfig)
	if err != nil {
		fmt.Println("eroroorororo", err)
	}
	db.SingularTable(true)

	return db
}
