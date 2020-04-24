package api

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"middlewares"
	"models"
)

type SignInRequest struct {
	UID      string `json:"user_id"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := new(SignInRequest)
		if err := c.Bind(request); err != nil {
			return err
		}

		password_seed := request.UID + ":" + request.Password
		password_sum256 := sha256.Sum256([]byte(password_seed))
		password_hex := strings.ToUpper(hex.EncodeToString(password_sum256[:]))

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		user := models.User{}
		if dbs.DB.Table("users").Where(&models.User{UID: request.UID, Password: password_hex}).First(&user).RecordNotFound() {
			return c.JSON(fasthttp.StatusUnauthorized, "ユーザ名もしくはパスワードが間違っています。")
		}

		tokenString, _ := middlewares.Generate(user)
		res := SignInResponse{
			Token: tokenString,
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
