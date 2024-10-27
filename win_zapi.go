package escpos
import (
    "syscall"
    "unsafe"
)

var _ unsafe.Pointer
var (
	modwinspool = syscall.NewLazyDLL("winspool.drv")

	procClosePrinter       = modwinspool.NewProc("ClosePrinter")
	procOpenPrinterW       = modwinspool.NewProc("OpenPrinterW")
	procStartDocPrinterW   = modwinspool.NewProc("StartDocPrinterW")
	procEndDocPrinter      = modwinspool.NewProc("EndDocPrinter")
	procWritePrinter       = modwinspool.NewProc("WritePrinter")
)


func OpenPrinter(name *uint16, h *syscall.Handle, defaults uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procOpenPrinterW.Addr(), 3, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(h)), uintptr(defaults))
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func WritePrinter(h syscall.Handle, buf *byte, bufN uint32, written *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procWritePrinter.Addr(), 4, uintptr(h), uintptr(unsafe.Pointer(buf)), uintptr(bufN), uintptr(unsafe.Pointer(written)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func StartDocPrinter(h syscall.Handle, level uint32, docinfo *DOC_INFO_1) (err error) {
	r1, _, e1 := syscall.Syscall(procStartDocPrinterW.Addr(), 3, uintptr(h), uintptr(level), uintptr(unsafe.Pointer(docinfo)))
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetPrinterDriver(h syscall.Handle, env *uint16, level uint32, di *byte, n uint32, needed *uint32) (err error) {
    r1, _, e1 := syscall.Syscall6(procGetPrinterDriverW.Addr(), 6, uintptr(h), uintptr(unsafe.Pointer(env)), uintptr(level), uintptr(unsafe.Pointer(di)), uintptr(n), uintptr(unsafe.Pointer(needed)))
    if r1 == 0 {
        if e1 != 0 {
            err = error(e1)
        } else {
            err = syscall.EINVAL
        }
    }
    return
}

func EndDocPrinter(h syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procEndDocPrinter.Addr(), 1, uintptr(h), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func ClosePrinter(h syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procClosePrinter.Addr(), 1, uintptr(h), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
