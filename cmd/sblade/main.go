package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"

	bio "github.com/crmackay/gobioinfo"
	"github.com/crmackay/switchblade/workers"
)

func processFile(in io.Reader, out io.Writer, numWorkers int) {

	// TODO set path to input file

	// outFile := inFile + "_trimmed.fastq"
	// TODO set path to output file

	// TODO start file parser

	rawReads := make(chan *bio.FASTQRead, 100000)
	finishedReads := make(chan *bio.FASTQRead, 100000)
	outputData := make(chan []string, 50000)
	doneSignal := make(chan bool)

	var wg sync.WaitGroup

	// start the single IO worker
	go workers.ReadWrite(in, out, rawReads, finishedReads, outputData, doneSignal)
	wg.Add(1)

	// start numCPU - 1 number of workers
	for c := 0; c < numWorkers; c++ {
		go workers.Trim(rawReads, finishedReads, outputData, doneSignal)
		wg.Add(1)
	}

	// wait until the are all done
	numDones := 0
	for numDones < numWorkers {
		<-doneSignal
		numDones++
		wg.Done()
	}

	// close the channels
	close(finishedReads)
	close(outputData)

	// consume the ReadWrite done signal
	<-doneSignal
	wg.Done()

	fmt.Println("this is the main function being done")
}

func main() {

	// TODO parse command line arguments

	// find the number of logical CPUs on the system
	totalCPUS := runtime.NumCPU()

	// set the golang runtime to use all the available processors
	runtime.GOMAXPROCS(totalCPUS)

	CPUWorkers := totalCPUS - 1
	// CPUWorkers = 1
	//rawReads := make(chan bio.FASTQRead, 1000)

	//processedReads := make(chan bio.FASTQRead, 1000)

	if len(os.Args) != 2 {
		panic("please provide a single path to a .fastq or .fastq.gz file")
	}

	inFilePath := os.Args[1]

	// TODO handle gzipped files

	// if filepath.Ext(filePath) == ".gz" {
	// 	reader, err = gzip.NewReader(file)
	// } else {
	// 	reader = file
	// }

	if strings.HasSuffix(inFilePath, ".fastq") || strings.HasSuffix(inFilePath, ".fastq.gz") {
		inReader, err := os.Open(inFilePath)
		if err != nil {
			fmt.Println("problem reading input")
		}

		outWriter, err := os.Create(inFilePath + "_trimmed.fastq")
		if err != nil {
			fmt.Println("problem creating output file")
		}
		processFile(inReader, outWriter, CPUWorkers)
	} else {
		panic("please provide a single path to a .fastq or .fastq.gz file")
	}
}
