package services

import (
	"crypto/tls"
	"fmt"

	"github.com/mikestefanello/pagoda/config"

	"github.com/labstack/echo/v4"
	gomail "gopkg.in/mail.v2"
)

// MailClient provides a client for sending email
// This is purposely not completed because there are many different methods and services
// for sending email, many of which are very different. Choose what works best for you
// and populate the methods below
type MailClient struct {
	// config stores application configuration
	config *config.Config

	// templates stores the template renderer
	templates *TemplateRenderer
}

// NewMailClient creates a new MailClient
func NewMailClient(cfg *config.Config, templates *TemplateRenderer) (*MailClient, error) {
	return &MailClient{
		config:    cfg,
		templates: templates,
	}, nil
}

// Send sends an email to a given email address with a given body
func (c *MailClient) Send(ctx echo.Context, contactorEmailAddr, body string) error {
	if c.skipSend() {
		ctx.Logger().Debugf("skipping email sent from: %s", contactorEmailAddr)
	}

	ctx.Logger().Debugf("Sending mail to %s from %s - Contactor %s", c.config.Mail.ToAddress, c.config.Mail.FromAddress, contactorEmailAddr)

	msg := gomail.NewMessage()
	msg.SetHeader("From", c.config.Mail.FromAddress)
	msg.SetHeader("To", c.config.Mail.ToAddress)
	msg.SetHeader("Subject", "Contact Me - "+contactorEmailAddr)
	msg.SetBody("text/plain", body)

	dialer := gomail.NewDialer(c.config.Mail.Hostname, int(c.config.Mail.Port), c.config.Mail.User, c.config.Mail.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(msg)
	if err != nil {
		return err
	}
	return nil
}

// SendTemplate sends an email to a given email address using a template and data which is passed to the template
// The template name should only include the filename without the extension or directory.
// The funcmap will be automatically added to the template and the data will be passed in.
func (c *MailClient) SendTemplate(ctx echo.Context, from, template string, data interface{}) error {
	if c.skipSend() {
		ctx.Logger().Debugf("skipping template email sent from: %s", from)
	}

	// Parse and execute template
	// Uncomment the first variable when ready to use
	_, err := c.templates.ParseAndExecute(
		"mail",
		template,
		template,
		[]string{fmt.Sprintf("emails/%s", template)},
		[]string{},
		data,
	)
	if err != nil {
		return err
	}

	return nil
}

// skipSend determines if mail sending should be skipped
func (c *MailClient) skipSend() bool {
	//return c.config.App.Environment != config.EnvProduction
	return false
}
