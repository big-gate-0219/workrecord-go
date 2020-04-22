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
	db.Debug().Unscoped().Delete(&models.User{})
	db.Debug().Create(&models.User{UID: "test1@mail.com", MailAddress: "test1@mail.com", Name: "テスト１", Password: "password"})
	db.Debug().Create(&models.User{UID: "test2@mail.com", MailAddress: "test2@mail.com", Name: "テスト２", Password: "password"})
	db.Debug().Create(&models.User{UID: "test3@mail.com", MailAddress: "test3@mail.com", Name: "テスト３", Password: "password"})
	db.Debug().Create(&models.User{UID: "test4@mail.com", MailAddress: "test4@mail.com", Name: "テスト４", Password: "password"})
	db.Debug().Create(&models.User{UID: "test5@mail.com", MailAddress: "test5@mail.com", Name: "テスト５", Password: "password"})
	db.Debug().Create(&models.User{UID: "test6@mail.com", MailAddress: "test6@mail.com", Name: "テスト６", Password: "password"})
	db.Debug().Create(&models.User{UID: "test7@mail.com", MailAddress: "test7@mail.com", Name: "テスト７", Password: "password"})

	db.Debug().AutoMigrate(&models.WorkRecord{})
	db.Debug().AutoMigrate(&models.Group{})
	db.Debug().AutoMigrate(&models.GroupUser{})

}
