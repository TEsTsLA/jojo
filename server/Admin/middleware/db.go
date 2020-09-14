package middleware

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BindDB(db *gorm.DB) echo.MiddlewareFunc  {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if context.GetDB() == nil {
				context.SetDB(db)
			}
			return next(context)
		}
	}
}