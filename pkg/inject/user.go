package inject

import (
	subjectp "github.com/mutezebra/subject-review/app/subject/persistence"
	subjects "github.com/mutezebra/subject-review/app/subject/service"
	subjectUsecase "github.com/mutezebra/subject-review/app/subject/usecase"
	userp "github.com/mutezebra/subject-review/app/user/persistence"
	users "github.com/mutezebra/subject-review/app/user/service"
	userUsecase "github.com/mutezebra/subject-review/app/user/usecase"
	"github.com/mutezebra/subject-review/pkg/email"
	"github.com/mutezebra/subject-review/pkg/oss"
	"github.com/mutezebra/subject-review/pkg/utils"
)

func UserInject() {
	svc := users.NewUserService(&users.UserService{
		UserDB:    userp.NewUserDB(),
		UserCache: userp.NewUserCache(),
		Email:     email.NewEmail(),
		OSS:       oss.NewOss(),
		PwdRe:     utils.GetPasswordVerifyRe(),
		EmailRe:   utils.GetEmailVerifyRe(),
	})

	userUsecase.InitUsecase(svc)
}

func SubjectInject() {
	svc := subjects.NewSubjectService(&subjects.SubjectService{
		SubjectDB:    subjectp.NewSubjectDB(),
		SubjectCache: subjectp.NewSubjectCache(),
		Email:        email.NewEmail(),
	})

	subjectUsecase.InitUsecase(svc)
}
