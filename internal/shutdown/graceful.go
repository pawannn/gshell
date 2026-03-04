package shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Wait(cleanup func()) {

	sig := make(chan os.Signal, 1)

	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-sig

	fmt.Println("\nShutting down gshell...")

	cleanup()

	os.Exit(0)
}
