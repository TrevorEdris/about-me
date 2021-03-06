//go:build integration
// +build integration

package services

import (
	"os"
	"testing"

	"github.com/TrevorEdris/about-me/config"
	"github.com/TrevorEdris/about-me/ent"
	"github.com/TrevorEdris/about-me/tests"

	"github.com/labstack/echo/v4"
)

var (
	c   *Container
	ctx echo.Context
	usr *ent.User
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Create a new container
	c = NewContainer()
	defer func() {
		if err := c.Shutdown(); err != nil {
			c.Web.Logger.Fatal(err)
		}
	}()

	// Create a web context
	ctx, _ = tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)

	// TODO: Maybe re-enable at a later date
	// // Create a test user
	// var err error
	// if usr, err = tests.CreateUser(c.ORM); err != nil {
	// 	panic(err)
	// }

	// Run tests
	exitVal := m.Run()
	os.Exit(exitVal)
}
