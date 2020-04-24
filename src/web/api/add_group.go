package api

import (
	"models"
	"middlewares"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type AddGroupRequest struct {
	GroupName string `json:"group_name"`
}

type AddGroupResponse struct {
	Status string `json:"status"`
	Group models.Group `json:"group"`
}

func AddGroup() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := new(AddGroupRequest)
		if err:= c.Bind(request); err != nil {
			return err
		}

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		group := models.Group{Name: request.GroupName}
		dbs.DB.Create(&group)

		auth := c.Get("auth").(*models.User)
		groupUser := models.GroupUser{GroupId: group.ID, UserId:auth.ID}
		dbs.DB.Create(&groupUser)
		
		response := AddGroupResponse {
			Status: "SUCCESS",
			Group: group,
		}

		return c.JSON(fasthttp.StatusOK, response)
	}
}