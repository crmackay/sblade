//

package main

import (
	"fmt"
	bio "github.com/crmackay/gobioinfo"
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

	rawReads := chan FASTARead

	processedReads := chan FASTARead

	// TODO set path to input file

	// TODO set path to output file

	// TODO start file parser

	// TODO start file writer
	//for c := 0; c < numberCPUs; c++ {
	//go work(rawReads, processedReads)
	//}

	// TODO start CPU# of workers

	// wait until the are all done

	//bio.Testfastq()
	bio.Testalign()
	bio.Testfastqwriter()

	fmt.Println("this is the main function")
}
