package cmd

import "strings"

// TODO: refactoring
func ParseImageStr(i string) (tag, digest string) {
	parts := strings.Split(i, "@")
	tag = strings.Split(parts[0], ":")[1]
	digest = strings.Split(parts[1], ":")[1]
	return
}
