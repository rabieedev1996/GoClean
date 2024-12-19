package Commons

import (
	Services2 "GoClean/GoClean.Application/Contract/Services"
	GoClean_Common "GoClean/GoClean.Common"
	"GoClean/GoClean.Domain/Enums"
	"net/http"
)

type ResponseModel struct {
	ResponseDetil  ResponseDetail
	HTTPStatusCode int `json:"httpStatusCode"`
}
type ResponseDetail struct {
	ResultCode    int                          `json:"ResultCode"`
	Data          any                          `json:"Data"`
	Message       string                       `json:"Message"`
	IsSuccess     bool                         `json:"IsSuccess"`
	ErrorMessages GoClean_Common.Slice[string] `json:"ErrorMessages"`
}

const (
	RESULT_CODE_SUCCESS int = iota
	RESULT_CODE_EXCEPTION
	RESULT_CODE_VALIDATION_ERROR
	RESULT_CODE_INVALID_OPERATION
	RESULT_CODE_INVALID_ACTIVATION_CODE
	RESULT_CODE_SIMILAR_USERNAME_EXIST
	RESULT_CODE_UNAUTHORIZED
)

type ResponseGenerator struct {
	messageService *Services2.IMessageService
}

func NewResponseGenerator(messageService *Services2.IMessageService) *ResponseGenerator {
	return &ResponseGenerator{messageService}
}

func (rg *ResponseGenerator) GetResponseObjectResult(code int, errors GoClean_Common.Slice[string], data ...any) ResponseModel {
	message := rg.fillMessage(code)
	httpStatusCode := rg.fillStatusCode(code)
	isSuccess := rg.fillSuccessStatus(code)
	var resData any = nil
	if len(data) != 0 {
		resData = data[0]
	}

	return ResponseModel{
		ResponseDetil: ResponseDetail{
			ResultCode:    int(code),
			Data:          resData,
			Message:       message,
			IsSuccess:     isSuccess,
			ErrorMessages: errors,
		},
		HTTPStatusCode: httpStatusCode,
	}
}

func (rg *ResponseGenerator) fillStatusCode(code int) int {
	switch code {
	case RESULT_CODE_SUCCESS:
		return http.StatusOK
	case RESULT_CODE_EXCEPTION:
		return http.StatusInternalServerError
	case RESULT_CODE_VALIDATION_ERROR, RESULT_CODE_INVALID_OPERATION:
		return http.StatusBadRequest
	case RESULT_CODE_UNAUTHORIZED:
		return http.StatusUnauthorized
	case RESULT_CODE_INVALID_ACTIVATION_CODE, RESULT_CODE_SIMILAR_USERNAME_EXIST:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}

func (rg *ResponseGenerator) fillSuccessStatus(code int) bool {
	switch code {
	case RESULT_CODE_SUCCESS:
		return true
	default:
		return false
	}
}

func (rg *ResponseGenerator) fillMessage(code int) string {
	switch code {
	case RESULT_CODE_SUCCESS:
		return (*rg.messageService).GetMessage(Enums.MESSAGE_CODE_SUCCESS)
	case RESULT_CODE_EXCEPTION:
		return (*rg.messageService).GetMessage(Enums.MESSAGE_CODE_EXCEPTION)
	case RESULT_CODE_VALIDATION_ERROR:
		return (*rg.messageService).GetMessage(Enums.MESSAGE_CODE_VALIDATION_ERROR)
	case RESULT_CODE_INVALID_ACTIVATION_CODE:
		return (*rg.messageService).GetMessage(Enums.MESSAGE_CODE_INVALID_ACTIVATION)
	case RESULT_CODE_INVALID_OPERATION:
		return (*rg.messageService).GetMessage(Enums.MESSAGE_CODE_INVALID_OPERATION)
	case RESULT_CODE_SIMILAR_USERNAME_EXIST:
		return (*rg.messageService).GetMessage(Enums.MESSAGE_CODE_SIMILAR_USERNAME_EXIST)
	case RESULT_CODE_UNAUTHORIZED:
		return (*rg.messageService).GetMessage(Enums.MESSAGE_CODE_UNAUTHORIZED)
	default:
		return "Unknown error"
	}
}
