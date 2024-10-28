package errno

// Verify
const (
	InvalidEmailFormat = 1000 + iota
	InvalidUserNameFormat
	InvalidPasswordFormat
	InvalidVerifyCodeFormat
	SpaceUserName
	LackToken
	WrongToken
	TokenExpire
	LackOFAvatarFile
	InvalidParam
)

// User
const (
	EmailHaveExisted = 2000 + iota
	ExpiredVerifyCode
	WrongVerifyCode
	UserNameExisted
	WrongPassword
	UnsupportedAvatarFormat
	AvatarFileOverLimit
	AvatarFileTooSmall
	EmailNotExisted
)

// subject
const (
	PagesLessThanZero = 3000 + iota
	SizeLessThanZero
	UnknownSubjectType
	UnknownSubjectReqType
	LengthOFSubjectNameOverLimit
	SubjectNotExist
	NotManager
	SubjectIDLessThanZero
)
