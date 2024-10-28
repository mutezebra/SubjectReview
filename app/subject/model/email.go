package model

type Email interface {
	SendRemindMessage(to string, msg *string) error
}
