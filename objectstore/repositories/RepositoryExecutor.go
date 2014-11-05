package repositories

import (
	"duov6.com/objectstore/messaging"
)

func Execute(request *messaging.ObjectRequest, repository AbstractRepository) (response RepositoryResponse) {

	switch request.Header.Operation { //CREATE, READ, UPDATE, DELETE, SPECIAL
	case "CREATE":
		if request.Header.Multiplicity == "SINGLE" {
			response = repository.InsertSingle(request)
		} else {
			response = repository.InsertMultiple(request)
		}
	case "READ":
		switch request.Body.Query.What { //QUERYING, SEARCHING, KEY, ALL
		case "QUERYING":
			response = repository.GetQuery(request)
		case "SEARCHING":
			response = repository.GetSearch(request)
		case "KEY":
			response = repository.GetByKey(request)
		case "ALL":
			response = repository.GetAll(request)
		}
	case "UPDATE":
		if request.Header.Multiplicity == "SINGLE" {
			response = repository.UpdateSingle(request)
		} else {
			response = repository.UpdateMultiple(request)
		}
	case "DELETE":
		if request.Header.Multiplicity == "SINGLE" {
			response = repository.DeleteSingle(request)
		} else {
			response = repository.DeleteMultiple(request)
		}
	case "SPECIAL":
		response = repository.Special(request)
	}

	return
}
