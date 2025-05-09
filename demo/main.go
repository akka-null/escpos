package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"net"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/akka-null/escpos"
	"github.com/akka-null/escpos/raster"
)

var (
	// lpDev     = flag.String("p", "/dev/usb/lp0", "Printer dev file")
	imgPath   = flag.String("i", "image.png", "Input image")
	threshold = flag.Float64("t", 0.5, "Black/white threshold")
	align     = flag.String("a", "center", "Alignment (left, center, right)")
	doCut     = flag.Bool("c", false, "Cut after print")
	maxWidth  = flag.Int("printer-max-width", 512, "Printer max width in pixels")
)

func main() {
	Print_Network()
	Print_USB()
	// Print_Img_USB()
	// Print_Img_Network()

}

func Print_USB() {

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

func Print_Network() {
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
	printer.Cut()
	printer.End()
}

func Print_Img_USB() {
	flag.Parse()

	imgFile, err := os.Open(*imgPath)
	if err != nil {
		log.Fatal(err)
	}

	img, imgFormat, err := image.Decode(imgFile)
	imgFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Loaded image, format: ", imgFormat)

	// ----------------------------------------------------------------------

	ep, err := escpos.NewWindowsPrinter("R80160")
	if err != nil {
		log.Fatal(err)
	}

	defer ep.Close()

	ep.Init()

	ep.SetAlign(*align)

	rasterConv := &raster.Converter{
		MaxWidth:  *maxWidth,
		Threshold: *threshold,
	}

	rasterConv.Print(img, ep)
	ep.Cut()
	ep.End()
}


func Print_Img_Network() {
	flag.Parse()

	imgFile, err := os.Open(*imgPath)
	if err != nil {
		log.Fatal(err)
	}

	img, imgFormat, err := image.Decode(imgFile)
	imgFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Loaded image, format: ", imgFormat)

	// ----------------------------------------------------------------------

	socket, err := net.Dial("tcp", "192.168.100.50:9100")
	if err != nil {
		log.Fatal(err)
	}

	defer socket.Close()
	ep := escpos.New(socket)

	ep.Init()

	ep.SetAlign(*align)

	rasterConv := &raster.Converter{
		MaxWidth:  *maxWidth,
		Threshold: *threshold,
	}

	rasterConv.Print(img, ep)

	ep.Cut()
	ep.End()
}

/******
// "github.com/conejoninja/go-escpos"
	p, err := escpos.NewUSBPrinterByPath("") // empry string will do a self discovery
	p, err := escpos.NewWindowsPrinterByPath("R80160")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer p.Close()

	p.Init() // start
	p.Smooth(true) // use smootth printing
	p.Size(2, 2) // set font size
	p.PrintLn("HELLO GO")

	p.Size(1, 1)
	p.Font(escpos.FontB) // change font
	p.PrintLn("This is a test of MECT go-escpos")
	p.Font(escpos.FontA)

	p.Align(escpos.AlignRight) // change alignment
	p.PrintLn("An all Go\neasy to use\nEpson POS Printer library")
	p.Align(escpos.AlignLeft)

	p.Size(2, 2)
	p.PrintLn("* No magic numbers")
	p.PrintLn("* ISO8859-15 ŠÙþþØrt")
	p.Underline(true)
	p.PrintLn("* Extended layout")
	p.Underline(false)
	p.PrintLn("* All in Go!")

	p.Align(escpos.AlignCenter)
	p.Barcode("MECT", escpos.BarcodeTypeCODE39) // print barcode
	p.Align(escpos.AlignLeft)

	// p.FeedN(2) // feed 2
	p.Cut() // cut
	p.End() // stop

*/
