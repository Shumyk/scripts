package cmd

import "os"

func TerminateOnSigint(result string) {
	if len(result) == 0 {
		Purple("heh, ctrl+C combination was gently pressed. see you")
		os.Exit(0)
	}
}
