package databases

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Connect() (db *gorm.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	db, err = gorm.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logrus.Fatal(err)
	}

	return db, err
}
