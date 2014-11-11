package repositories

import (
	"duov6.com/objectstore/messaging"
	"fmt"
)

func getDefaultNotImplemented() RepositoryResponse {
	return RepositoryResponse{IsSuccess: false, IsImplemented: false, Message: "Operation Not Implemented"}
}

func FillControlHeaders(request *messaging.ObjectRequest) {
	controlObject := messaging.ControlHeaders{}
	controlObject.Version = "xxx-xxx-xxx-xxx"
	controlObject.Namespace = request.Controls.Namespace
	controlObject.Class = request.Controls.Class
	controlObject.Tenant = "123"
	controlObject.LastUdated = "xx/xx/xxxx xx:xx:xx"

	fmt.Println("Filling Headers...")
	request.Body.Body["__osHeaders"] = controlObject
	fmt.Println("Headers Filled...")
}

func getNoSqlKey(request *messaging.ObjectRequest) string {
	key := request.Controls.Namespace + "." + request.Controls.Class + "." + request.Controls.Id
	return key
}
