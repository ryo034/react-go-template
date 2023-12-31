package response

import (
	"google.golang.org/protobuf/proto"
)

type ErrorResponse struct {
	Errors []errDetail `json:"errors"`
}

type errDetail struct {
	//Message string `json:"message"`
	Message proto.Message
}
