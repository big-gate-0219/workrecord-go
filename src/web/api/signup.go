package api

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"middlewares"
	"models"
	"validate"

	"strings"
)

type SignUpRequest struct {
	UID         string `json:"userId" validate:"required,min=5,max=20"`
	MailAddress string `json:"email" validate:"required,email"`
	Name        string `json:"userName" validate:"required,min=5,max=20,excludesall=!()#@{}"`
	Password    string `json:"password" validate:"required,min=5,max=16"`
}

type SignUpResponse struct {
	Status string      `json:"status"`
	User   models.User `json:user"`
}

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := new(SignUpRequest)
		if err := c.Bind(request); err != nil {
			return err
		}
		if err := c.Validate(request); err != nil {
			validateError := validate.TranslateError(err.(validator.ValidationErrors))
			errResponse := validate.CreateErrorResponse(validateError)
			return c.JSON(fasthttp.StatusBadRequest, errResponse)
		}

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		u := models.User{}
		if !dbs.Transaction.Where(&models.User{UID: request.UID}).First(&u).RecordNotFound() {
			validateError := validate.CreateSingleErrors("duplicated", "userId")
			errResponse := validate.CreateErrorResponse(validateError)
			return c.JSON(fasthttp.StatusBadRequest, errResponse)
		}

		password_seed := request.UID + ":" + request.Password
		password_sum256 := sha256.Sum256([]byte(password_seed))
		password_hex := strings.ToUpper(hex.EncodeToString(password_sum256[:]))

		user := models.User{
			UID:         request.UID,
			MailAddress: request.MailAddress,
			Name:        request.Name,
			Password:    password_hex,
		}

		dbs.Transaction.Create(&user)
		response := SignUpResponse{
			Status: "SUCCESS",
			User:   user,
		}

		return c.JSON(fasthttp.StatusOK, response)
	}
}
