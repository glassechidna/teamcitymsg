package teamcitymsg

import "fmt"

type single struct {
	msgType string
	value   string
}

func (b *single) MessageType() string {
	return b.msgType
}

func (b *single) Attributes() map[string]string {
	return map[string]string{"": b.value}
}

func (b *single) Single() bool {
	return true
}

func (b *single) String() string {
	return fmt.Sprintf("##teamcity[%s '%s']", b.msgType, EscapeField(b.value))
}
