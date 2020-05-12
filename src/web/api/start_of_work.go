package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"middlewares"
	"models"
	"time"
)

type StartOfWorkResponse struct {
	Status     string `json:"status"`
	User       models.User
	WorkRecord models.WorkRecord
}

func StartOfWork() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		auth := c.Get("auth").(*models.User)

		user := models.User{}
		dbs.DB.Table("users").Where(models.User{ID: auth.ID}).First(&user)

		time.Local = time.FixedZone("Local", 9*60*60)
		jst, _ := time.LoadLocation("Local")
		date := time.Now().In(jst)
		today := date.Format("2006-01-02")
		currentTime := date.Format("15:04")

		workrecord := models.WorkRecord{}
		if dbs.DB.Where(models.WorkRecord{UserId: user.ID, Date: today}).First(&workrecord).RecordNotFound() {
			workrecord = models.WorkRecord{UserId: user.ID, Date: today, StartOfWork: currentTime}
			dbs.DB.Create(&workrecord)
		} else {
			dbs.DB.Where(models.WorkRecord{UserId: user.ID, Date: today}).First(&workrecord)
		}

		response := StartOfWorkResponse{Status: "SUCCESS", User: user, WorkRecord: workrecord}

		return c.JSON(fasthttp.StatusOK, response)
	}
}
