// +build linux darwin

package sequences

import (
	"fmt"
	"syscall"
)

func EnableVirtualTerminalProcessing(stream syscall.Handle, enable bool) error {
	return fmt.Errorf("windows only package")
}
