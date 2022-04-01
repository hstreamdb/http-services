package errorno

type ErrorResponse struct {
	Code     uint32 `json:"code"`
	Message  string `json:"message"`
	FullText string `json:"full_text"`
}

func NewErrorResponse(code ErrorCode, err error) ErrorResponse {
	return ErrorResponse{
		Code:     uint32(code),
		Message:  GetMsg(code),
		FullText: err.Error(),
	}
}
