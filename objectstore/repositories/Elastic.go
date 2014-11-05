package repositories

import (
	"duov6.com/objectstore/messaging"
	"fmt"
)

type ElasticRepository struct {
}

func (repository ElasticRepository) GetAll(request *messaging.ObjectRequest) RepositoryResponse {
	return getDefaultNotImplemented()
}

func (repository ElasticRepository) GetSearch(request *messaging.ObjectRequest) RepositoryResponse {
	return getDefaultNotImplemented()
}

func (repository ElasticRepository) GetQuery(request *messaging.ObjectRequest) RepositoryResponse {
	return getDefaultNotImplemented()
}

func (repository ElasticRepository) GetByKey(request *messaging.ObjectRequest) RepositoryResponse {
	return getDefaultNotImplemented()
}

func (repository ElasticRepository) InsertMultiple(request *messaging.ObjectRequest) RepositoryResponse {
	fmt.Println("ELASTIC Insert Multiple")
	return getDefaultNotImplemented()
}

func (repository ElasticRepository) InsertSingle(request *messaging.ObjectRequest) RepositoryResponse {
	fmt.Println("ELASTIC Insert Single")
	return getDefaultNotImplemented()
}

func (repository ElasticRepository) UpdateMultiple(request *messaging.ObjectRequest) RepositoryResponse {
	return getDefaultNotImplemented()
}

func (repository ElasticRepository) UpdateSingle(request *messaging.ObjectRequest) RepositoryResponse {
	return getDefaultNotImplemented()
}

func (repository ElasticRepository) DeleteMultiple(request *messaging.ObjectRequest) RepositoryResponse {
	return getDefaultNotImplemented()
}

func (repository ElasticRepository) DeleteSingle(request *messaging.ObjectRequest) RepositoryResponse {
	return getDefaultNotImplemented()
}

func (repository ElasticRepository) Special(request *messaging.ObjectRequest) RepositoryResponse {
	return getDefaultNotImplemented()
}

func (repository ElasticRepository) Test(request *messaging.ObjectRequest) {

}
