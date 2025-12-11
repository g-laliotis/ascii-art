//go:build unix || linux || darwin

package ascii

import (
	"os"
	"syscall"
	"unsafe"
)

func getTerminalWidthOS() int {
	type winsize struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}

	ws := &winsize{}
	retVal, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(os.Stdout.Fd()),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retVal) == -1 {
		return -1
	}
	return int(ws.Col)
}
