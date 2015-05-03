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
	var fp *os.File
	var fh int
	var fd uintptr
	var err error

	if runtime.GOOS == "windows" {
		fh, err = syscall.Open("CONOUT$", syscall.O_RDWR, 0)
		fd = uintptr(fh)
	} else {
		fp, err = os.OpenFile("/dev/tty", syscall.O_WRONLY, 0)
		fd = fp.Fd()
	}

	if err != nil {
		return err, int(0), int(0)
	}

	_, _, _ = syscall.Syscall(
		syscall.SYS_IOCTL,
		fd,
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&sz)))

	return nil, int(sz.cols), int(sz.rows)
}
