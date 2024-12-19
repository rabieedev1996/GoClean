package ServiceImpl

import "GoClean/GoClean.Common"

type ISMSImpl interface {
	Send(dest GoClean_Common.Slice[string], message string)
	SendCode(dest string, code string)
}
