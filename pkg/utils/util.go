package utils

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/mutezebra/subject-review/biz/model/subject"
)

var sb = strings.Builder{}

func GenerateCode(length int) string {
	defer sb.Reset()
	codes := [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := 0; i < length; i++ {
		sb.WriteByte(codes[rand.Intn(len(codes))])
	}
	return sb.String()
}

func Ptr[T any](v T) *T {
	return &v
}

var (
	once sync.Once
	ctx  context.Context
)

func NewContextCancelWhenExit() context.Context {
	once.Do(func() {
		var stop context.CancelFunc
		ctx, stop = signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		go func() {
			defer stop()
			<-ctx.Done()
			stop()
		}()
	})
	return ctx
}

func BuildSubjectRecordResp(records []*subject.SubjectRecord) []*subject.SubjectRecordResp {
	results := make([]*subject.SubjectRecordResp, len(records))
	for i, record := range records {
		result := &subject.SubjectRecordResp{
			ID:           record.ID,
			Name:         record.Name,
			Answer:       record.Answer,
			SubjectType:  record.SubjectType,
			Phase:        record.Phase,
			LearnTimes:   record.LearnTimes,
			LastReviewAt: Ptr(time.Unix(record.GetLastReviewAt(), 0).Format("2006-01-02 15:04:05")),
			NextReviewAt: Ptr(time.Unix(record.GetNextReviewAt(), 0).Format("2006-01-02 15:04:05")),
		}
		results[i] = result
	}
	return results
}
