package config

import "testing"

func TestInitConfig(t *testing.T) {
	InitConfig()
	t.Logf("Email: \n")
	t.Log(Email.UserName, Email.Port, Email.UserName, Email.Secret)
	t.Logf("\n")
}
