package dbconnection

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func DbConn() {
	err := godotenv.Load("dbConfig.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	localhost := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	sslmode := os.Getenv("SSLMODE")
	dbname := os.Getenv("DBNAME")
	connectStr := ("host=" + localhost + " port=" + port + " user=" + user + " password=" + password + " sslmode=" + sslmode + " dbname=" + dbname)
	DB, err = gorm.Open("postgres", connectStr)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("connection successfully")
}
