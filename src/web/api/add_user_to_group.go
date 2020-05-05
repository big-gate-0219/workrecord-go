package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"middlewares"
	"models"
	"strconv"
	"validate"
)

type AddUserToGroupRequest struct {
	UserId uint64 `json:"user_id" validate:"required"`
}

func AddUserToGroup() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		groupId, _ := strconv.ParseUint(c.Param("group_id"), 10, 64)

		request := new(AddUserToGroupRequest)
		if err := c.Bind(request); err != nil {
			return err
		}
		if err := c.Validate(request); err != nil {
			validateError := validate.TranslateError(err.(validator.ValidationErrors))
			errResponse := validate.CreateErrorResponse(validateError)
			return c.JSON(fasthttp.StatusBadRequest, errResponse)
		}

		group := models.Group{}
		if dbs.DB.Where(models.Group{ID: groupId}).First(&group).RecordNotFound() {
			validateError := validate.CreateSingleErrors("not_found", "groupId")
			errResponse := validate.CreateErrorResponse(validateError)
			return c.JSON(fasthttp.StatusBadRequest, errResponse)
		}

		user := models.User{}
		if dbs.DB.Where(models.User{ID: request.UserId}).First(&user).RecordNotFound() {
			validateError := validate.CreateSingleErrors("not_found", "user_id")
			errResponse := validate.CreateErrorResponse(validateError)
			return c.JSON(fasthttp.StatusBadRequest, errResponse)
		}

		groupUser := models.GroupUser{
			GroupId: group.ID,
			UserId:  user.ID,
		}
		dbs.DB.Create(&groupUser)

		return c.JSON(fasthttp.StatusOK, "")
	}
}
