package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"middlewares"
	"models"
	"time"
)

type EndOfWorkResponse struct {
	Status     string `json:"status"`
	User       models.User
	WorkRecord models.WorkRecord
}

func EndOfWork() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		auth := c.Get("auth").(*models.User)

		user := models.User{}
		dbs.Transaction.Table("users").Where(models.User{ID: auth.ID}).First(&user)

		jst, _ := time.LoadLocation("Asia/Tokyo")
		date := time.Now().In(jst)
		today := date.Format("2006-01-02")
		currentTime := date.Format("15:04")

		workrecord := models.WorkRecord{}
		if dbs.Transaction.Where(models.WorkRecord{UserId: user.ID, Date: today}).First(&workrecord).RecordNotFound() {
			workrecord = models.WorkRecord{UserId: user.ID, Date: today, EndOfWork: currentTime}
			dbs.Transaction.Create(&workrecord)
		} else {
			dbs.Transaction.Model(&workrecord).UpdateColumns(models.WorkRecord{EndOfWork: currentTime})
		}

		response := EndOfWorkResponse{Status: "SUCCESS", User: user, WorkRecord: workrecord}

		return c.JSON(fasthttp.StatusOK, response)
	}
}
