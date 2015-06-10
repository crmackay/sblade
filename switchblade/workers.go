package switchblade

import (
	"encoding/csv"
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	"os"
	"unicode/utf8"
	//"path/filepath"
)

// Trim3pWorker consumes the raw reads channel, takes each read and enacts a 3'-trim
// action to it, and then puts the trimmed read into the finishedReads channel
// and puts some meta data about the trimming work that was completed in a
//tab-deliminted format into a
func Trim3pWorker(rawReads <-chan *bio.FASTQRead, finishedReads chan<- *bio.FASTQRead,
	outputData chan<- []string, threePLinker *bio.DNASequence) {

	for rawRead := range rawReads {

		inProcessRead := newInProcessRead(rawRead, threePLinker)

		// loop while the last trim is labelled as a linker (ie stop once the last
		// trim is not a linker), or is the trims array is empty
		// TODO: turn this into a new function
		for inProcessRead.threePTrims[len(inProcessRead.threePTrims)-1].isLinker == true || inProcessRead.threePTrims == nil {

			// align and test
			//alignAndTest3p(inProcessRead)

		}

		finishedRead, data := process3p(inProcessRead)

		finishedReads <- &finishedRead

		outputData <- data

	}

	close(finishedReads)
	close(outputData)
}

// IOWorker ..
func IOWorker(inFile string, outFile string, rawReads chan bio.FASTQRead,
	finishedReads chan bio.FASTQRead, outputData chan []string) {

	reader := bio.NewFASTQScanner(inFile)
	defer reader.Close()

	doneWriter := bio.NewFASTQWriter(outFile)
	defer doneWriter.Close()

	//

	dataPath := "this is a path" // outfile path + processing_data.txt

	csvfile, err := os.Create(dataPath + ".csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()

	dataWriter := csv.NewWriter(csvfile)

	dataWriter.Comma, _ = utf8.DecodeRuneInString("\t")

	// TODO: how do you detect when the file is done?...

loop:
	for {

		newRead := reader.NextRead()

		//is the sequence attribute is empty, the whole read is likely empty,
		//indicating that the end of the file was found
		if newRead.Sequence == nil {
			close(rawReads)
			for read := range rawReads {
				doneWriter.Write(read)
			}
			break loop
		}
		select {

		//input is not full, great fill it up!
		case rawReads <- newRead:

		//once the input is full, drain the two output channels
		default:
			currenChanLen := len(finishedReads)

			for i := 0; i < currenChanLen; i++ {

				doneWriter.Write(<-finishedReads)

				dataWriter.WriteAll([][]string{<-outputData})

			}
		}
	}
}
