package workers

import (
	"encoding/csv"
	//"fmt"
	bio "github.com/crmackay/gobioinfo"
	"github.com/crmackay/switchblade/fiveprime"
	"github.com/crmackay/switchblade/research"
	"github.com/crmackay/switchblade/threeprime"
	"github.com/crmackay/switchblade/types"
	"os"
	"unicode/utf8"
	// 	//"path/filepath"
)

// Trimmer consumes the raw reads channel, takes each read and enacts a 3'-trim
// action to it, and then puts the trimmed read into the finishedReads channel
// and puts some meta data about the trimming work that was completed in a
//tab-deliminted format into a
func Trimmer(rawReads chan *bio.FASTQRead, finishedReads chan<- *bio.FASTQRead, outputData chan<- []string, doneChan chan bool) {
	//fmt.Println("starting worker")
	for rawRead := range rawReads {

		inProcessRead := types.NewRead(rawRead)

		threeprime.Process3p(inProcessRead)

		fiveprime.Process5p(inProcessRead)

		data := research.GetDataCSV(inProcessRead)

		//fmt.Println(inProcessRead.DNASequence.Sequence)
		//fmt.Println(inProcessRead.End5p)
		//fmt.Println(inProcessRead.End3p)
		var finishedRead *bio.FASTQRead
		if inProcessRead.End5p < inProcessRead.End3p {
			finishedRead = &bio.FASTQRead{
				DNASequence: bio.DNASequence{
					Sequence: inProcessRead.DNASequence.Sequence[inProcessRead.End5p:inProcessRead.End3p],
				},
				PHRED: bio.PHRED{
					Encoded:  inProcessRead.PHRED.Encoded[inProcessRead.End5p:inProcessRead.End3p],
					Decoded:  inProcessRead.PHRED.Decoded[inProcessRead.End5p:inProcessRead.End3p],
					Encoding: inProcessRead.PHRED.Encoding,
				},
				ID:   inProcessRead.ID,
				Misc: inProcessRead.Barcode + ":" + inProcessRead.DegenBases,
			}
		} else {
			//fmt.Println(string(inProcessRead.DNASequence.Sequence))
			finishedRead = &bio.FASTQRead{
				DNASequence: bio.DNASequence{
					Sequence: inProcessRead.DNASequence.Sequence[inProcessRead.End5p:inProcessRead.End5p],
				},
				PHRED: bio.PHRED{
					Encoded:  inProcessRead.PHRED.Encoded[inProcessRead.End5p:inProcessRead.End5p],
					Decoded:  inProcessRead.PHRED.Decoded[inProcessRead.End5p:inProcessRead.End5p],
					Encoding: inProcessRead.PHRED.Encoding,
				},
				ID:   inProcessRead.ID,
				Misc: inProcessRead.Barcode + ":" + inProcessRead.DegenBases,
			}
		}

		finishedReads <- finishedRead

		outputData <- data

	}
	doneChan <- true
	//fmt.Println("Sending Done")
	return
}

// ReadWriter ..
func ReadWriter(inFile string, outFile string, rawReads chan *bio.FASTQRead, finishedReads chan *bio.FASTQRead, outputData chan []string, doneSignal chan bool) {
	reader := bio.NewFASTQScanner(inFile)
	defer reader.Close()

	doneWriter := bio.NewFASTQWriter(outFile)
	defer doneWriter.Close()

	//

	dataPath := "/Users/christophermackay/Desktop/deepseq_data/pir1/hits-clip/working_data/sample_data/fastq/sblade-test/data" // outfile path + processing_data.txt

	csvfile, err := os.Create(dataPath + ".csv")
	if err != nil {
		//fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()

	dataWriter := csv.NewWriter(csvfile)

	dataWriter.Comma, _ = utf8.DecodeRuneInString("\t")

	for {
		newRead, err := reader.NextRead()
		if err != nil {
			break
		}

		select {

		//input is not full, great fill it up!
		case rawReads <- &newRead:

		//once the input is full, drain the two output channels
		default:
			currChanLen := len(finishedReads)

			for i := 0; i < currChanLen; i++ {

				read := <-finishedReads
				doneWriter.Write(*read)

				dataWriter.WriteAll([][]string{<-outputData})

			}
		}
	}

	close(rawReads)

	for read := range finishedReads {
		doneWriter.Write(*read)
		data := <-outputData
		dataWriter.WriteAll([][]string{data})

	}

	return
}
