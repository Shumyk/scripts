package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

const LengthInfoLine = 19

var (
	termWidth, _, _ = term.GetSize(int(os.Stdin.Fd()))
	header          = color.New(color.Bold, color.BgHiGreen).SprintFunc()
	green           = color.New(color.Bold, color.FgHiGreen).SprintFunc()
	purple          = color.New(color.Bold, color.FgMagenta).SprintFunc()
	red             = color.New(color.Bold, color.FgHiRed).SprintFunc()
)

func Goodbye(s ...any) {
	fmt.Println(purple(s))
	os.Exit(0)
}

func Error(s ...any) {
	_, _ = fmt.Fprintln(os.Stderr, red(s))
	os.Exit(1)
}

go down
func PrintEnvironmentInfo(service, namespace string) {
	wrapHeader(buildHeader("ENVIRONMENT"))
	fmt.Println(buildInfoLine("service", green(service)))
	fmt.Println(buildInfoLine("namespace", green(namespace)))
	hrLine()
}

func PrintImageInfo(i string) (tag, digest string) {
	wrapHeader(buildHeader("CURRENT IMAGE"))
	tag, digest = ParseImageStr(i)
	imageInfo(tag, digest)
	hrLine()
	return
}

func ParseImageStr(i string) (tag, digest string) {
	parts := strings.Split(i, "@")
	tag = strings.Split(parts[0], ":")[1]
	digest = strings.Split(parts[1], ":")[1]
	return
}

func PrintDeployedImageInfo(tag, digest string) {
	wrapHeader(buildHeader("DEPLOYED IMAGE"))
	imageInfo(tag, digest)
	hrLine()
}

func imageInfo(tag, digest string) {
	fmt.Println(buildInfoLine("tag", green(tag)))
	fmt.Println(buildInfoLine("digest", green(digest)))
}

func hrLine() {
	fmt.Printf("%s\n", strings.Repeat("-", termWidth))
}

func wrapHeader(head string) {
	hrLine()
	fmt.Println(header(fillSpaces(head)))
	hrLine()
}

func fillSpaces(s string) string {
	return s + strings.Repeat(" ", termWidth-len(s))
}

// terminal indentation helpers
// ↓↓↓						↓↓↓
func buildLine(msg, suffix string) string {
	prefix := fmt.Sprintf("|   %v", msg)
	freeSpace := LengthInfoLine - len(prefix)
	spaces := strings.Repeat(" ", freeSpace)
	return fmt.Sprintf("%v%v%v", prefix, spaces, suffix)
}

func buildHeader(header string) string {
	return buildLine(header, "|")
}

func buildInfoLine(key, value string) string {
	suffix := fmt.Sprintf(":  %v", value)
	return buildLine(key, suffix)
}

// end terminal indentation helpers
// ↑↑↑				        ↑↑↑
