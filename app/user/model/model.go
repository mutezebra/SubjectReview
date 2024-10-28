package model

import (
	"context"
	"io"
)

type Email interface {
	SendVerifyCode(to string, message *string) error
}

type OSS interface {
	UploadAvatar(ctx context.Context, filename string, reader io.Reader) error
}
