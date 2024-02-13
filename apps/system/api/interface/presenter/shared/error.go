package shared

type ErrorResponse struct {
	Title  string      `json:"title"`
	Detail string      `json:"detail"`
	Type   string      `json:"type"`
	Errors []errDetail `json:"errors"`
	Code   string      `json:"code"`
}

type errDetail struct {
	Message string `json:"message"`
}
