package api

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"middlewares"
	"models"
	"strings"
	"validate"
)

type SignInRequest struct {
	UID      string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := new(SignInRequest)
		if err := c.Bind(request); err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err)
		}
		if err := c.Validate(request); err != nil {
			validateError := validate.TranslateError(err.(validator.ValidationErrors))
			errResponse := validate.CreateErrorResponse(validateError)
			return c.JSON(fasthttp.StatusBadRequest, errResponse)
		}

		password_seed := request.UID + ":" + request.Password
		password_sum256 := sha256.Sum256([]byte(password_seed))
		password_hex := strings.ToUpper(hex.EncodeToString(password_sum256[:]))

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		user := models.User{}
		if dbs.DB.Table("users").Where(&models.User{UID: request.UID, Password: password_hex}).First(&user).RecordNotFound() {
			validateError := validate.CreateSingleErrors("unauthorized", "user_id+password")
			errResponse := validate.CreateErrorResponse(validateError)
			return c.JSON(fasthttp.StatusUnauthorized, errResponse)
		}

		tokenString, _ := middlewares.Generate(user)
		res := SignInResponse{
			Token: tokenString,
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
