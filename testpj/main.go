package main

import (
	"duov6.com/authlib"
	"duov6.com/term"
)

func main() {
	term.Write(term.FgRed, term.BgBlue, term.Read(term.FgBlack, term.BgCyan, "Mother Fucker"), term.ErrorBlink)
	a := authlib.Auth{}
	term.Writec(a.GetAuthCode("sss", "ssss"))
}
