# escpos #
This is a simple [Golang](http://www.golang.org/project) package that provides
[ESC-POS](https://en.wikipedia.org/wiki/ESC/P) library functions to help with
sending control codes to a ESC-POS capable printer

# Acknowledgments
This project is heavily based on the following GitHub repos:
- [kenshaw-escpos](https://github.com/kenshaw/escpos)
- [alexbrainman-printer](https://github.com/alexbrainman/printer)
- [conejoninja-escpos](https://github.com/conejoninja/go-escpos)

## examples ##

### USB Windows ###
```go
package main
import "github.com/akka-null/escpos"

func main() {

    // Printer Name
	printer, err := escpos.NewWindowsPrinter("R80160")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer printer.Close()

	printer.Init()
	printer.Write("akka\n how are you")
	printer.Linefeed()
	printer.SetReverse(1)
	printer.Write("Hello World!")
	printer.Linefeed()
	printer.Cut()
	printer.End()
}
```

### over Network ###

```go
package main

import( 
    "net"
    "github.com/akka-null/escpos"
    )

func main() {
	socket, err := net.Dial("tcp", "192.168.100.50:9100")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer socket.Close()
	printer := escpos.New(socket)

	printer.Init()
	printer.Write("akka\n how are you")
	printer.Linefeed()
	printer.SetReverse(1)
	printer.Write("Hello World!")
	printer.Linefeed()
	printer.Cut()
	printer.End()
}
```
