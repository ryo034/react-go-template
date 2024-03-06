package error

type BadRequest struct {
	msg string
}

func NewBadRequest(msg string) *BadRequest {
	return &BadRequest{msg}
}

func (e *BadRequest) Error() string {
	return e.msg
}

func (e *BadRequest) MessageKey() MessageKey {
	return BadRequestMessageKey
}
