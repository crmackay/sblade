//

package main

import (
	"fmt"
	sblade "github.com/crmackay/SwitchBlade/switchblade"
	bio "github.com/crmackay/gobioinfo"
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

	rawReads := make(chan bio.FASTQRead, 1000)

	processedReads := make(chan bio.FASTQRead, 1000)

	// TODO set path to input file

	// TODO set path to output file

	// TODO start file parser

	var wg sync.WaitGroup

	// start the single IO worker
	go sblade.IOWork(inFile, outFile, rawReads, processedReads)
	wg.Add(1)

	// start CPU-1 number of workers
	for c := 0; c < CPUWorkers; c++ {
		go sblade.TrimWork(rawReads, processedReads)
		wg.Add(1)
	}

	// wait until the are all done
	wg.Wait()

	fmt.Println("this is the main function")
}
