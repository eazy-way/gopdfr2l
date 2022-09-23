package main

import (
    "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
    "github.com/pdfcpu/pdfcpu/pkg/api"
    "flag"
	"fmt"
	"os"
)

func main() {
	var infile = flag.String("i","input.pdf","input file name")
	var outfile = flag.String("o","output.pdf","output file name")

	flag.Parse()
	ctx, err := api.ReadContextFile(*infile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	xRefTable := ctx.XRefTable
	rootDict, err := xRefTable.Catalog()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	rootDict.Insert("ViewerPreferences",
		pdfcpu.Dict(
			map[string]Object{
				"FitWindow":    Boolean(true),
				"CenterWindow": Boolean(true),
			},
		),
	)
	if err := api.WriteContextFile(ctx, *outfile); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
