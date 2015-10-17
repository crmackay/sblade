package main

import (
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	"github.com/crmackay/switchblade/workers"
	"runtime"
	"sync"
)

/*
func work (input chan FASTARead, ouput chan FASTARead) {
    //processRead.Process
    fmt.Println("Starting Worker:")
}*/

func main() {

	// TODO parse command line arguments

	// find the number of logical CPUs on the system
	totalCPUS := runtime.NumCPU()

	// set the golang runtime to use all the available processors
	runtime.GOMAXPROCS(totalCPUS)

	CPUWorkers := totalCPUS - 1
	CPUWorkers = 1
	//rawReads := make(chan bio.FASTQRead, 1000)

	//processedReads := make(chan bio.FASTQRead, 1000)

	// TODO set path to input file
	inFile := "/Users/christophermackay/Desktop/deepseq_data/pir1/hits-clip/working_data/sample_data/fastq/sblade-test/sample_25000_2.fastq"

	outFile := "/Users/christophermackay/Desktop/deepseq_data/pir1/hits-clip/working_data/sample_data/fastq/sblade-test/output.fastq"
	// TODO set path to output file

	// TODO start file parser

	rawReads := make(chan *bio.FASTQRead, 1000)
	finishedReads := make(chan *bio.FASTQRead, 1000)
	outputData := make(chan []string, 1000)
	doneSignal := make(chan bool)

	var wg sync.WaitGroup

	// start the single IO worker
	go workers.ReadWriter(inFile, outFile, rawReads, finishedReads, outputData, doneSignal)
	wg.Add(1)

	// start CPU-1 number of workers
	for c := 0; c < CPUWorkers; c++ {
		go workers.Trimmer(rawReads, finishedReads, outputData, doneSignal)
		wg.Add(1)
	}

	// wait until the are all done
	numDones := 0
	for numDones < CPUWorkers {
		<-doneSignal
		numDones++
		wg.Done()
		fmt.Println("Number of Dones: ", numDones)
	}
	close(finishedReads)
	close(outputData)
	<-doneSignal
	wg.Done()
	wg.Wait()

	fmt.Println("this is the main function")
}
