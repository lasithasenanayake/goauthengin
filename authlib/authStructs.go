package authlib

import (
//"duov6.com/config"
)

func NewUser(userID, EmailAddress, Name, Password string) User {
	return User{userID, EmailAddress, Name, Password, Password, false}
}

func GetConfig() AuthConfig {
	return AuthConfig{Cirtifcate: "", PrivateKey: ""}
}

func SetConfig(c AuthConfig) {
	//c.PrivateKey
}

type AppAutherize struct {
	Name          string
	AppliccatioID string
	AutherizeKey  string
	OtherData     map[string]interface{}
}

type AppCertificate struct {
	AuthKey       string
	UserID        string
	ApplicationID string
	AppSecretKey  string
	Otherdata     map[string]interface{}
}

type User struct {
	UserID          string
	EmailAddress    string
	Name            string
	Password        string
	ConfirmPassword string
	Active          bool
}

type AuthConfig struct {
	Cirtifcate    string
	PrivateKey    string
	Https_Enabled bool
	StoreID       string
	Smtpserver    string
	Smtpusername  string
	Smtppassword  string
}

type AuthCode struct {
	ApplicationID string
	Code          string
	UserID        string
	URI           string
}
