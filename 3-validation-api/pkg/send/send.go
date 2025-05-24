package send

import (
	"fmt"
	"go/hw/3-validation-api/configs"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(conf *configs.Config, link string, toAddr string) (bool, error) {
	e := email.NewEmail()
	e.From = conf.Email
	e.To = []string{toAddr}
	e.Subject = "Email validation"
	e.HTML = []byte(fmt.Sprintf(
		"<h1>Validation</h1><p>To validate email, go to: <a>%s</a>",
		link,
	))
	fmt.Println(conf)
	fmt.Println(link, toAddr)
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", conf.Email, conf.Password, conf.Address))
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
