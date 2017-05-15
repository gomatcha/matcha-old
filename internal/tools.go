package internal

import (
	"github.com/overcyn/mochibridge"
	"os"
	"runtime/pprof"
)

func init() {
	mochibridge.RegisterFunc("github.com/overcyn/mochi/internal printStack", printStack)
}

func printStack() {
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
}
