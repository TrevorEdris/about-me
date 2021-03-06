//go:build integration
// +build integration

package middleware

import (
	"os"
	"testing"

	"github.com/TrevorEdris/about-me/config"
	"github.com/TrevorEdris/about-me/ent"
	"github.com/TrevorEdris/about-me/services"
)

var (
	c   *services.Container
	usr *ent.User
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Create a new container
	c = services.NewContainer()
	defer func() {
		if err := c.Shutdown(); err != nil {
			c.Web.Logger.Fatal(err)
		}
	}()

	// TODO: Maybe re-enable at a later date
	// // Create a user
	// var err error
	// if usr, err = tests.CreateUser(c.ORM); err != nil {
	// 	panic(err)
	// }

	// Run tests
	exitVal := m.Run()
	os.Exit(exitVal)
}
