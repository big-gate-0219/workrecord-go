package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"middlewares"
	"models"
	"time"
)

type GetWorkRecordResponse struct {
	Status      string `json:"status"`
	User        models.User
	WorkRecords []models.WorkRecord
}

func GetWorkRecord() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		auth := c.Get("auth").(*models.User)

		time.Local = time.FixedZone("Local", 9*60*60)
		jst, _ := time.LoadLocation("Local")
		date := time.Now().In(jst)
		todayYearMonth := date.Format("2006-01")

		user := models.User{}
		workrecords := []models.WorkRecord{}
		dbs.DB.Where(models.User{ID: auth.ID}).First(&user)
		dbs.DB.Where("user_id = ?", user.ID).Where("Date like ?", todayYearMonth+"%").Find(&workrecords)

		response := GetWorkRecordResponse{Status: "SUCCESS", User: user, WorkRecords: workrecords}

		return c.JSON(fasthttp.StatusOK, response)
	}
}
