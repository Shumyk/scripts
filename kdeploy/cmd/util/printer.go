package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

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

func Error(s ...string) {
	fmt.Fprintln(os.Stderr, red(s))
	os.Exit(1)
}

func PrintEnvInfo(service, namespace string) {
	wrapHeader("|    ENVIRONMENT   |")
	fmt.Printf("|   service        :  %v\t\n", green(service))
	fmt.Printf("|   namespace      :  %v\t\n", green(namespace))
}

func PrintImageInfo(i string) (tag, digest string) {
	wrapHeader("|   CURRENT IMAGE  |")
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
	wrapHeader("    DEPLOYED IMAGE |")
	imageInfo(tag, digest)
	hrLine()
}

func imageInfo(tag, digest string) {
	fmt.Printf("|   tag            :   %v\t\n", green(tag))
	fmt.Printf("|   digest         :   %v\t\n", green(digest))
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
