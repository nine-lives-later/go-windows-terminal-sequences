// +build linux darwin

package sequences

import (
	"fmt"
)

func EnableVirtualTerminalProcessing() error {
	return fmt.Errorf("windows only package")
}
