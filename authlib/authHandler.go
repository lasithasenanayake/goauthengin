package authlib

import (
	"duov6.com/common"
	"duov6.com/config"
	"duov6.com/objectstore/client"
	"duov6.com/term"
	"encoding/json"
	//"fmt"
)

type AuthHandler struct {
	Config config.File
}

func newAuthHandler() *AuthHandler {
	authhld := new(AuthHandler)
	authhld.Config = config.File{Filename: "auth.cofig"}
	return authhld
}

func (h *AuthHandler) ChangePassword() {

}

func (h *AuthHandler) AppAutherize(ApplicationID, UserID string) bool {
	bytes, err := client.Go("ignore", "com.duosoftware.auth", "Atherized").GetOne().BySearching(ApplicationID + "-" + UserID).Ok()
	term.Write("AppAutherize For Application "+ApplicationID+" UserID "+UserID, term.Debug)
	if err == "" {
		if bytes != nil {
			var uList []AppAutherize
			err := json.Unmarshal(bytes, &uList)
			if err == nil && len(uList) != 0 {
				return true
			}
		}
	} else {
		term.Write("AppAutherize Error "+err, term.Error)
	}
	return false
}

func (h *AuthHandler) GetAuthCode(ApplicationID, UserID, URI string) string {
	var a AuthCode
	a.ApplicationID = ApplicationID
	a.UserID = UserID
	a.URI = URI
	a.Code = common.RandText(10)
	client.Go("ignore", "com.duosoftware.auth", "AuthCode").StoreObject().WithKeyField("Code").AndStoreOne(a).Ok()
	term.Write("GetAuthCode for "+ApplicationID+" with SecurityToken :"+UserID, term.Debug)
	return a.Code
}

func AutherizeApp(Code, ApplicationID, AppSecret, UserID string) bool {
	bytes, _ := client.Go("ignore", "com.duosoftware.auth", "AuthCode").GetOne().BySearching(Code).Ok()
	term.Write("AutherizeApp For SecurityToken "+ApplicationID, term.Debug)
	var uList []AuthCode
	json.Unmarshal(bytes, &uList)
	if len(uList) != 0 {

		return true
	}
	return false
}

func (h *AuthHandler) AddSession(a AuthCertificate) {
	client.Go("ignore", "com.duosoftware.auth", "sessions").StoreObject().WithKeyField("SecurityToken").AndStoreOne(a).Ok()
	term.Write("AddSession for "+a.Name+" with SecurityToken :"+a.SecurityToken, term.Debug)
}

func (h *AuthHandler) GetSession(key string) (AuthCertificate, string) {
	bytes, err := client.Go("ignore", "com.duosoftware.auth", "sessions").GetOne().BySearching(key).Ok()
	term.Write("GetSession For SecurityToken "+key, term.Debug)

	var c AuthCertificate
	if err == "" {
		if bytes != nil {
			var uList []AuthCertificate
			err := json.Unmarshal(bytes, &uList)
			if err == nil && len(uList) != 0 {
				return uList[0], ""
			}
		}
	} else {
		term.Write("GetSession Error "+err, term.Error)
	}
	term.Write("GetSession No Session for SecurityToken "+key, term.Debug)

	return c, "Error Session Not Found"
}

func (h *AuthHandler) SaveUser(u User) User {
	term.Write("SaveUser saving user  "+u.Name, term.Debug)

	bytes, err := client.Go("ignore", "com.duosoftware.auth", "users").GetOne().BySearching(u.EmailAddress).Ok()
	if err == "" {
		var uList []User
		err := json.Unmarshal(bytes, &uList)
		if err == nil {
			if len(uList) == 0 {

				u.UserID = common.GetGUID()
				term.Write("SaveUser saving user  "+u.Name+" New User "+u.UserID, term.Debug)

				client.Go("ignore", "com.duosoftware.auth", "users").StoreObject().WithKeyField("EmailAddress").AndStoreOne(u).Ok()
			} else {
				u.UserID = uList[0].UserID
				term.Write("SaveUser saving user  "+u.Name+" Update User "+u.UserID, term.Debug)
				client.Go("ignore", "com.duosoftware.auth", "users").StoreObject().WithKeyField("EmailAddress").AndStoreOne(u).Ok()
			}
		} else {
			term.Write("SaveUser saving user store Error #"+err.Error(), term.Error)
		}
	} else {
		term.Write("SaveUser saving user fetech Error #"+err, term.Error)
	}
	return u
}

func (h *AuthHandler) Login(email, password string) (User, string) {
	term.Write("Login  user  email"+email, term.Debug)
	bytes, err := client.Go("ignore", "com.duosoftware.auth", "users").GetOne().BySearching(email).Ok()
	var user User
	if err == "" {
		if bytes != nil {
			var uList []User
			err := json.Unmarshal(bytes, &uList)

			if err == nil && len(uList) != 0 {
				if uList[0].Password == password && uList[0].EmailAddress == email {
					return uList[0], ""
				} else {
					term.Write("password incorrect", term.Error)
				}
			} else {
				if err != nil {
					term.Write("Login  user Error "+err.Error(), term.Error)
				}
			}
		}
	} else {
		term.Write("Login  user  Error "+err, term.Error)
	}

	return user, "Error Validating user"
}

func SendNotification(u User, Message string) {

}
