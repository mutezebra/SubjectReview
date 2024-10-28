package oss

import (
	"context"
	"fmt"
	"io"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/pkg/errors"

	"github.com/mutezebra/subject-review/config"
	"github.com/mutezebra/subject-review/pkg/logger"
)

var cfg *oss.Config

func initOSS() {
	if config.OSS == nil {
		logger.Fatal("config`s oss is nil")
	}
	region := config.OSS.Region
	provider := credentials.NewStaticCredentialsProvider(config.OSS.OssAccessKeyID, config.OSS.OssAccessKeySecret)
	cfg = oss.LoadDefaultConfig().WithCredentialsProvider(provider).WithRegion(region)
}

type moss struct {
	cfg *oss.Config
}

func NewOss() *moss {
	if cfg == nil {
		initOSS()
	}
	return &moss{cfg: cfg}
}

func (m *moss) UploadAvatar(ctx context.Context, filename string, reader io.Reader) error {
	c := oss.NewClient(m.cfg)
	if _, err := c.NewUploader().UploadFrom(ctx, &oss.PutObjectRequest{
		Bucket: oss.Ptr(config.OSS.Bucket),
		Key:    oss.Ptr(fmt.Sprintf("%s/%s", config.OSS.AvatarPrefix, filename)),
	}, reader); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("faled upload %s to oss,error: ", filename))
	}
	return nil
}
