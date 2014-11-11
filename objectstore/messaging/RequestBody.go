package messaging

type RequestBody struct {
	Parameters ObjectParameters
	Query      Query
	Body       map[string]interface{}
}
