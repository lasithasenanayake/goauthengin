package repositories

func getDefaultNotImplemented()(RepositoryResponse){
	 return RepositoryResponse {IsSuccess:false,IsImplemented:false,Message:"Operation Not Implemented" };
}
