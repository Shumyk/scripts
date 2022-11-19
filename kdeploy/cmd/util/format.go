package cmd

import (
	"fmt"
	"strings"
	"time"
)

const (
	Divider            string = "     "
	Delimiter          string = ","
	DigestPrefix       string = "sha256:"
	FriendlyDateFormat string = "2006-01-02 15:04:05"
)

func Date(t time.Time) string {
	return t.Format(FriendlyDateFormat)
}

func TrimDigestPrefix(digest string) string {
	return strings.TrimPrefix(digest, DigestPrefix)
}

func JoinComma(parts []string) string {
	return strings.Join(parts, Delimiter)
}

func AppendSemicolon(tag string) string {
	if len(tag) > 0 {
		return fmt.Sprintf(":%v", tag)
	}
	return ""
}
