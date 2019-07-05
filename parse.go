package teamcitymsg

import (
	"bufio"
	"io"
	"regexp"
)

var serviceMessageRegex = regexp.MustCompile(`##teamcity\[(\S+)(.+?)$`)
var serviceMessageParamsRegex = regexp.MustCompile(`([a-zA-Z][a-zA-Z0-9-]*)\s*=\s*'([^']+)'`)

func Parse(reader io.Reader) []ServiceMessage {
	scanner := bufio.NewScanner(reader)
	var messages []ServiceMessage

	for scanner.Scan() {
		line := scanner.Text()
		matches := serviceMessageRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			msgType := match[1]
			paramMatches := serviceMessageParamsRegex.FindAllStringSubmatch(match[2], -1)
			attrs := map[string]string{}
			for _, pm := range paramMatches {
				name := pm[1]
				val := pm[2]
				attrs[name] = UnescapeField(val)
			}
			messages = append(messages, &multi{msgType: msgType, attrs: attrs})
		}
	}

	return messages
}
