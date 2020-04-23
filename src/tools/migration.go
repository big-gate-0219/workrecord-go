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
	user1 := models.User{UID: "test1@mail.com", MailAddress: "test1@mail.com", Name: "テスト１", Password: "password"}
	user2 := models.User{UID: "test2@mail.com", MailAddress: "test2@mail.com", Name: "テスト２", Password: "password"}
	user3 := models.User{UID: "test3@mail.com", MailAddress: "test3@mail.com", Name: "テスト３", Password: "password"}
	user4 := models.User{UID: "test4@mail.com", MailAddress: "test4@mail.com", Name: "テスト４", Password: "password"}
	user5 := models.User{UID: "test5@mail.com", MailAddress: "test5@mail.com", Name: "テスト５", Password: "password"}
	user6 := models.User{UID: "test6@mail.com", MailAddress: "test6@mail.com", Name: "テスト６", Password: "password"}
	user7 := models.User{UID: "test7@mail.com", MailAddress: "test7@mail.com", Name: "テスト７", Password: "password"}
	db.Debug().Create(&user1)
	db.Debug().Create(&user2)
	db.Debug().Create(&user3)
	db.Debug().Create(&user4)
	db.Debug().Create(&user5)
	db.Debug().Create(&user6)
	db.Debug().Create(&user7)

	db.Debug().AutoMigrate(&models.Group{})
	db.Debug().Unscoped().Delete(&models.Group{})
	group1 := models.Group{Name: "Group-1"}
	group2 := models.Group{Name: "Group-2"}
	group3 := models.Group{Name: "Group-3"}
	db.Debug().Create(&group1)
	db.Debug().Create(&group2)
	db.Debug().Create(&group3)

	db.Debug().AutoMigrate(&models.GroupUser{})
	db.Debug().Unscoped().Delete(&models.Group{})
	groupUser1_1 := models.GroupUser{GroupId:group1.ID, UserId:user1.ID}
	groupUser1_2 := models.GroupUser{GroupId:group1.ID, UserId:user2.ID}
	groupUser1_4 := models.GroupUser{GroupId:group1.ID, UserId:user4.ID}
	groupUser2_1 := models.GroupUser{GroupId:group2.ID, UserId:user1.ID}
	db.Debug().Create(&groupUser1_1)
	db.Debug().Create(&groupUser1_2)
	db.Debug().Create(&groupUser1_4)
	db.Debug().Create(&groupUser2_1)


	
	
	db.Debug().AutoMigrate(&models.WorkRecord{})


}
