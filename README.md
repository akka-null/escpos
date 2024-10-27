escpos pkg

# printer.go
type Printer struct 

func Open(name string) (*Printer, error) 
    zapi: func OpenPrinter(name *uint16, h *syscall.Handle, defaults uintptr) (err error) 

func (p *Printer) StartRawDocument(name string) error 
    func (p *Printer) DriverInfo() (*DriverInfo, error) 
        zapi: func GetPrinterDriver(h syscall.Handle, env *uint16, level uint32, di *byte, n uint32, needed *uint32) (err error) 
            func (p *Printer) StartDocument(name, datatype string) error 
                zapi: func StartDocPrinter(h syscall.Handle, level uint32, docinfo *DOC_INFO_1) (err error) 

func (p *Printer) Write(b []byte) (int, error) 
    zapi: func WritePrinter(h syscall.Handle, buf *byte, bufN uint32, written *uint32) (err error) 

func (p *Printer) EndDocument() error 
    zapi: func EndDocPrinter(h syscall.Handle) (err error) 

func (p *Printer) Close() error 
    zapi: func ClosePrinter(h syscall.Handle) (err error) 


## zapi.go
func OpenPrinter(name *uint16, h *syscall.Handle, defaults uintptr) (err error) 
func GetPrinterDriver(h syscall.Handle, env *uint16, level uint32, di *byte, n uint32, needed *uint32) (err error) 
func StartDocPrinter(h syscall.Handle, level uint32, docinfo *DOC_INFO_1) (err error) 
func WritePrinter(h syscall.Handle, buf *byte, bufN uint32, written *uint32) (err error) 
func EndDocPrinter(h syscall.Handle) (err error) 
func ClosePrinter(h syscall.Handle) (err error) 
