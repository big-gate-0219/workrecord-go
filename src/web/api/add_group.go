package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"middlewares"
	"models"
	"validate"
)

type AddGroupRequest struct {
	GroupName string `json:"group_name" validate:"required,min=5,max=20"`
}

type AddGroupResponse struct {
	Status string       `json:"status"`
	Group  models.Group `json:"group"`
}

func AddGroup() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := new(AddGroupRequest)
		if err := c.Bind(request); err != nil {
			return err
		}
		if err := c.Validate(request); err != nil {
			validateError := validate.TranslateError(err.(validator.ValidationErrors))
			errResponse := validate.CreateErrorResponse(validateError)
			return c.JSON(fasthttp.StatusBadRequest, errResponse)
		}

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)

		g := models.Group{}
		if !dbs.Transaction.Where(&models.Group{Name: request.GroupName}).First(&g).RecordNotFound() {
			validateError := validate.CreateSingleErrors("duplicated", "group_name")
			errResponse := validate.CreateErrorResponse(validateError)
			return c.JSON(fasthttp.StatusBadRequest, errResponse)
		}

		group := models.Group{Name: request.GroupName}
		dbs.Transaction.Create(&group)

		auth := c.Get("auth").(*models.User)
		groupUser := models.GroupUser{GroupId: group.ID, UserId: auth.ID}
		dbs.Transaction.Create(&groupUser)

		response := AddGroupResponse{
			Status: "SUCCESS",
			Group:  group,
		}

		return c.JSON(fasthttp.StatusOK, response)
	}
}
