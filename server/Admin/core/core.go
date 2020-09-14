package core

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	middleware2 "github.com/testsla/jojo/server/Admin/middleware"
	"github.com/testsla/jojo/server/Admin/model"
	"github.com/testsla/jojo/server/Admin/router"
	"github.com/testsla/jojo/util"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

type server struct {
	Instance *echo.Echo
	//HttpPort int
	DB     *gorm.DB
	Config util.Map
}

var Server *server

func New() *server {
	config := util.GetServerConfig("admin")
	Server = &server{
		Instance: echo.New(),
		Config:   config,
	}
	Server.Instance.HideBanner = false
	return Server
}

func (s *server) HttpListen() *server {
	s.Instance.Debug = true
	s.Instance.Logger.Fatal(
		s.Instance.Start(fmt.Sprintf(":%v", s.Config["port"])),
	)
	return s
}
func (s *server) DBConnect() *server {
	var err error
	config := util.GetDataBaseConfig("admin")
	s.DB, err = gorm.Open(sqlite.Open(config["uri"].(string)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	s.Instance.Logger.Debugf("连接成功：[ %d => %d]", config["type"], config["uri"])
	s.DB.AutoMigrate(&model.User{})
	return s
}
func customHTTPErrorHandler(err error, c echo.Context) {
	c.JSON(http.StatusBadRequest, err)
	c.Logger().Error(err)
}
func (s *server) Middleware() *server {
	s.Instance.Use(middleware2.BindDB(s.DB))
	s.Instance.Use(middleware.CORS())
	s.Instance.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_custom} ${remote_ip}>${id} [${method}] - ${uri} { ${query} }\n",
	}))
	//s.Instance.Use(middleware.Recover())
	//s.Instance.Use(middleware.JWT([]byte("secret")))
	// 错误处理
	s.Instance.HTTPErrorHandler = customHTTPErrorHandler
	return s
}
func (s *server) Router() *server {
	for i, route := range router.RouteMap {
		s.Instance.Match(route.Method, route.Path, route.Handler)
		//fmt.Printf("%v:  %v  %v \n", i, route.Method, route.Path)
		s.Instance.Logger.Infof("%v:  %v  %v \n", i, route.Method, route.Path)
	}
	return s
}
