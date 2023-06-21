package mail

import (
	config "mygram/infrastructures"

	"gopkg.in/gomail.v2"
)

type Mail interface {
	SendMail(to string, message string)
}

type MailService struct {
	Dialer *gomail.Dialer
}

func NewMailService(configuration config.Config) Mail {
	host := "sandbox.smtp.mailtrap.io"
	user := "62ae93b062182d"
	password := "102fc868c74199"
	port := 465
	d := gomail.NewDialer(host, port, user, password)
	d.SSL = false

	return &MailService{
		Dialer: d,
	}
}

func (service *MailService) SendMail(to string, message string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", message)

	// Send the email to Bob, Cora and Dan.
	if err := service.Dialer.DialAndSend(m); err != nil {
		panic(err)
	}
}

// import (
// 	"strconv"

// 	config "mygram/infrastructures"

// 	"gopkg.in/gomail.v2"
// )

// type MailService interface {
// 	configMail(config.Config) *gomail.Dialer
// 	sendMail(d *gomail.Dialer)
// }

// func NewMailService(service MailService) MailService {
// 	return &MailServiceImpl{
// 		MailService: service,
// 	}
// }

// type MailServiceImpl struct {
// 	MailService MailService
// }

// // configMail implements MailService
// func (service *MailServiceImpl) configMail(configuration config.Config) *gomail.Dialer {
// 	host := configuration.Get("MAIL_HOST")
// 	user := configuration.Get("MAIL_USER")
// 	password := configuration.Get("MAIL_PASSWORD")
// 	port, _ := strconv.Atoi(configuration.Get("MAIL_PORT"))
// 	d := gomail.NewDialer(host, port, user, password)
// 	return d
// }

// // sendMail implements MailService
// func (service *MailServiceImpl) sendMail(d *gomail.Dialer) {
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", "alex@example.com")
// 	m.SetHeader("To", "bob@example.com", "cora@example.com")
// 	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
// 	m.SetHeader("Subject", "Hello!")
// 	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

// 	// Send the email to Bob, Cora and Dan.
// 	if err := d.DialAndSend(m); err != nil {
// 		panic(err)
// 	}
// }
