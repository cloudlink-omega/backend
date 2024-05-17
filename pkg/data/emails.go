package data

import (
	"bytes"
	"crypto/tls"
	"html/template"

	structs "github.com/cloudlink-omega/backend/pkg/structs"
	gomail "gopkg.in/mail.v2"
)

// SendHTMLEmail sends an HTML email using the provided Email Arguments and Template Data.
//
// args: EmailArgs struct containing email details
// data: TemplateData struct containing template data
func (mgr *Manager) SendHTMLEmail(args *structs.EmailArgs, data *structs.TemplateData) {

	var t *template.Template
	var err error

	// Create new message
	m := gomail.NewMessage()

	// Format headers
	m.SetHeader("From", m.FormatAddress(mgr.MailConfig.Username, "CloudLink Omega"))
	m.SetHeader("To", args.To)
	m.SetHeader("Subject", args.Subject)

	// Use HTML template
	if t, err = template.ParseFiles("./email_templates/" + args.Template + ".html"); err != nil {
		panic(err)
	}

	var body bytes.Buffer
	var temp = *data
	temp.ServerName = mgr.ServerNickname
	t.Execute(&body, temp)
	m.SetBody("text/html", body.String())

	// Prepare message for SMTP transmission
	d := gomail.NewDialer(
		mgr.MailConfig.Server,
		mgr.MailConfig.Port,
		mgr.MailConfig.Username,
		mgr.MailConfig.Password,
	)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send E-Mail
	if err = d.DialAndSend(m); err != nil {
		panic(err)
	}
}

// SendPlainEmail sends a plain text email using the provided Email Arguments and data.
//
// args: EmailArgs struct containing email details
// data: string containing the plain text email content
func (mgr *Manager) SendPlainEmail(args *structs.EmailArgs, data string) {

	// Create new message
	m := gomail.NewMessage()

	// Format headers
	m.SetHeader("From", m.FormatAddress(mgr.MailConfig.Username, "CloudLink Omega"))
	m.SetHeader("To", args.To)
	m.SetHeader("Subject", args.Subject)

	// Use plaintext
	m.SetBody("text/plain", data)

	// Prepare message for SMTP transmission
	d := gomail.NewDialer(
		mgr.MailConfig.Server,
		mgr.MailConfig.Port,
		mgr.MailConfig.Username,
		mgr.MailConfig.Password,
	)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send E-Mail
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
