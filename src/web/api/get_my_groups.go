package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"models"
	"middlewares"
)

type GetMyGroupsResponse struct {
	Status string `json:"status"`
	Groups []models.Group
}

func GetMyGroups() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		auth := c.Get("auth").(*models.User)

		user := models.User{}
		dbs.DB.Where(models.User{ID: auth.ID}).First(&user)

		groupUsers := []models.GroupUser{}
		dbs.DB.Where(models.GroupUser{UserId: user.ID}).Find(&groupUsers)

		groups := []models.Group{};
		for _, groupUser := range groupUsers {
			group := models.Group{}
			dbs.DB.Where(models.Group{ID: groupUser.GroupId}).First(&group)
			groups = append(groups, group)
		}

		response := GetMyGroupsResponse{Status: "SUCCESS", Groups: groups}
		return c.JSON(fasthttp.StatusOK, response)
	}
}
