package GoClean_Common

type StringTools string

func (receiver StringTools) Substring(characterCount int) string {
	if receiver == "" {
		return ""
	}
	if characterCount >= len(receiver) {
		return string(receiver)
	}
	return string(receiver)[:characterCount]
}
