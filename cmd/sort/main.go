package main

import (
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	"github.com/crmackay/switchblade/config"
	"strings"
)

func main() {

	inPath := ""

	outPath := ""

	outPaths := make(map[string]string)

	for bc := range config.Barcodes {
		outPaths[bc] = outPath + bc + ".fastq"
	}

	reader := bio.NewFASTQScanner(inPath)
	defer reader.Close()

	writer := bio.NewFASTQWriter(outPath)
	defer writer.Close()

	for {
		next, err := reader.NextRead()
		if err != nil {
			break
		}

		if len(next.Sequence) > 15 {
			barcode := strings.Split(next.Misc, ":")[0]
			writer.Write(next)
		}

	}

}
