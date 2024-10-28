package email

import (
	"testing"

	"github.com/mutezebra/subject-review/pkg/utils"

	"github.com/mutezebra/subject-review/config"
)

func TestSendEmail(t *testing.T) {
	config.InitConfig()
	mail := NewEmail()
	t.Log(mail.SendVerifyCode("827545521@qq.com", utils.Ptr("Hello")))
}
