package common

import (
	"crypto/md5"
	"encoding/hex"
	"os/exec"
)

func GetGUID() string {

	//h.Write()
	out, err := exec.Command("uuidgen").Output()
	h := md5.New()
	h.Write(out)
	if err == nil {
		return hex.EncodeToString(h.Sum(nil))
	} else {
		return ""
	}
}

func GetHash(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
