/*
* original source: https://github.com/conejoninja/go-escpos/blob/main/windows.go
 */
package escpos

import "github.com/akka-null/escpos/windows"

type Windows_Printer struct {
	ptr *windows.WinPrinter
}

func NewWindowsPrinter(path string) (*Printer, error) {
	p, err := windows.Open(path)
	var wPritner Windows_Printer
	if err != nil {
		return &Printer{dist: wPritner}, err
	}
	wPritner.ptr = p
	wPritner.ptr.StartRawDocument("invoice.txt")

	return &Printer{dist: wPritner}, nil

}

func (wp Windows_Printer) Write(data []byte) (n int, err error) {
	return wp.ptr.Write(data)
}

func (wp Windows_Printer) Read(data []byte) (n int, err error) {
	return 0, nil
}

func (wp Windows_Printer) Close() error {
	wp.ptr.EndDocument()
	return wp.ptr.Close()
}
