package SMSIR

import (
	GoClean_Common "GoClean/GoClean.Common"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type SMSIRMethods struct {
	Token            string
	BaseUrl          string
	TemplateId       int
	Sender           string
	PositionVariable string
}

func NewSMSIRMethods(token string) SMSIRMethods {
	sMSIRMethods := SMSIRMethods{
		Sender:           "30007732003736",
		Token:            token,
		BaseUrl:          "https://api.sms.ir/v1/send",
		TemplateId:       638137,
		PositionVariable: "CODE",
	}
	return sMSIRMethods
}
func (r SMSIRMethods) Send(dest GoClean_Common.Slice[string], message string) {
	body := SMSIRSendRequest{
		LineNumber:  r.Sender,
		MessageText: message,
		Mobiles:     dest,
	}

	// کلاینت Resty
	client := resty.New()

	// ارسال درخواست
	res, err := client.R().
		SetHeader("x-api-key", r.Token).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(r.BaseUrl + "/bulk")
	if err != nil {
		fmt.Print(err.Error())
	}
	if res != nil {
		fmt.Print(res.RawResponse)
	}

}

func (r SMSIRMethods) SendCode(dest string, code string) {
	body := SMSIRSendTemplateRequest{
		Mobile:     dest,
		TemplateId: int64(r.TemplateId),
		Parameters: []SMSIRSendTemplateParameter{
			{
				Name:  r.PositionVariable,
				Value: code,
			},
		},
	}

	// کلاینت Resty
	client := resty.New()

	// ارسال درخواست
	res, err := client.R().
		SetHeader("x-api-key", r.Token).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(r.BaseUrl + "/verify")
	if err != nil {
		fmt.Print(err.Error())
	}
	if res != nil {
		fmt.Print(res.RawResponse)
	}
}
