package authlib

import ()

type AuthCertificate struct {
	UserID, Username, Name, Email, SecurityToken string
}

//Authendication Struct
type Auth struct {
}

//var A = Auth{}

func (A Auth) Login(username, password, domain string) (AuthCertificate, int) {
	if username == "admin" {
		//fmt.Println("login succeful")
		return AuthCertificate{"UserID", "UserName", "Name", "Email", "SecurityToke"}, 1
	} else {

		return AuthCertificate{}, 0
	}
}

func (A Auth) Autherize(SecurityToken string, ApplicationIDs []string) AuthCertificate {
	return AuthCertificate{"UserID", "UserName", "Name", "Email", "SecurityToke"}
}

func (A Auth) GetAuthCode(SecurityToken, ApplicationID string) string {
	return "12233"
}

//Auth end
