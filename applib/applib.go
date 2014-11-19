package applib

import (
	"duov6.com/common"
	"duov6.com/objectstore/client"
	"duov6.com/term"
	"encoding/json"
)

func NewApphanler() AppSvc {
	var apphdl AppSvc
	return apphdl
}

type Application struct {
	ApplicationID string
	SecretKey     string
	Name          string
	Description   string
	AppType       string
	AppUri        string
	OtherData     map[string]interface{}
}

type AppSvc struct {
}

func (app AppSvc) Get(ApplicationID string) Application {
	term.Write("Get  App  by ID"+ApplicationID, term.Debug)
	bytes, err := client.Go("ignore", "com.duosoftware.application", "apps").GetOne().BySearching(ApplicationID).Ok()
	var a Application
	if err == "" {
		if bytes != nil {
			var uList []Application
			err := json.Unmarshal(bytes, &uList)

			if err == nil && len(uList) != 0 {
				return uList[0]
			} else {
				if err != nil {
					term.Write("Login  user Error "+err.Error(), term.Error)
				}
			}
		}
	} else {
		term.Write("Login  user  Error "+err, term.Error)
	}

	return a
}

func (app AppSvc) Add(a Application) Application {
	term.Write("Add saving Application  "+a.Name, term.Debug)

	bytes, err := client.Go("ignore", "com.duosoftware.application", "apps").GetOne().BySearching(a.ApplicationID).Ok()
	if err == "" {
		var uList []Application
		err := json.Unmarshal(bytes, &uList)
		if err == nil {
			if len(uList) == 0 {

				a.ApplicationID = common.GetGUID()
				a.SecretKey = common.RandText(10)
				term.Write("Add saving Addaplication  "+a.Name+" New App "+a.ApplicationID, term.Debug)

				client.Go("ignore", "com.duosoftware.Application", "apps").StoreObject().WithKeyField("ApplicationID").AndStoreOne(a).Ok()
			} else {
				a.ApplicationID = uList[0].ApplicationID
				term.Write("SaveUser saving user  "+a.Name+" Update User "+a.ApplicationID, term.Debug)
				client.Go("ignore", "com.duosoftware.Application", "apps").StoreObject().WithKeyField("ApplicationID").AndStoreOne(a).Ok()
			}
		} else {
			term.Write("SaveUser saving user store Error #"+err.Error(), term.Error)
		}
	} else {
		term.Write("SaveUser saving user fetech Error #"+err, term.Error)
	}
	return a
}
