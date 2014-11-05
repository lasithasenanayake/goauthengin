package storageengines

import (
	"duov6.com/objectstore/messaging"
	"duov6.com/objectstore/repositories"
	"fmt"
)

type ReplicatedStorageEngine struct {
}

func (r ReplicatedStorageEngine) Store(request *messaging.ObjectRequest) (response repositories.RepositoryResponse) {

	//1 = COMMIT, 2 = ROLLBACK, 3 = BREAK
	var successAction int = 0
	var failAction int = 0
	var engineMappings map[string]string

	switch request.Header.Operation { //CREATE, READ, UPDATE, DELETE, SPECIAL
	case "CREATE":
		successAction = 1
		failAction = 2
		if request.Header.Multiplicity == "SINGLE" {
			engineMappings = request.StoreConfiguration.StoreConfiguration["INSERT-SINGLE"]
		} else {
			engineMappings = request.StoreConfiguration.StoreConfiguration["INSERT-MULTIPLE"]
		}
	case "READ":
		successAction = 3
		failAction = 1
		switch request.Body.Query.What { //QUERYING, SEARCHING, KEY, ALL
		case "QUERYING":
			engineMappings = request.StoreConfiguration.StoreConfiguration["GET-QUERY"]
		case "SEARCHING":
			engineMappings = request.StoreConfiguration.StoreConfiguration["GET-SEARCH"]
		case "KEY":
			engineMappings = request.StoreConfiguration.StoreConfiguration["GET-KEY"]
		case "ALL":
			engineMappings = request.StoreConfiguration.StoreConfiguration["GET-ALL"]
		}
	case "UPDATE":
		successAction = 1
		failAction = 2
		if request.Header.Multiplicity == "SINGLE" {
			engineMappings = request.StoreConfiguration.StoreConfiguration["UPDATE-SINGLE"]
		} else {
			engineMappings = request.StoreConfiguration.StoreConfiguration["UPDATE-MULTIPLE"]
		}
	case "DELETE":
		successAction = 1
		failAction = 2
		if request.Header.Multiplicity == "SINGLE" {
			engineMappings = request.StoreConfiguration.StoreConfiguration["DELETE-SINGLE"]
		} else {
			engineMappings = request.StoreConfiguration.StoreConfiguration["DELETE-MULTIPLE"]
		}
	case "SPECIAL":
		successAction = 3
		failAction = 1
		engineMappings = request.StoreConfiguration.StoreConfiguration["SPECIAL"]

	}

	convertedRepositories := getRepositories(engineMappings)

	response = startAtomicOperation(request, convertedRepositories, successAction, failAction)

	return
}

func getRepositories(engineMappings map[string]string) []repositories.AbstractRepository {
	var outRepositories []repositories.AbstractRepository

	outRepositories = make([]repositories.AbstractRepository, len(engineMappings))

	count := -1

	for k, v := range engineMappings {
		count++
		fmt.Println(k)
		absRepository := repositories.Create(v)
		outRepositories[count] = absRepository
	}

	return outRepositories
}

func startAtomicOperation(request *messaging.ObjectRequest, repositoryList []repositories.AbstractRepository, successAction int, failAction int) (response repositories.RepositoryResponse) {

	canRollback := false
	for index, repository := range repositoryList {
		fmt.Println(index)
		tmpResponse := repositories.Execute(request, repository)
		canBreak := false

		if tmpResponse.IsSuccess {
			switch successAction {
			case 1:
				response = tmpResponse
				continue
			case 3:
				response = tmpResponse
				canBreak = true
			}
		} else {
			switch failAction {
			case 1:
				continue
			case 2:
				canRollback = true
				canBreak = true
			case 3:
				response = tmpResponse
				canBreak = true
			}
		}

		if canBreak == true {
			break
		}

		fmt.Println(tmpResponse.Message)

		//1 = COMMIT, 2 = ROLLBACK, 3 = BREAK

	}

	if canRollback {
		fmt.Println("Rollbacking!!!")
	}
	return
}
