//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package machine

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"syscall"
)

func DialNamedPipe(ctx context.Context, path string) (net.Conn, error) {
	return nil, errors.New("not implemented")
}

func GetEnvSetString(env string, val string) string {
	return fmt.Sprintf("export %s='%s'", env, val)
}

// CheckProcessRunning checks non blocking if the pid exited
// returns nil if process is running otherwise an error if not
func CheckProcessRunning(processHint string, pid int, stderrBuf *bytes.Buffer) error {
	var status syscall.WaitStatus
	pid, err := syscall.Wait4(pid, &status, syscall.WNOHANG, nil)
	if err != nil {
		return fmt.Errorf("failed to read %s process status: %w", processHint, err)
	}
	if pid > 0 {
		stderr := ""
		if stderrBuf != nil {
			stderr = fmt.Sprintf(", stderr: %s", stderrBuf)
		}
		// Child exited, process is no longer running
		if status.Exited() {
			return fmt.Errorf("%s exited unexpectedly with exit code %d%s", processHint, status.ExitStatus(), stderr)
		}
		if status.Signaled() {
			return fmt.Errorf("%s was terminated by signal: %s%s", processHint, status.Signal().String(), stderr)
		}
		return fmt.Errorf("%s exited unexpectedly%s", processHint, stderr)
	}
	return nil
}
