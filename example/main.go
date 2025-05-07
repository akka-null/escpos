/*
* original source: https://github.com/kenshaw/escpos/blob/master/example/image2pos/main.go
 */
package main

import (
	"flag"
	"image"
	"log"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/akka-null/escpos"
	"github.com/akka-null/escpos/raster"
)

var (
	pritnerName = flag.String("p", "", "Printer name")
	imgPath     = flag.String("i", "image.png", "Input image")
	threshold   = flag.Float64("t", 0.5, "Black/white threshold")
	align       = flag.String("a", "center", "Alignment (left, center, right)")
	doCut       = flag.Bool("c", false, "Cut after print")
	maxWidth    = flag.Int("max-width", 512, "Printer max width in pixels")
)

func main() {
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

	ep, err := escpos.NewWindowsPrinter(*pritnerName)
	if err != nil {
		log.Fatal(err)
	}

	defer ep.Close()
	log.Print(*pritnerName, " open.")

	ep.Init()

	ep.SetAlign(*align)

	rasterConv := &raster.Converter{
		MaxWidth:  *maxWidth,
		Threshold: *threshold,
	}

	rasterConv.Print(img, ep)

	if *doCut {
		ep.Cut()
	}
	ep.End()
}
