package Services

type ISMSService interface {
	Send(dest string, text string)
	SendCode(dest string, code string)
}
