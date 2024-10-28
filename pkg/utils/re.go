package utils

import "regexp"

func GetEmailVerifyRe() *regexp.Regexp {
	return regexp.MustCompile(`^[A-Za-z0-9\x{4e00}-\x{9fa5}]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
}

func GetPasswordVerifyRe() *regexp.Regexp {
	return regexp.MustCompile(`^[a-zA-Z0-9_.]+$`) // password pattern (without checks for lower/upper case)
}
