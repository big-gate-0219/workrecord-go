package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"middlewares"
	"models"
	"strconv"
	"validate"
)

type GetGroupUsersResponse struct {
	Status string        `json:"status"`
	Users  []models.User `json:"users"`
}

func GetGroupUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		groupId, _ := strconv.ParseUint(c.Param("group_id"), 10, 64)
		g := models.Group{}
		if dbs.Transaction.Where(&models.Group{ID: groupId}).First(&g).RecordNotFound() {
			validateError := validate.CreateSingleErrors("not_found", "group_id")
			errResponse := validate.CreateErrorResponse(validateError)
			return c.JSON(fasthttp.StatusBadRequest, errResponse)
		}

		group := models.Group{}
		dbs.Transaction.Where(models.Group{ID: groupId}).Find(&group)

		groupUsers := []models.GroupUser{}
		dbs.Transaction.Where(models.GroupUser{GroupId: groupId}).Find(&groupUsers)

		users := []models.User{}
		for _, groupUser := range groupUsers {
			user := models.User{}
			dbs.Transaction.Where(models.User{ID: groupUser.UserId}).First(&user)
			users = append(users, user)
		}

		response := GetGroupUsersResponse{
			Status: "SUCCESS",
			Users:  users,
		}
		return c.JSON(fasthttp.StatusOK, response)
	}
}
