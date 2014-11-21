package main

import (
	"code.google.com/p/gorest"
	"duov6.com/applib"
	"duov6.com/authlib"
	"duov6.com/term"
	"net/http"
)

func main() {
	authlib.SetupConfig()
	//authlib.GetConfig()
	//term.Write(authlib.Config.Cirtifcate, term.Splash)
	go webServer()
	go runRestFul()
	term.SplashScreen("splash.art")
	term.Write("================================================================", term.Splash)
	term.Write("|     Admintration Console running on :9000                    |", term.Splash)
	term.Write("|     https RestFul Service running on :3048                   |", term.Splash)
	term.Write("|     Duo v6 Auth Service 6.0                                  |", term.Splash)
	term.Write("================================================================", term.Splash)
	term.StartCommandLine()

}

func status() {
	term.Write("Status is running", term.Information)
}

func webServer() {
	http.Handle(
		"/",
		http.StripPrefix(
			"/",
			http.FileServer(http.Dir("html")),
		),
	)
	http.ListenAndServe(":9000", nil)
}

func runRestFul() {
	gorest.RegisterService(new(authlib.Auth))
	gorest.RegisterService(new(applib.AppSvc))
	err := http.ListenAndServeTLS(":3048", "apache.crt", "apache.key", gorest.Handle())
	if err != nil {
		term.Write(err.Error(), term.Error)
		return
	}
}

func Setup() {

}
