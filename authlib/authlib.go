package authlib

import (
	"code.google.com/p/gorest"
	"duov6.com/common"
	//"duov6.com/config"
)

type AuthCertificate struct {
	UserID, Username, Name, Email, SecurityToken, Domain, DataContract string
}

//Authendication Struct
type Auth struct {
	gorest.RestService
	login        gorest.EndPoint `method:"GET" path:"/Login/{username:string}/{password:string}/{domain:string}" output:"AuthCertificate"`
	autherize    gorest.EndPoint `method:"GET" path:"/Autherize/{SecurityToken:string}/{ApplicationID:string}" output:"AuthCertificate"`
	getAuthCode  gorest.EndPoint `method:"GET" path:"/GetAuthCode/{SecurityToken:string}/{ApplicationID:string}" output:"string"`
	autherizeApp gorest.EndPoint `method:"GET" path:"/AutherizeApp/{SecurityToken:string}/{Code:string}/{ApplicationID:string}/{AppSecret:string}" output:"bool"`
	//Config config.File
}

//var A = Auth{}

func (A *Auth) NewAuth(username, password, domain string) AuthCertificate {
	if username == "admin" {
		//fmt.Println("login succeful")
		securityToken := common.GetGUID()

		return AuthCertificate{"0", "Admin", "Administrator", "lasitha.senanayake@gmail.com", securityToken, "http://192.168.0.58:9000/instaltionpath", "0so0936"}

	} else {

		return AuthCertificate{}
	}
}

func (A *Auth) Autherize(SecurityToken string, ApplicationID string) AuthCertificate {
	return AuthCertificate{"0", "Admin", "Administrator", "lasitha.senanayake@gmail.com", "SecurityToke", "http://192.168.0.58:9000/instaltionpath", "0so0936"}

}

func (A Auth) GetAuthCode(SecurityToken, ApplicationID string) string {
	return "12233"

}

func (A *Auth) AutherizeApp(SecurityToken, Code, ApplicationID, AppSecret string) bool {
	return true

}

//Auth end
