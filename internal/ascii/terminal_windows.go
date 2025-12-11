//go:build windows

package ascii

import (
	"syscall"
	"unsafe"
)

func getTerminalWidthOS() int {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("GetConsoleScreenBufferInfo")

	type coord struct {
		X, Y int16
	}

	type smallRect struct {
		Left, Top, Right, Bottom int16
	}

	type consoleScreenBufferInfo struct {
		Size              coord
		CursorPosition    coord
		Attributes        uint16
		Window            smallRect
		MaximumWindowSize coord
	}

	var info consoleScreenBufferInfo
	ret, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(unsafe.Pointer(&info)))
	if ret == 0 {
		return -1
	}
	return int(info.Size.X)
}