package termSize

import (
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

type size struct {
	rows uint16
	cols uint16
}

func TermSize() (error, int, int) {

	var sz size
	var fd uintptr

	if runtime.GOOS == "windows" {
		if fh, err := syscall.Open("CONOUT$", syscall.O_RDWR, 0); err != nil {
			return err, int(0), int(0)
		} else {
			fd = uintptr(fh)
		}
	} else {
		if fp, err := os.OpenFile("/dev/tty", syscall.O_WRONLY, 0); err != nil {
			return err, int(0), int(0)
		} else {
			fd = fp.Fd()
		}
	}

	_, _, _ = syscall.Syscall(
		syscall.SYS_IOCTL,
		fd,
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&sz)))

	return nil, int(sz.cols), int(sz.rows)
}
