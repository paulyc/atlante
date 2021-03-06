package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-spatial/atlante/svg2pdf"
)

var height, width float64

func init() {
	flag.Float64Var(&height, "height", 2500, "height")
	flag.Float64Var(&width, "width", 3000, "width")
}

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		fmt.Printf("incorrect number of args (%d)", flag.NArg())
		os.Exit(1)
	}

	fmt.Println(flag.Args())
	err := svg2pdf.GeneratePDF(flag.Arg(0), flag.Arg(1), height, width)
	if err != nil {
		panic(err)
	}
}
