package router

import (
	"github.com/labstack/echo/v4"
	"github.com/testsla/jojo/server/Admin/controller"
	"net/http"
)

type Route struct {
	Path    string
	Handler echo.HandlerFunc
	Method  []string
}

var RouteMap = [...]Route{
	{
		Path:    "/",
		Handler: controller.Home,
		Method:  []string{http.MethodGet},
	},
	{
		Path:    "GetAllUser",
		Handler: controller.GetAllUser,
		Method:  []string{http.MethodGet,http.MethodPost},
	},
	{
		Path:    "login",
		Handler: controller.Login,
		Method:  []string{http.MethodPost},
	},
}
