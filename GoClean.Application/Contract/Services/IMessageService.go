package Services

type IMessageService interface {
	GetMessage(code int) string
}
