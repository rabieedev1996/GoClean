package SMSIR

type SMSIRSendTemplateRequest struct {
	Mobile     string                       `json:"mobile"`
	TemplateId int64                        `json:"templateId"`
	Parameters []SMSIRSendTemplateParameter `json:"parameters"`
}

type SMSIRSendTemplateParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SMSIRSendResponse struct {
	Status  int64                 `json:"status"`
	Message string                `json:"message"`
	Data    SMSIRSendResponseData `json:"data"`
}

type SMSIRSendResponseData struct {
	MessageId int64 `json:"messageId"`
	Cost      int64 `json:"cost"`
}

type SMSIRSendRequest struct {
	LineNumber  string   `json:"lineNumber"`
	MessageText string   `json:"messageText"`
	Mobiles     []string `json:"mobiles"`
}
