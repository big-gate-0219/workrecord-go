package middlewares

import (
	"databases"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type DatabaseClient struct {
	DB          *gorm.DB
	Transaction *gorm.DB
}

func DatabasesService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := databases.Connect()
			d := DatabaseClient{
				DB:          session,
				Transaction: session.Begin(),
			}

			defer d.DB.Close()

			d.DB.LogMode(true)
			c.Set("dbs", &d)
			if err := next(c); err != nil {
				d.Transaction.Rollback()
				logrus.Debug("Transction Rollback: ", err)
				return err
			}
			if c.Response().Status != 200 {
				d.Transaction.Rollback()
				logrus.Debug("Transction Rollback Status: ", c.Response().Status)
				return nil
			}

			logrus.Debug("Transaction Commit")
			d.Transaction.Commit()
			return nil
		}
	}
}
