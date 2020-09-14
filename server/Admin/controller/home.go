package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/testsla/jojo/server/Admin/model"
	"github.com/testsla/jojo/util"
)

func Home(c echo.Context) error {
	return util.NotLoginError
}

func GetAllUser(c echo.Context) (err error) {
	var allUsers []model.User
	c.GetDB().Find(&allUsers)
	c.Ok(echo.Map{
		"msg":   "success",
		"users": allUsers,
	})
	return nil
}
