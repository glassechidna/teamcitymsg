package teamcitymsg

import "regexp"

var escapeRegex = regexp.MustCompile(`(['\n\r|\[\]])`)
var unescapeRegex = regexp.MustCompile(`\|['nr\[\]|]`)

func EscapeField(input string) string {
	return escapeRegex.ReplaceAllStringFunc(input, func(s string) string {
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

func UnescapeField(input string) string {
	return unescapeRegex.ReplaceAllStringFunc(input, func(s string) string {
		s = s[1:]
		if s == "n" {
			return "\n"
		} else if s == "r" {
			return "\r"
		} else {
			return s
		}
	})
}
