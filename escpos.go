package escpos

import (
	"io"
)

const (
	// ASCII DLE (DataLinkEscape)
	DLE byte = 0x10

	// ASCII EOT (EndOfTransmission)
	EOT byte = 0x04

	// ASCII GS (Group Separator)
	GS byte = 0x1D
)

type Printer struct {
	dist io.ReadWriteCloser

	// font metrices
	width, height uint8

	// state toggles ESC[char]
	underline  uint8
	emphasize  uint8
	upsidedown uint8
	rotate     uint8

	// state toggles GS[char]
	reverse, smooth uint8
}

// reset toggles
func (e *Printer) reset() {
	e.width = 1
	e.height = 1

	e.underline = 0
	e.emphasize = 0
	e.upsidedown = 0
	e.rotate = 0

	e.reverse = 0
	e.smooth = 0
}

func New (dst io.ReadWriteCloser) (e *Printer) {
// func New (dst io.ReadWriter) (e *Printer) {
    e = &Printer{dist: dst}
    e.reset()
    return
}

// write raw bytes to printer
func (e *Printer) WriteRaw(data []byte) (n int, err error) {
	if len(data) > 0 {
		// e.logger.Printf("Writing %d bytes\n", len(data))
		e.dist.Write(data)
	} else {
		// e.logger.Printf("Wrote NO bytes\n")
	}

	return 0, nil
}

// read raw bytes from printer
func (e *Printer) ReadRaw(data []byte) (n int, err error) {
	return e.dist.Read(data)
}

// write a string to the printer
func (e *Printer) Write(data string) (int, error) {
	return e.WriteRaw([]byte(data))
}

// init/reset printer settings
func (e *Printer) Init() {
	e.reset()
	e.Write("\x1B@")
}

// end output
func (e *Printer) End() {
	e.Write("\xFA")
}

// send cut
func (e *Printer) Cut() {
	e.Write("\x1DVA0")
}

// send cut minus one point (partial cut)
func (e *Printer) CutPartial() {
	e.WriteRaw([]byte{GS, 0x56, 1})
}

// send cash
func (e *Printer) Cash() {
	e.Write("\x1B\x70\x00\x0A\xFF")
}

// send linefeed
func (e *Printer) Linefeed() {
	e.Write("\n")
}
// usage example we might need
/****** over tcp

	socket, err := net.Dial("tcp", path)
	if err != nil {
        fmt.print(err)
		return
	}
	defer socket.Close()

	printer := escpos.New(socket)
    printer.Init()

*/

/******** over USB on linux
	file, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)
	if err != nil {
        fmt.print(err)
        return
	}
	defer f.Close()

	printer := escpos.New(f)
	printer.Init()

*/
