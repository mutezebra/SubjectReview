package service

import (
	"bytes"
	"context"
	"github.com/mutezebra/subject-review/biz/model/subject"
	"github.com/mutezebra/subject-review/pkg/constants"
	"github.com/mutezebra/subject-review/pkg/logger"
	"github.com/mutezebra/subject-review/pkg/utils"
	"github.com/pkg/errors"
	"html/template"
	"time"
)

type sendEmail struct {
	records []*subject.SubjectRecord
	email   string
}

func (svc *SubjectService) Remind() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan *sendEmail, constants.SendEmailRoutines)
	for i := 0; i < constants.SendEmailRoutines; i++ {
		go svc.sendEmail(ctx, ch)
	}

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			svc.remind(ctx, ch)
		}
	}
}

func (svc *SubjectService) remind(ctx context.Context, ch chan *sendEmail) {
	now := time.Now().Unix()
	if now%constants.RemindInterval >= 59 {
		return
	}

	uids, err := svc.SubjectCache.GetManagers(ctx)
	if err != nil {
		logger.Errorf("Error getting managers: %v", err)
		return
	}

	for _, uid := range uids {
		records, _, err := svc.SubjectDB.GetNeedReviewedSubjects(ctx, uid, 1, 10, time.Now().Unix()+constants.DefaultRemindEndTime)
		if err != nil {
			logger.Errorf("Error getting need reviewed subjects: %v", err)
			continue
		}

		if len(records) > 0 {
			email, err := svc.SubjectDB.GetEmailByID(ctx, uid)
			if err != nil {
				logger.Errorf("Error getting email by ID: %v", err)
				continue
			}
			ch <- &sendEmail{records: records, email: email}
		}
	}

	time.Sleep(60 * time.Second)
}

func (svc *SubjectService) sendEmail(ctx context.Context, ch chan *sendEmail) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-ch:
			body := parse(utils.BuildSubjectRecordResp(data.records))
			if err := svc.Email.SendRemindMessage(data.email, body); err != nil {
				logger.Error(errors.WithMessage(err, "failed when send remind message: "))
				break
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func parse(records []*subject.SubjectRecordResp) *string {
	tmpl := template.Must(template.New("emailTemplate").Parse(constants.HTMLTemplate))

	var body bytes.Buffer
	err := tmpl.Execute(&body, map[string]interface{}{
		"Records": records,
	})
	if err != nil {
		logger.Fatal(err)
	}
	s := body.String()
	return &s
}
