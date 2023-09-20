package utils

import (
	"crypto/tls"
	"fmt"
	"n1h41/auth-service/config"

	"gopkg.in/mail.v2"
)

func SendMail(m *mail.Message) (err error) {
	config, _ := config.LoadConfig("../")
	d := mail.NewDialer(config.SmtpHost, 587, config.SmtpUser, config.SmtpPass)
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
  if err := d.DialAndSend(m); err != nil {
    fmt.Println("Error sending mail: ", err)
    return err
  }
  return nil
}
