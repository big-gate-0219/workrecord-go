package api

import (
	"strconv"
	"time"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"models"
	"middlewares"
)

type GetWorkrecordGroupTodayResponse struct {
	Status string `json:"status"`
	Date string `json:"date"`
	Group models.Group `json:"group"`
	UserWorkrecords []GetWorkrecordGroupTodayResponsePart `json:"user_workrecords"`
}

type GetWorkrecordGroupTodayResponsePart struct {
	User       models.User
	WorkRecord models.WorkRecord
}
func GetWorkrecordGroupToday() echo.HandlerFunc {
	return func(c echo.Context) error {
		groupId, _ := strconv.ParseUint(c.Param("group_id"), 10, 64)

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		time.Local = time.FixedZone("Local", 9*60*60)
		jst, _ := time.LoadLocation("Local")
		today := time.Now().In(jst).Format("2006-01-02")


		group := models.Group{}
		dbs.Transaction.Where(models.Group{ID: groupId}).Find(&group)


		groupUsers := []models.GroupUser{}
		dbs.Transaction.Where(models.GroupUser{GroupId: groupId}).Find(&groupUsers)

		userWorkrecords := []GetWorkrecordGroupTodayResponsePart{};
		for _, groupUser := range groupUsers {
			user := models.User{}
			dbs.Transaction.Where(models.User{ID: groupUser.UserId}).First(&user)

			workRecord := models.WorkRecord{}
			dbs.Transaction.Where(models.WorkRecord{UserId: user.ID, Date: today}).First(&workRecord)
			userWorkrecord := GetWorkrecordGroupTodayResponsePart{User: user, WorkRecord: workRecord}
			userWorkrecords = append(userWorkrecords, userWorkrecord)
		}

		response := GetWorkrecordGroupTodayResponse{
			Status: "SUCCESS",
			Date: today,
			Group: group,
			UserWorkrecords: userWorkrecords,
		}
		return c.JSON(fasthttp.StatusOK, response)
	}
}