package routes

import (
	"github.com/TrevorEdris/about-me/context"
	"github.com/TrevorEdris/about-me/controller"
	"github.com/TrevorEdris/about-me/ent"
	"github.com/TrevorEdris/about-me/ent/user"
	"github.com/TrevorEdris/about-me/msg"
	"github.com/labstack/echo/v4"
)

type VerifyEmail struct {
	controller.Controller
}

func (c *VerifyEmail) Get(ctx echo.Context) error {
	var usr *ent.User

	// Validate the token
	token := ctx.Param("token")
	email, err := c.Container.Auth.ValidateEmailVerificationToken(token)
	if err != nil {
		msg.Warning(ctx, "The link is either invalid or has expired.")
		return c.Redirect(ctx, "home")
	}

	// Check if it matches the authenticated user
	if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
		authUser := u.(*ent.User)

		if authUser.Email == email {
			usr = authUser
		}
	}

	// Query to find a matching user, if needed
	if usr == nil {
		usr, err = c.Container.ORM.User.
			Query().
			Where(user.Email(email)).
			Only(ctx.Request().Context())

		if err != nil {
			return c.Fail(ctx, err, "query failed loading email verification token user")
		}
	}

	// Verify the user, if needed
	if !usr.Verified {
		err = c.Container.ORM.User.
			Update().
			SetVerified(true).
			Where(user.ID(usr.ID)).
			Exec(ctx.Request().Context())

		if err != nil {
			return c.Fail(ctx, err, "failed to set user as verified")
		}
	}

	msg.Success(ctx, "Your email has been successfully verified.")
	return c.Redirect(ctx, "home")
}
