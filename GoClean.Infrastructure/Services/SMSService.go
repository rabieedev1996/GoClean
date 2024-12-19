package Services

import (
	Services2 "GoClean/GoClean.Application/Contract/Services"
	GoClean_Domain "GoClean/GoClean.Domain"
	ServiceImpl2 "GoClean/GoClean.Infrastructure/ServiceImpl"
	SMSIR2 "GoClean/GoClean.Infrastructure/ServiceImpl/SMSIR"
)

type SMSService struct {
	SMSImpl ServiceImpl2.ISMSImpl
}

func NewSMSService(config GoClean_Domain.Configs_SMS) Services2.ISMSService {
	var sMSService Services2.ISMSService = SMSService{
		SMSImpl: SMSIR2.NewSMSIRMethods(config.SMSIRToken),
	}
	return sMSService
}
func (r SMSService) Send(dest string, text string) {
	r.SMSImpl.Send([]string{
		dest,
	}, text)
}
func (r SMSService) SendCode(dest string, code string) {
	r.SMSImpl.SendCode(dest, code)
}
