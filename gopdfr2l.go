package main

import (
    "github.com/hhrutter/pdfcpu/pkg/pdfcpu"
    "github.com/hhrutter/pdfcpu/pkg/api"
    "flag"
	"fmt"
	"os"
)

func main() {
	var infile = flag.String("i","input.pdf","input file name")
	var outfile = flag.String("o","output.pdf","output file name")

	flag.Parse()
	ctx, err := ReadContextFile(infile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	rootDict, err := ctx.xRefTable.Catalog()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	rootDict.Insert("ViewerPreferences",
		Dict(
			map[string]Object{
				"FitWindow":    Boolean(true),
				"CenterWindow": Boolean(true),
			}
		)
	)
	if err := WriteContextFile(ctx, outfile); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
)
}