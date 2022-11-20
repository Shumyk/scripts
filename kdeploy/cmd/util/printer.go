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

	CurrentImageHeader  = "CURRENT IMAGE"
	DeployedImageHeader = "DEPLOYED IMAGE"
)

func DashLine() {
	fmt.Printf("%s", strings.Repeat("-", termWidth))
}

func PrintEnvironmentInfo(service, namespace string) {
	wrapHeader(buildHeaderLine("ENVIRONMENT"))
	fmt.Println(buildInfoLine("service", green(service)))
	fmt.Println(buildInfoLine("namespace", green(namespace)))
	DashLine()
}

func PrintImageInfo(tag, digest, header string) {
	wrapHeader(buildHeaderLine(header))
	fmt.Println(buildInfoLine("tag", green(tag)))
	fmt.Println(buildInfoLine("digest", green(digest)))
	DashLine()
}

func wrapHeader(title string) {
	DashLine()
	line := withTrailingWhitespaces(title)
	fmt.Println(header(line))
	DashLine()
}

func withTrailingWhitespaces(prefix string) string {
	trailingWhitespaces := strings.Repeat(" ", termWidth-len(prefix))
	return fmt.Sprintf("%v%v", prefix, trailingWhitespaces)
}

// terminal indentation helpers
// ↓↓↓						↓↓↓
func buildLine(msg, suffix string) string {
	prefix := fmt.Sprintf("|   %v", msg)
	freeSpace := LengthInfoLine - len(prefix)
	spaces := strings.Repeat(" ", freeSpace)
	return fmt.Sprintf("%v%v%v", prefix, spaces, suffix)
}

func buildHeaderLine(header string) string {
	return buildLine(header, "|")
}

func buildInfoLine(key, value string) string {
	suffix := fmt.Sprintf(":  %v", value)
	return buildLine(key, suffix)
}

// ↑↑↑				      	    ↑↑↑
// end terminal indentation helpers
