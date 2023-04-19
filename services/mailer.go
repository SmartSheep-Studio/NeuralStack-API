package services

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/viper"
	"net/mail"
	"net/smtp"
	"strings"
)

// parse map to mail body
func parse(headers map[string]string, content string) string {
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	return message + "\r\n" + content
}

// SendMail using smtp tls protocol
func SendMail(recipient string, subject string, content string) error {
	server := strings.Split(viper.GetString("mailer.host"), ":")[0]
	authorization := smtp.PlainAuth("", viper.GetString("mailer.username"), viper.GetString("mailer.password"), server)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         server,
	}

	// Prepare headers
	from := mail.Address{Name: viper.GetString("mailer.name"), Address: viper.GetString("mailer.mail")}
	to := mail.Address{Address: recipient}
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = "text/html; charset=UTF-8"

	// Start handshake and transfer data packs
	tcp, err := tls.Dial("tcp", viper.GetString("mailer.host"), tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(tcp, server)
	if err != nil {
		return err
	}
	if err = client.Auth(authorization); err != nil {
		return err
	}
	if err = client.Mail(from.Address); err != nil {
		return err
	}
	if err = client.Rcpt(to.Address); err != nil {
		return err
	}

	writer, err := client.Data()
	if err != nil {
		return err
	}

	message := parse(headers, content)
	if _, err = writer.Write([]byte(message)); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	return client.Quit()
}
