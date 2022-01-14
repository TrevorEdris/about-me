//go:build integration
// +build integration

package middleware

import (
	"testing"
)

func TestLoadUser(t *testing.T) {
	// TODO: Maybe re-enable at a later date
	// ctx, _ := tests.NewContext(c.Web, "/")
	// ctx.SetParamNames("user")
	// ctx.SetParamValues(fmt.Sprintf("%d", usr.ID))
	// _ = tests.ExecuteMiddleware(ctx, LoadUser(c.ORM))
	// ctxUsr, ok := ctx.Get(context.UserKey).(*ent.User)
	// require.True(t, ok)
	// assert.Equal(t, usr.ID, ctxUsr.ID)
}
