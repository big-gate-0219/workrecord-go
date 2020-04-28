package api

import (
	"strconv"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"models"
	"middlewares"
)

type AddUserToGroupRequest struct {
	UserId uint64 `json:"user_id"`
}

func AddUserToGroup() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		groupId, _ := strconv.ParseUint(c.Param("group_id"), 10, 64)
		request := new(AddUserToGroupRequest)
		if err := c.Bind(request); err != nil {
			return err
		}

		group := models.Group{}
		if dbs.DB.Where(models.Group{ID: groupId}).First(&group).RecordNotFound() {
			return c.JSON(fasthttp.StatusInternalServerError, "Not found group.")
		}

		user := models.User{}
		if dbs.DB.Where(models.User{ID: request.UserId}).First(&user).RecordNotFound() {
			return c.JSON(fasthttp.StatusInternalServerError, "Not found user.")
		}

		groupUser := models.GroupUser {
			GroupId: group.ID,
			UserId: user.ID,
		}
		dbs.DB.Create(&groupUser)

		return c.JSON(fasthttp.StatusOK, "")
	}
}