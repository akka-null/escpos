package escpos

type WindowsPrinter struct {
	ptr *WinPrinter
}

func NewWindowsPrinter(path string) (*Printer, error) {
    // p, err :=  printer.Open(path)
    p, err :=  Open(path)
    var wp WindowsPrinter
    if err != nil {
        return &Printer{dist: wp}, err
    }
    wp.ptr = p
    wp.ptr.StartRawDocument("ticket.txt")

    return &Printer{dist: wp}, nil

}

func (wprinter WindowsPrinter) Write(p []byte) (n int, err error) {
	return wprinter.ptr.Write(p)
}

func (wprinter WindowsPrinter) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (wprinter WindowsPrinter) Close() error {
    wprinter.ptr.EndDocument()
    return wprinter.ptr.Close()
}

