package router

import (
	mw "echo-task/middleware"

	"github.com/labstack/echo"
)

func ApiRouting() *echo.Echo {
	e := echo.New()
	r := e.Group("api/v1")
	e.POST("/login", mw.Login)
	r.GET("/Get/Student", mw.GetAllStudents)
	r.POST("/SignIn", mw.SignIN)
	e.DELETE("/Delete/Student/:roll_no", mw.DeleteStudentDetails)
	return e
}
