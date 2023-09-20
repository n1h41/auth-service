package utils

import (
	"testing"

	"gopkg.in/mail.v2"
)

func TestSendMail(t *testing.T) {
  m := mail.NewMessage()
  m.SetHeader("From", "sureshgopi@gmail.com")
  m.SetHeader("To", "nihalninu25@gmail.com")
  m.SetHeader("Subject", "Hello!")
  m.SetBody("text/html", "Hello! <b>Nihal</b>")

  SendMail(m)
}
