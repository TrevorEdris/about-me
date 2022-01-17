package routes

import (
	"github.com/TrevorEdris/about-me/controller"
	"github.com/labstack/echo/v4"
)

type (
	Resume struct {
		controller.Controller
	}
)

func (c *Resume) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "resume"
	page.Title = "Public Resume"
	page.Tall = true

	page.Cache.Enabled = false

	return c.RenderPage(ctx, page)
}
