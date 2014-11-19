package authlib

import (
	"code.google.com/p/gorest"
	"duov6.com/common"
)

type AuthCertificate struct {
	UserID, Username, Name, Email, SecurityToken, Domain, DataCaps, ClientIP string
	//Otherdata                                                                map[string]interface{}
}

type Auth struct {
	gorest.RestService
	login        gorest.EndPoint `method:"GET" path:"/Login/{username:string}/{password:string}/{domain:string}" output:"AuthCertificate"`
	autherize    gorest.EndPoint `method:"GET" path:"/Autherize/{SecurityToken:string}/{ApplicationID:string}" output:"AuthCertificate"`
	getAuthCode  gorest.EndPoint `method:"GET" path:"/GetAuthCode/{SecurityToken:string}/{ApplicationID:string}" output:"string"`
	autherizeApp gorest.EndPoint `method:"GET" path:"/AutherizeApp/{SecurityToken:string}/{Code:string}/{ApplicationID:string}/{AppSecret:string}" output:"bool"`
	addUser      gorest.EndPoint `method:"POST" path:"/AddUser/" postdata:"User"`
	//addApplication gorest.EndPoint `method:"POST" path:"/AddApplication/" postdata:"applib.Application"`
	getUser gorest.EndPoint `method:"GET" path:"/GetUser/" output:"User"`
}

func (A Auth) Login(username, password, domain string) AuthCertificate {
	h := newAuthHandler()
	u, err := h.Login(username, password)
	if err == "" {
		//fmt.Println("login succeful")
		securityToken := common.GetGUID()
		c := AuthCertificate{u.UserID, u.EmailAddress, u.Name, u.EmailAddress, securityToken, "http://192.168.0.58:9000/instaltionpath", "#0so0936#sdasd", "IPhere"}
		h.AddSession(c)
		return c
	} else {
		return AuthCertificate{}
	}
}

func (A Auth) Autherize(SecurityToken string, ApplicationID string) AuthCertificate {
	h := newAuthHandler()
	var a AuthCertificate
	c, err := h.GetSession(SecurityToken)
	if err == "" {
		if h.AppAutherize(ApplicationID, c.UserID) == true {
			a = c
			a.SecurityToken = common.GetGUID()
			h.AddSession(a)
			return a
		}

	}
	return a
}

func (A Auth) GetAuthCode(SecurityToken, ApplicationID, URI string) string {
	h := newAuthHandler()
	c, err := h.GetSession(SecurityToken)
	if err == "" {
		return h.GetAuthCode(ApplicationID, c.UserID, URI)
	}
	return ""
}

func (A Auth) AutherizeApp(SecurityToken, Code, ApplicationID, AppSecret string) bool {
	return true
}

func (A Auth) AddUser(u User) {
	h := newAuthHandler()
	h.SaveUser(u)
}

func (A Auth) GetUser() User {
	return User{"", "", "", "", "", false}
}
