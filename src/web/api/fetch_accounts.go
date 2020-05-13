package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"middlewares"
	"models"
)

type FetchAccountsResponse struct {
	Status string `json:"status"`
	Users  []models.User
}

func FetchAccounts() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		query := c.QueryParam("q")
		users := []models.User{}
		dbs.Transaction.Where("uid like ?", "%"+query+"%").Find(&users)

		response := FetchAccountsResponse{
			Status: "SUCCESS",
			Users:  users,
		}
		return c.JSON(fasthttp.StatusOK, response)
	}

}
