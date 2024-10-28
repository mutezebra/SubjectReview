package errno

type Errno interface {
	Code() int32
	Error() string
}

type errno struct {
	code int32
	msg  string
}

func (e *errno) Code() int32 {
	return e.code
}

func (e *errno) Error() string {
	return e.msg
}

func New(code int32, msg string) Errno {
	return &errno{
		code: code,
		msg:  msg,
	}
}
