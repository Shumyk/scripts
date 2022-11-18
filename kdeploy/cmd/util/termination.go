package cmd

func TerminateOnSigint(result string) {
	if len(result) == 0 {
		Goodbye("heh, ctrl+C combination was gently pressed. see you")
	}
}

func TerminateOnEmpty[T any](args []T, msg ...string) {
	if len(args) == 0 {
		Error(msg...)
	}
}
