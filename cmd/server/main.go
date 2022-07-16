package main

import (
	"github.com/ShariqT/csfood4/pkg/routes"
	"github.com/ShariqT/csfood4/pkg/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/siredwin/pongorenderer/renderer"
)

func main() {
	e := echo.New()
	MainRenderer := renderer.Renderer{Debug: true}
	e.Renderer = MainRenderer
	e.Use(middleware.Logger())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &utils.CommonstockContext{c}
			return next(cc)
		}
	})
	e.GET("/", routes.IndexHandler)
	e.GET("/login", routes.ShowLogin)
	e.POST("/login", routes.IndexHandler)
	e.Start(":5000")
}
