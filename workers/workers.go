package workers

import (
	"fmt"
	"io"

	bio "github.com/crmackay/gobioinfo"
	"github.com/crmackay/switchblade/fiveprime"
	"github.com/crmackay/switchblade/threeprime"
	"github.com/crmackay/switchblade/types"
)

// Trimmer consumes the raw reads channel, takes each read and enacts a 3'-trim
// action to it, and then puts the trimmed read into the finishedReads channel
// and puts some meta data about the trimming work that was completed in a
//tab-deliminted format into a
func Trimmer(rawReads chan *bio.FASTQRead, finishedReads chan<- *bio.FASTQRead, outputData chan<- []string, doneChan chan bool) {
	fmt.Println("starting worker")
	for rawRead := range rawReads {

		inProcessRead := types.NewRead(rawRead)

		threeprime.Process3p(inProcessRead)

		fiveprime.Process5p(inProcessRead)

		// data := research.GetDataCSV(inProcessRead)

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

		// outputData <- data

	}
	fmt.Println("Sending Done")
	doneChan <- true
	fmt.Println("closing worker")
	return
}

func partTimeTrimmer(n int, rawReads chan *bio.FASTQRead, finishedReads chan<- *bio.FASTQRead, outputData chan<- []string) {
	fmt.Println("starting part-time worker")
	for i := 0; i < n; i++ {
		// fmt.Println(i)
		rawRead := <-rawReads
		inProcessRead := types.NewRead(rawRead)

		threeprime.Process3p(inProcessRead)

		fiveprime.Process5p(inProcessRead)

		// data := research.GetDataCSV(inProcessRead)

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
			finishedReads <- finishedRead
		}

		// else {
		// 	//fmt.Println(string(inProcessRead.DNASequence.Sequence))
		// 	finishedRead = &bio.FASTQRead{
		// 		DNASequence: bio.DNASequence{
		// 			Sequence: inProcessRead.DNASequence.Sequence[inProcessRead.End5p:inProcessRead.End5p],
		// 		},
		// 		PHRED: bio.PHRED{
		// 			Encoded:  inProcessRead.PHRED.Encoded[inProcessRead.End5p:inProcessRead.End5p],
		// 			Decoded:  inProcessRead.PHRED.Decoded[inProcessRead.End5p:inProcessRead.End5p],
		// 			Encoding: inProcessRead.PHRED.Encoding,
		// 		},
		// 		ID:   inProcessRead.ID,
		// 		Misc: inProcessRead.Barcode + ":" + inProcessRead.DegenBases,
		// 	}
		// }

		// outputData <- data

	}
	fmt.Println("closing part-time worker")
	return
}

// ReadWriter ..
func ReadWriter(in io.Reader, out io.Writer, rawReads chan *bio.FASTQRead, finishedReads chan *bio.FASTQRead, outputData chan []string, doneSignal chan bool) {
	fmt.Println("starting ReadWriter")

	reader := bio.NewFASTQScanner(in)
	defer reader.Close()

	doneWriter := bio.NewFASTQWriter(out)
	defer doneWriter.Close()

	//

	// dataPath := outFile + "_processing_data"
	//
	// csvfile, err := os.Create(dataPath + ".csv")
	// if err != nil {
	// 	//fmt.Println("Error:", err)
	// 	return
	// }
	// defer csvfile.Close()

	// dataWriter := csv.NewWriter(csvfile)

	// dataWriter.Comma, _ = utf8.DecodeRuneInString("\t")
	written := 0

loop:
	for {

		spaceRawReads := cap(rawReads) - len(rawReads)
		fmt.Println("spaceRawReads", spaceRawReads)

		for i := 0; i < 100000; i++ {
			newRead, err := reader.NextRead()
			//fmt.Println("reading", i)
			if err != nil {
				fmt.Println("done reading")
				close(rawReads)
				break loop
			}
			rawReads <- &newRead
		}
		// TODO this number should not be hardcoded, and will lead to a deadlock on highly pararllel systems or small files

		//partTimeTrimmer(20000, rawReads, finishedReads, outputData)

		currChanLen := len(finishedReads)
		fmt.Println("currChanLen", currChanLen)
		for i := 0; i < 100000; i++ {

			read := <-finishedReads
			doneWriter.Write(*read)
			written++
			//fmt.Println("writing num: ", written)
			// dataWriter.WriteAll([][]string{<-outputData})
		}

	}
	for read := range finishedReads {
		// fmt.Println(ct)
		// fmt.Println("here")
		// fmt.Println(string(read.Sequence))
		doneWriter.Write(*read)
		written++
		//fmt.Println("writing num: ", written)
		// data := <-outputData
		// dataWriter.WriteAll([][]string{data})
	}

	doneSignal <- true
	return
}

// TODO: optimization - use a buffer of FASTQ objects and just recycle them via channels so that
// they are not discarded and invoke the GC
