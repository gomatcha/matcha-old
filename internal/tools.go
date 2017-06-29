package internal

import (
	"os"
	"runtime/pprof"

	"github.com/overcyn/matchabridge"
)

func init() {
	matchabridge.RegisterFunc("github.com/gomatcha/matcha/internal printStack", printStack)
}

func printStack() {
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
}
