package Sample

import (
	"GoClean/GoClean.Application/Contract/Services"
)

type SampleBusinessServiceHandler struct {
	MessageService *Services.IMessageService
}

func NewSampleBusinessServiceHandler(messageService *Services.IMessageService) *SampleBusinessServiceHandler {
	return &SampleBusinessServiceHandler{
		MessageService: messageService,
	}
}

type SampleBusinessServiceInput struct {
	Field string
}
type SampleBusinessServiceOutput struct {
}

func (receiver SampleBusinessServiceHandler) Handle(input *SampleBusinessServiceInput) *SampleBusinessServiceOutput {

	////Error Generation
	//responseGenerator := Commons.NewResponseGenerator(receiver.MessageService)
	//responseObject := responseGenerator.GetResponseObjectResult(Commons.RESULT_CODE_VALIDATION_ERROR, nil)
	//panic(&responseObject)

	return &SampleBusinessServiceOutput{}
}
