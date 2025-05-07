/*
* original source: https://github.com/alexbrainman/printer/blob/master/printer.go
 */
package windows

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type DOC_INFO_1 struct {
	DocName    *uint16
	OutputFile *uint16
	Datatype   *uint16
}

// DriverInfo stores information about printer driver.
type DriverInfo struct {
	Name        string
	Environment string
	DriverPath  string
	Attributes  uint32
}

type PRINTER_INFO_5 struct {
	PrinterName              *uint16
	PortName                 *uint16
	Attributes               uint32
	DeviceNotSelectedTimeout uint32
	TransmissionRetryTimeout uint32
}

type DRIVER_INFO_8 struct {
	Version                  uint32
	Name                     *uint16
	Environment              *uint16
	DriverPath               *uint16
	DataFile                 *uint16
	ConfigFile               *uint16
	HelpFile                 *uint16
	DependentFiles           *uint16
	MonitorName              *uint16
	DefaultDataType          *uint16
	PreviousNames            *uint16
	DriverDate               syscall.Filetime
	DriverVersion            uint64
	MfgName                  *uint16
	OEMUrl                   *uint16
	HardwareID               *uint16
	Provider                 *uint16
	PrintProcessor           *uint16
	VendorSetup              *uint16
	ColorProfiles            *uint16
	InfPath                  *uint16
	PrinterDriverAttributes  uint32
	CoreDriverDependencies   *uint16
	MinInboxDriverVerDate    syscall.Filetime
	MinInboxDriverVerVersion uint32
}


const (
	PRINTER_DRIVER_XPS = 0x00000002
)

type WinPrinter struct {
	h syscall.Handle
}

func Open(name string) (*WinPrinter, error) {
	var p WinPrinter
	err := OpenPrinter(&(syscall.StringToUTF16(name))[0], &p.h, 0)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *WinPrinter) Write(b []byte) (int, error) {
	var written uint32
	err := WritePrinter(p.h, &b[0], uint32(len(b)), &written)
	if err != nil {
		return 0, err
	}
	return int(written), nil
}

func (p *WinPrinter) StartDocument(name, datatype string) error {
	d := DOC_INFO_1{
		DocName:    &(syscall.StringToUTF16(name))[0],
		OutputFile: nil,
		Datatype:   &(syscall.StringToUTF16(datatype))[0],
	}
	return StartDocPrinter(p.h, 1, &d)
}

// StartRawDocument calls StartDocument and passes either "RAW" or "XPS_PASS"
// as a document type, depending if printer driver is XPS-based or not.
func (p *WinPrinter) StartRawDocument(name string) error {
	di, err := p.DriverInfo()
	if err != nil {
		return err
	}
	// See https://support.microsoft.com/en-us/help/2779300/v4-print-drivers-using-raw-mode-to-send-pcl-postscript-directly-to-the
	// for details.
	datatype := "RAW"
	if di.Attributes&PRINTER_DRIVER_XPS != 0 {
		datatype = "XPS_PASS"
	}
	return p.StartDocument(name, datatype)
}


// DriverInfo returns information about printer p driver.
func (p *WinPrinter) DriverInfo() (*DriverInfo, error) {
	var needed uint32
	b := make([]byte, 1024*10)
	for {
		err := GetPrinterDriver(p.h, nil, 8, &b[0], uint32(len(b)), &needed)
		if err == nil {
			break
		}
		if err != syscall.ERROR_INSUFFICIENT_BUFFER {
			return nil, err
		}
		if needed <= uint32(len(b)) {
			return nil, err
		}
		b = make([]byte, needed)
	}
	di := (*DRIVER_INFO_8)(unsafe.Pointer(&b[0]))
	return &DriverInfo{
		Attributes:  di.PrinterDriverAttributes,
		Name:        windows.UTF16PtrToString(di.Name),
		DriverPath:  windows.UTF16PtrToString(di.DriverPath),
		Environment: windows.UTF16PtrToString(di.Environment),
	}, nil
}


func (p *WinPrinter) EndDocument() error {
	return EndDocPrinter(p.h)
}


func (p *WinPrinter) Close() error {
	return ClosePrinter(p.h)
}
