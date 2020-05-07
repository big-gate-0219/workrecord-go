package middlewares

import (
	"databases"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type DatabaseClient struct {
	DB *gorm.DB
}

func DatabasesService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := databases.Connect()
			d := DatabaseClient{DB: session}
			tx := d.DB.Begin()

			defer d.DB.Close()

			d.DB.LogMode(true)
			c.Set("dbs", &d)
			c.Set("tx", &tx)
			if err := next(c); err != nil {
				tx.Rollback()
				logrus.Debug("Transction Rollback: ", err)
				return err
			}
			if c.Response().Status != 200 {
				tx.Rollback()
				logrus.Debug("Transction Rollback Status: ", c.Response().Status)
				return nil
			}

			logrus.Debug("Transaction Commit")
			tx.Commit()
			return nil
		}
	}
}
