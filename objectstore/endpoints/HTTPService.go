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

	//READ BY KEY
	m.Get("/:namespace/:class/:id", handleRequest)
	//READ BY KEYWORD
	m.Get("/:namespace/:class?keyword=:keyword", handleRequest)
	//READ ADVANCED, INSERT
	m.Post("/:namespace/:class", handleRequest)

	//UPDATE
	m.Put("/:namespace/:class", handleRequest)
	//DELETE
	m.Delete("/:namespace/:class", handleRequest)

	m.Run()
}

func handleRequest(params martini.Params, res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini

	responseMessage, isSuccess := dispatchRequest(req, params)

	if isSuccess {
		res.WriteHeader(200)
		fmt.Println("Success")
	} else {
		res.WriteHeader(500)
		fmt.Println("Failed!!!")
	}

	fmt.Fprintf(res, "%s", responseMessage)
}

func (h *HTTPService) Stop() {
}

func dispatchRequest(r *http.Request, params martini.Params) (responseMessage string, isSuccess bool) { //result is JSON

	objectRequest := messaging.ObjectRequest{}

	paramMap := make(map[string]interface{})
	objectRequest.Extras = paramMap

	message, isSuccess := getObjectRequest(r, &objectRequest, params)

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

func getObjectRequest(r *http.Request, objectRequest *messaging.ObjectRequest, params martini.Params) (message string, isSuccess bool) {

	missingFields := ""
	isSuccess = true

	headerToken := r.Header.Get("securityToken")

	var headerOperation string
	headerMultipliciry := r.Header.Get("multiplicity")

	headerNamespace := params["namespace"]
	headerClass := params["class"]

	headerId := params["id"]
	headerKeyword := params["keyword"]

	if len(headerToken) == 0 {
		isSuccess = false
		missingFields = missingFields + "securityToken"
	}

	var requestBody messaging.RequestBody

	if isSuccess {

		if r.Method != "GET" {
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
				}
			}
		}

		if isSuccess {

			canAddHeader := true
			switch r.Method {
			case "GET": //read keyword, and unique key
				if len(headerId) != 0 {
					headerOperation = "read-key"
				} else if len(headerKeyword) != 0 {
					headerOperation = "read-keyword"
				} else if len(headerNamespace) != 0 && len(headerClass) != 0 {
					headerOperation = "read-all"
				}
				canAddHeader = false
			case "POST": //read query, read special, insert
				if requestBody.Body != nil {
					fmt.Println("Inset by POST : " + objectRequest.Body.Parameters.KeyProperty)
					headerOperation = "insert"
					headerId = objectRequest.Body.Body[objectRequest.Body.Parameters.KeyProperty].(string)
				} else if &requestBody.Query != nil {
					headerOperation = "read-filter"
					canAddHeader = false
				}

			case "PUT": //update
				headerId = objectRequest.Body.Body[objectRequest.Body.Parameters.KeyProperty].(string)
				headerOperation = "update"

			case "DELETE": //delete
				headerId = objectRequest.Body.Body[objectRequest.Body.Parameters.KeyProperty].(string)
				headerOperation = "delete"
			}

			headerMultipliciry = "single"

			objectRequest.Controls = messaging.RequestControls{SecurityToken: headerToken, Namespace: headerNamespace, Class: headerClass, Multiplicity: headerMultipliciry, Id: headerId, Operation: headerOperation}

			configObject := configuration.ConfigurationManager{}.Get(headerToken, headerNamespace, headerClass)
			objectRequest.Configuration = configObject

			if canAddHeader {
				repositories.FillControlHeaders(objectRequest)
			}
		}

	} else {
		message = "Missing attributes in request header : " + missingFields
	}

	return
}
