package middlewares

import (
	"github.com/pkg/errors"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func AuthenticationGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logrus.Debug("[Start ] AuthenticationGuard")

			authToken := c.Request().Header.Get("x-auth-token")
			auth, err := Parse(authToken)
			if err != nil {
				return errors.Wrapf(err, "failed to parse authorization")
			}

			c.Set("auth", auth)
			if err:= next(c); err !=nil {
				return err
			}

			logrus.Debug("[Finish] AuthenticationGuard")
			return nil
		}
	}
}