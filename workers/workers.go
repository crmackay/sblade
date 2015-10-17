package workers

// import (
// 	"encoding/csv"
// 	"fmt"
// 	bio "github.com/crmackay/gobioinfo"
// 	tp "github.com/crmackay/switchblade/threeprime"
// 	"os"
// 	"unicode/utf8"
// 	//"path/filepath"
// )
//
// // Trim3pWorker consumes the raw reads channel, takes each read and enacts a 3'-trim
// // action to it, and then puts the trimmed read into the finishedReads channel
// // and puts some meta data about the trimming work that was completed in a
// //tab-deliminted format into a
// func Trim3pWorker(rawReads <-chan *bio.FASTQRead, finishedReads chan<- *bio.FASTQRead,
// 	outputData chan<- []string, threePLinker *bio.DNASequence) {
//
// 	for rawRead := range rawReads {
//
// 		inProcessRead := newInProcessRead(rawRead, threePLinker)
//
// 		finishedRead, data := threeprime.Process3p(inProcessRead)
//
// 		finishedReads <- &finishedRead
//
// 		outputData <- data
//
// 	}
//
// 	close(finishedReads)
// 	close(outputData)
// }
//
// // IOWorker ..
// func IOWorker(inFile string, outFile string, rawReads chan bio.FASTQRead,
// 	finishedReads chan bio.FASTQRead, outputData chan []string) {
//
// 	reader := bio.NewFASTQScanner(inFile)
// 	defer reader.Close()
//
// 	doneWriter := bio.NewFASTQWriter(outFile)
// 	defer doneWriter.Close()
//
// 	//
//
// 	dataPath := "this is a path" // outfile path + processing_data.txt
//
// 	csvfile, err := os.Create(dataPath + ".csv")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer csvfile.Close()
//
// 	dataWriter := csv.NewWriter(csvfile)
//
// 	dataWriter.Comma, _ = utf8.DecodeRuneInString("\t")
//
// 	// TODO: this loop looks funky...
//
// 	newRead, err := reader.NextRead()
// 	if err != nil {
// 		close(rawReads)
// 		for read := range rawReads {
// 			doneWriter.Write(read)
// 		}
// 	}
//
// 	select {
//
// 	//input is not full, great fill it up!
// 	case rawReads <- newRead:
// 		newRead, err = reader.NextRead()
// 		if err != nil {
// 			break
// 		}
//
// 	//once the input is full, drain the two output channels
// 	default:
// 		currChanLen := len(finishedReads)
//
// 		for i := 0; i < currChanLen; i++ {
//
// 			doneWriter.Write(<-finishedReads)
//
// 			dataWriter.WriteAll([][]string{<-outputData})
//
// 		}
// 	}
// }
