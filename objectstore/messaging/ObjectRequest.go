package messaging

import (
	"duov6.com/objectstore/configuration"
)

type ObjectRequest struct {
	Header             RequestHeader
	Body               RequestBody
	StoreConfiguration configuration.StoreConfiguration
}
