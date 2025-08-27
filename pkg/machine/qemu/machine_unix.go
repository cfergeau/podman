//go:build dragonfly || freebsd || linux || netbsd || openbsd

package qemu

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/unix"
)

func isProcessAlive(pid int) bool {
	err := unix.Kill(pid, syscall.Signal(0))
	if err == nil || err == unix.EPERM {
		return true
	}
	return false
}

func sigKill(pid int) error {
	return unix.Kill(pid, unix.SIGKILL)
}

func findProcess(pid int) (int, error) {
	if err := unix.Kill(pid, 0); err != nil {
		if err == unix.ESRCH {
			return -1, nil
		}
		return -1, fmt.Errorf("pinging QEMU process: %w", err)
	}
	return pid, nil
}
