package cmd

import (
	"fmt"
	"strings"
	"time"
)

const (
	DIVIDER       = "     "
	TAGS_DELIM    = ","
	DIGEST_PREFIX = "sha256:"
)

func Date(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TrimDigestPrefix(digest string) string {
	return strings.TrimPrefix(digest, DIGEST_PREFIX)
}

func ToString(strs []string) string {
	return strings.Join(strs, TAGS_DELIM)
}

func AppendSemicolon(tag string) string {
	if len(tag) > 0 {
		return fmt.Sprintf(":%v", tag)
	}
	return ""
}
