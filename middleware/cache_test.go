//go:build integration
// +build integration

package middleware

import (
	"testing"
)

func TestServeCachedPage(t *testing.T) {
	// TODO: Maybe re-enable at a later date
	// 	// Cache a page
	// 	cp := CachedPage{
	// 		URL:        "/cache",
	// 		HTML:       []byte("html"),
	// 		Headers:    make(map[string]string),
	// 		StatusCode: http.StatusCreated,
	// 	}
	// 	cp.Headers["a"] = "b"
	// 	cp.Headers["c"] = "d"
	// 	err := marshaler.New(c.Cache).Set(context.Background(), cp.URL, cp, nil)
	// 	require.NoError(t, err)

	// 	// Request the URL of the cached page
	// 	ctx, rec := tests.NewContext(c.Web, cp.URL)
	// 	err = tests.ExecuteMiddleware(ctx, ServeCachedPage(c.Cache))
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, cp.StatusCode, ctx.Response().Status)
	// 	assert.Equal(t, cp.Headers["a"], ctx.Response().Header().Get("a"))
	// 	assert.Equal(t, cp.Headers["c"], ctx.Response().Header().Get("c"))
	// 	assert.Equal(t, cp.HTML, rec.Body.Bytes())

	// 	// Login and try again
	// 	tests.InitSession(ctx)
	// 	err = c.Auth.Login(ctx, usr.ID)
	// 	require.NoError(t, err)
	// 	_ = tests.ExecuteMiddleware(ctx, LoadAuthenticatedUser(c.Auth))
	// 	err = tests.ExecuteMiddleware(ctx, ServeCachedPage(c.Cache))
	// 	assert.Nil(t, err)
	// }

	// func TestCacheControl(t *testing.T) {
	// 	ctx, _ := tests.NewContext(c.Web, "/")
	// 	_ = tests.ExecuteMiddleware(ctx, CacheControl(time.Second*5))
	// 	assert.Equal(t, "public, max-age=5", ctx.Response().Header().Get("Cache-Control"))
	// 	_ = tests.ExecuteMiddleware(ctx, CacheControl(0))
	// 	assert.Equal(t, "no-cache, no-store", ctx.Response().Header().Get("Cache-Control"))
}
