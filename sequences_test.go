package sequences

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

func TestStdoutSequencesOn(T *testing.T) {
	EnableVirtualTerminalProcessing(syscall.Stdout, true)
	defer EnableVirtualTerminalProcessing(syscall.Stdout, false)

	fmt.Fprintf(os.Stdout, "\x1b[34mHello \x1b[35mWorld\x1b[0m!\n")
}

func TestStdoutSequencesOff(T *testing.T) {
	EnableVirtualTerminalProcessing(syscall.Stdout, false)

	fmt.Fprintf(os.Stdout, "\x1b[34mHello \x1b[35mWorld\x1b[0m!\n")
}

func TestStderrSequencesOn(T *testing.T) {
	EnableVirtualTerminalProcessing(syscall.Stderr, true)
	defer EnableVirtualTerminalProcessing(syscall.Stderr, false)

	fmt.Fprintf(os.Stderr, "\x1b[34mHello \x1b[35mWorld\x1b[0m!\n")
}

func TestStderrSequencesOff(T *testing.T) {
	EnableVirtualTerminalProcessing(syscall.Stderr, false)

	fmt.Fprintf(os.Stderr, "\x1b[34mHello \x1b[35mWorld\x1b[0m!\n")
}
