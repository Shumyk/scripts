package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

var (
	width  int
	header func(a ...any) string
	green  func(a ...any) string
	purple func(a ...any) string
	red    func(a ...any) string
)

func InitPrinter() {
	width, _, _ = term.GetSize(int(os.Stdin.Fd()))
	header = color.New(color.BgHiGreen).SprintFunc()
	green = color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()
	purple = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()
	red = color.New(color.FgHiRed).Add(color.Bold).SprintFunc()
}

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
	fmt.Printf("%s\n", strings.Repeat("-", width))
}

func wrapHeader(head string) {
	hrLine()
	fmt.Println(header(fillSpaces(head)))
	hrLine()
}

func fillSpaces(s string) string {
	return s + strings.Repeat(" ", width-len(s))
}
