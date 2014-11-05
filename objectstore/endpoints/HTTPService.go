package endpoints

import (
	"duov6.com/objectstore/configuration"
	"duov6.com/objectstore/messaging"
	"duov6.com/objectstore/repositories"
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"io/ioutil"
	"net/http"
)

type HTTPService struct {
}

func (h *HTTPService) Start() {
	fmt.Println("Object Store Listening on Port : 8090")
	m := martini.Classic()
	m.Post("/:namespace/:class", func(params martini.Params, res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini

		responseMessage, isSuccess := dispatchRequest(req, params["namespace"], params["class"])

		if isSuccess {
			res.WriteHeader(200)
			fmt.Println("Success")
		} else {
			res.WriteHeader(500)
			fmt.Println("Failed!!!")
		}

		fmt.Fprintf(res, "%s", responseMessage)
	})
	m.Run()
}

func (h *HTTPService) Stop() {
}

func dispatchRequest(r *http.Request, namespace string, class string) (responseMessage string, isSuccess bool) { //result is JSON

	objectRequest := messaging.ObjectRequest{}
	message, isSuccess := getObjectRequest(r, &objectRequest, namespace, class)

	if isSuccess == false {
		responseMessage = getQueryResponseString("Invalid Query Request", message, false)
	} else {

		dispatcher := Dispatcher{}
		var repResponse repositories.RepositoryResponse = dispatcher.Dispatch(&objectRequest)
		isSuccess = repResponse.IsSuccess

		if isSuccess {
			if repResponse.Body != nil {
				responseMessage = string(repResponse.Body)
			} else {
				responseMessage = getQueryResponseString("Successfully completed request", repResponse.Message, isSuccess)
			}

		} else {
			responseMessage = getQueryResponseString("Error occured while processing", repResponse.Message, isSuccess)
		}

	}

	return
}

func getQueryResponseString(mainError string, reason string, isSuccess bool) string {
	response := messaging.ResponseBody{}
	response.IsSuccess = isSuccess
	response.Message = mainError + " : " + reason

	result, err := json.Marshal(&response)

	if err == nil {
		return string(result)
	} else {
		return "Invalid Query"
	}
}

func getObjectRequest(r *http.Request, objectRequest *messaging.ObjectRequest, headerNamespace string, headerClass string) (message string, isSuccess bool) {

	missingFields := ""
	isSuccess = true

	headerToken := r.Header.Get("securityToken")
	headerOperation := r.Header.Get("operation")
	headerMultipliciry := r.Header.Get("multiplicity")

	headerId := r.Header.Get("id")
	headerVersion := r.Header.Get("version")

	if len(headerToken) == 0 {
		isSuccess = false
		missingFields = missingFields + "securityToken, "
	}

	if len(headerOperation) == 0 {
		isSuccess = false
		missingFields = missingFields + "operation, "
	}

	if len(headerMultipliciry) == 0 {
		isSuccess = false
		missingFields = missingFields + "multiplicity, "
	}

	if isSuccess {
		objectRequest.Header = messaging.RequestHeader{SecurityToken: headerToken, Namespace: headerNamespace, Class: headerClass, Operation: headerOperation, Multiplicity: headerMultipliciry, Id: headerId, Version: headerVersion}

		var requestBody messaging.RequestBody

		rb, rerr := ioutil.ReadAll(r.Body)

		if rerr != nil {
			message = "Error converting request : " + rerr.Error()
			isSuccess = false
		} else {

			err := json.Unmarshal(rb, &requestBody)

			if err != nil {
				message = "JSON Parse error in Request : " + err.Error()
				isSuccess = false
			} else {
				objectRequest.Body = requestBody

				configObject := configuration.ConfigurationManager{}.Get(headerToken, headerNamespace, headerClass)
				objectRequest.StoreConfiguration = configObject
			}
		}

	} else {
		message = "Missing attributes in request header : " + missingFields
	}

	return
}
