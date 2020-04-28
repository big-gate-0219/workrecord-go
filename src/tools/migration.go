package main

import (
	"databases"
	"models"

	"github.com/sirupsen/logrus"
)

func main() {
	db, err := databases.Connect()
	defer db.Close()
	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Group{})
	db.Debug().AutoMigrate(&models.GroupUser{})
	db.Debug().AutoMigrate(&models.WorkRecord{})

}
