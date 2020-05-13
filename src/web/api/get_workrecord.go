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

		date := time.Now()
		todayYearMonth := date.Format("2006-01")

		user := models.User{}
		workrecords := []models.WorkRecord{}
		dbs.Transaction.Where(models.User{ID: auth.ID}).First(&user)
		dbs.Transaction.Where("user_id = ?", user.ID).Where("Date like ?", todayYearMonth+"%").Find(&workrecords)

		response := GetWorkRecordResponse{Status: "SUCCESS", User: user, WorkRecords: workrecords}

		return c.JSON(fasthttp.StatusOK, response)
	}
}
