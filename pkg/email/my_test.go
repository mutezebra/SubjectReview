package email

import (
	"github.com/mutezebra/subject-review/pkg/utils"
	"testing"

	"github.com/mutezebra/subject-review/config"
)

func TestSendEmail(t *testing.T) {
	config.InitConfig()
	mail := NewEmail()
	t.Log(mail.SendVerifyCode("827545521@qq.com", utils.Ptr("Hello")))
}
