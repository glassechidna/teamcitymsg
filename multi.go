package teamcitymsg

import (
	"fmt"
	"sort"
	"strings"
)

type multi struct {
	msgType string
	attrs   map[string]string
}

func (b *multi) MessageType() string {
	return b.msgType
}

func (b *multi) Attributes() map[string]string {
	return b.attrs
}

func (b *multi) Single() bool {
	return false
}

func (b *multi) String() string {
	var kvps []string

	for k, v := range b.attrs {
		kvps = append(kvps, fmt.Sprintf("%s='%s'", k, EscapeField(v)))
	}

	sort.Strings(kvps)
	joined := strings.Join(kvps, " ")
	return fmt.Sprintf("##teamcity[%s %s]", b.msgType, joined)
}
