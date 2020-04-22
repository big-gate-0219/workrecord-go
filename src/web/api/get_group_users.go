package api

import (
	"strconv"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"models"
	"middlewares"
)

type GetGroupUsersResponse struct {
	Status string `json:"status"`
	Users []models.User `json:"users"`
}

func GetGroupUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		groupId, _ := strconv.ParseUint(c.Param("group_id"), 10, 64)

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		group := models.Group{}
		dbs.DB.Where(models.Group{ID: groupId}).Find(&group)

		groupUsers := []models.GroupUser{}
		dbs.DB.Where(models.GroupUser{GroupId: groupId}).Find(&groupUsers)

		users := []models.User{}
		for _, groupUser := range groupUsers {
			user := models.User{}
			dbs.DB.Where(models.User{ID: groupUser.UserId}).First(&user)
			users = append(users, user)
		}

		response := GetGroupUsersResponse {
			Status: "SUCCESS",
			Users: users,
		}
		return c.JSON(fasthttp.StatusOK, response)
	}
}