package messaging

type RequestBody struct {
	Query Query
	Body  map[string]interface{}
}
