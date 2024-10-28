package email

import (
	"crypto/tls"
	"fmt"

	"github.com/pkg/errors"
	gomail "gopkg.in/gomail.v2"

	"github.com/mutezebra/subject-review/config"
)

type email struct {
	host     string
	port     int
	userName string
	secret   string
}

func NewEmail() *email {
	c := config.Email
	return &email{host: c.Host, port: c.Port, userName: c.UserName, secret: c.Secret}
}

func (e *email) SendVerifyCode(to string, message *string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.userName)              // 发件人
	m.SetHeader("To", to)                        // 收件人
	m.SetHeader("Subject", "Subject-Review 验证码") // 邮件主题
	m.SetBody("text/html", *message)             // 邮件正文

	d := gomail.NewDialer(
		e.host,
		e.port,
		e.userName,
		e.secret,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed send verify code to %s,error: ", to))
	}
	return nil
}

func (e *email) SendRemindMessage(to string, msg *string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.userName)   // 发件人
	m.SetHeader("To", to)             // 收件人
	m.SetHeader("Subject", "你待复习的题目") // 邮件主题
	m.SetBody("text/html", *msg)      // 邮件正文

	d := gomail.NewDialer(
		e.host,
		e.port,
		e.userName,
		e.secret,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return errors.New(fmt.Sprintf("failed send verify code to %s", to))
	}
	return nil
}
