package teamcitymsg

import "regexp"

func EscapeField(input string) string {
	re := regexp.MustCompile(`(['\n\r|\[\]])`)
	return re.ReplaceAllStringFunc(input, func(s string) string {
		switch s {
		case "'":
			return "|'"
		case "\n":
			return "|n"
		case "\r":
			return "|r"
		case "|":
			return "||"
		case "[":
			return "|["
		case "]":
			return "|]"
		default:
			panic("unexpected")
		}
	})
}
