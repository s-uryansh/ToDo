package jobs

import (
	"CRUD-SQL/utils"
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"text/template"
)

// Helpers
var t *template.Template

func init() {
	const templatePath = utils.REGISTER_TemplatePath
	var err error
	t, err = template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
}

// SendEmailRegistration sends a registration email to the user
func SendEmailRegistration(email string, name string) error {
	if t == nil {
		return fmt.Errorf("email template is not initialized")
	}

	from := utils.FROM_MAIL
	password := utils.PASSWORD_MAIL

	// SMTP setup
	auth := smtp.PlainAuth("", from, password, utils.SMTP_ADDRESS)
	conn, err := smtp.Dial(utils.SMTP_DIAL)
	if err != nil {
		return fmt.Errorf("error during SMTP Dial: %v", err)
	}
	defer conn.Close()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         utils.SMTP_ADDRESS,
	}
	if err := conn.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("error during StartTLS: %v", err)
	}

	if err := conn.Auth(auth); err != nil {
		return fmt.Errorf("error during SMTP Auth: %v", err)
	}

	if err := conn.Mail(from); err != nil {
		return fmt.Errorf("error during setting Mail sender: %v", err)
	}
	if err := conn.Rcpt(email); err != nil {
		return fmt.Errorf("error during setting Mail recipient: %v", err)
	}

	w, err := conn.Data()
	if err != nil {
		return fmt.Errorf("error during getting Data writer: %v", err)
	}
	defer w.Close()

	data := struct {
		Name string
	}{
		Name: name,
	}

	var bodyBuffer bytes.Buffer
	if err := t.Execute(&bodyBuffer, data); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	// Properly format the email with HTML content
	headers := "MIME-Version: 1.0\r\n"
	headers += "Content-Type: text/html; charset=\"UTF-8\"\r\n"
	headers += "Subject: " + utils.REGISTER_Subject + "\r\n"
	headers += "\r\n"

	emailBody := headers + bodyBuffer.String()
	_, err = w.Write([]byte(emailBody))
	if err != nil {
		return fmt.Errorf("error writing email body: %v", err)
	}
	return nil
}
