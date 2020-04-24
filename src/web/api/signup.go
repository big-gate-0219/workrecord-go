package api

import (
	"models"
	"middlewares"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type SignUpRequest struct {
	UID string `json:"userId"`
	MailAddress string `json:"email"`
	Name string `json:"userName"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Status string `json:"status"`
	User models.User `json:user"`
}


func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := new(SignUpRequest)
		if err:= c.Bind(request); err != nil {
			return err
		}

		password_seed := request.UID + ":" + request.Password
		password_sum256 := sha256.Sum256([]byte(password_seed))
		password_hex := strings.ToUpper(hex.EncodeToString(password_sum256[:]))

		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		user := models.User{
			UID: request.UID,
			MailAddress: request.MailAddress,
			Name: request.Name,
			Password: password_hex,
		}

		dbs.DB.Create(&user)
		response := SignUpResponse{
			Status: "SUCCESS",
			User: user,
		}

		return c.JSON(fasthttp.StatusOK, response)
	}
}