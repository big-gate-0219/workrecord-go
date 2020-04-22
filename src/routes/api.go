package routes

import (
	api "web/api"
	"middlewares"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	apiGroup := e.Group("/api")
	{
		accountsGroup := apiGroup.Group("/accounts")
		{
			accountsGroup.POST("/signin", api.SignIn())
		}

		workrecordGroup := apiGroup.Group("/work-records", middlewares.AuthenticationGuard())
		{
			workrecordGroup.GET("", api.GetWorkRecord())
			todayGroup := workrecordGroup.Group("/today")
			{
				todayGroup.POST("/start", api.StartOfWork())
				todayGroup.POST("/end", api.EndOfWork())
			}
			workrecordGroup.GET("/groups/:group_id/today", api.GetWorkrecordGroupToday())
		}

		groupGroup := apiGroup.Group("/groups", middlewares.AuthenticationGuard())
		{
			groupGroup.GET("", api.GetMyGroups())
			groupGroup.GET("/:group_id", api.GetGroupUsers())
		}
	}

}
