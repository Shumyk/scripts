package cmd

import (
	"strings"
	"time"
)

func Date(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TrimDigestPrefix(digest string) string {
	return strings.TrimPrefix(digest, "sha256:")
}

func ToString(strs []string) string {
	return strings.Join(strs, ",")
}
