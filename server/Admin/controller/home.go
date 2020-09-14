package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/testsla/jojo/server/Admin/model"
	"net/http"
	"reflect"
)

// [GET /home]
func Home(c echo.Context) error {
	c.JSON(http.StatusOK, echo.Map{
		"msg": "ok",
	})
	s := reflect.ValueOf(GetAllUser).Type()
	fmt.Println(s)
	return nil
}

func GetAllUser(c echo.Context) (err error) {
	c.GetDB().Save(model.NewUser())
	c.JSON(http.StatusOK, echo.Map{
		"msg":   "success",
		"users": "new User",
	})
	return nil
}
