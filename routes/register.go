package routes

import (
	"fmt"

	"github.com/TrevorEdris/about-me/context"
	"github.com/TrevorEdris/about-me/controller"
	"github.com/TrevorEdris/about-me/ent"
	"github.com/TrevorEdris/about-me/msg"

	"github.com/labstack/echo/v4"
)

type (
	Register struct {
		controller.Controller
	}

	RegisterForm struct {
		Name            string `form:"name" validate:"required"`
		Email           string `form:"email" validate:"required,email"`
		Password        string `form:"password" validate:"required"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
		Submission      controller.FormSubmission
	}
)

func (c *Register) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "auth"
	page.Name = "register"
	page.Title = "Register"
	page.Form = RegisterForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*RegisterForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *Register) Post(ctx echo.Context) error {
	var form RegisterForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(ctx, err, "unable to parse register form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(ctx, err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.Get(ctx)
	}

	// Hash the password
	pwHash, err := c.Container.Auth.HashPassword(form.Password)
	if err != nil {
		return c.Fail(ctx, err, "unable to hash password")
	}

	// Attempt creating the user
	u, err := c.Container.ORM.User.
		Create().
		SetName(form.Name).
		SetEmail(form.Email).
		SetPassword(pwHash).
		Save(ctx.Request().Context())

	switch err.(type) {
	case nil:
		ctx.Logger().Infof("user created: %s", u.Name)
	case *ent.ConstraintError:
		msg.Warning(ctx, "A user with this email address already exists. Please log in.")
		return c.Redirect(ctx, "login")
	default:
		return c.Fail(ctx, err, "unable to create user")
	}

	// Log the user in
	err = c.Container.Auth.Login(ctx, u.ID)
	if err != nil {
		ctx.Logger().Errorf("unable to log in: %v", err)
		msg.Info(ctx, "Your account has been created.")
		return c.Redirect(ctx, "login")
	}

	msg.Success(ctx, "Your account has been created. You are now logged in.")

	// Send the verification email
	c.sendVerificationEmail(ctx, u)

	return c.Redirect(ctx, "home")
}

func (c *Register) sendVerificationEmail(ctx echo.Context, usr *ent.User) {
	// Generate a token
	token, err := c.Container.Auth.GenerateEmailVerificationToken(usr.Email)
	if err != nil {
		ctx.Logger().Errorf("unable to generate email verification token: %v", err)
		return
	}

	// Send the email
	err = c.Container.Mail.Send(ctx, usr.Email, fmt.Sprintf(
		"Confirm your email address: %s",
		ctx.Echo().Reverse("verify_email", token),
	))
	if err != nil {
		ctx.Logger().Errorf("unable to send email verification token: %v", err)
		return
	}

	msg.Info(ctx, "An email was sent to you to verify your email address.")
}
