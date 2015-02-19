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


// this function uses the `select` method to keep IO to just one goroutine 
// this combined with our CPU accounting allows for fast loading and offloading
// of data to disk, while ensuring that that maximum CPU utilization is happening

/*func IOSwitch (inFile Path, outFile Path, input chan FASTARead, output chan FASTARead) {

    file := FASTQFile(inFile)
    fmt.Println("Starting IO (writer/reader):")
    fileCompleted := false
    nextRead, err := file.Readline()
        //if err
    for fileCompleted == false {
        select {
            //when read channel is not full put next read in
            case reads <- nextRead:
                nextRead, err = file.Readline()
                    if err.(
            // when read channel is full, purge the output channel to disk
            default:
                outputSize := len(output)
                for i := 0; i < outputSize ; i++ {
                    nextProduct := <- output
                    nextProduct.write(outFile)
                }
        }
    }
    
    
    
    
}
*/

func main() {
    
    //totalCPUS := 4
    
    //CPUWorkers := totalCPUS - 1
    
    //TEMP set number of CPUs to 
    
    //rawReads := chan FASTARead
    
    //processedReads := chan FASTARead
    
    // TODO parse command line arguments
    
        // TODO set number of CPUs
        
        // TODO set path to input file
        
        // TODO set path to output file
    
    // TODO start file parser
    
    // TODO start file writer
    //for c := 0; c < numberCPUs; c++ {
        //go work(rawReads, processedReads)
    //}
    
    // TODO start CPU# of workers
    
    // wait until the are all done
    
    bio.Testfastq()
    bio.Testalign()
    
    fmt.Println("this is the main function")
}
