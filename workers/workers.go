package workers

import (
	"fmt"
	"io"

	bio "github.com/crmackay/gobioinfo"
	"github.com/crmackay/switchblade/fiveprime"
	"github.com/crmackay/switchblade/threeprime"
	"github.com/crmackay/switchblade/types"
)

// Trim consumes the raw reads channel, takes each read and enacts a 3'-trim
// action to it, and then puts the trimmed read into the finishedReads channel
// and puts some meta data about the trimming work that was completed in a
// tab-deliminted format into a
func Trim(rawReads chan *bio.FASTQRead, finishedReads chan<- *bio.FASTQRead, outputData chan<- []string, doneChan chan bool) {
	fmt.Println("starting worker")
	for rawRead := range rawReads {

		inProcessRead := types.NewRead(rawRead)

		threeprime.Process3p(inProcessRead)

		fiveprime.Process5p(inProcessRead)

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

	}
	fmt.Println("Sending Done")
	doneChan <- true
	fmt.Println("closing worker")
	return
}

// ReadWrite is a . ..
func ReadWrite(in io.Reader, out io.Writer, rawReads chan *bio.FASTQRead, finishedReads chan *bio.FASTQRead, outputData chan []string, doneSignal chan bool) {
	fmt.Println("starting ReadWriter")

	reader := bio.NewFASTQScanner(in)
	defer reader.Close()

	doneWriter := bio.NewFASTQWriter(out)
	defer doneWriter.Close()

loop:
	for {

		spaceRawReads := cap(rawReads) - len(rawReads)
		fmt.Println("space on raw reads", spaceRawReads)

		for i := 0; i < 100000; i++ {
			// fmt.Println(i)
			newRead, err := reader.NextRead()
			// fmt.Println("reading", i)
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
		// if currChanLen == 0 {
		// 	fmt.Println("waiting...")
		// 	time.Sleep(5 * time.Second)
		// }
		fmt.Println("len of finished reads", currChanLen)
		for i := 0; i < currChanLen; i++ {
			// fmt.Println(i)
			read := <-finishedReads
			doneWriter.Write(*read)
		}

	}
	for read := range finishedReads {
		doneWriter.Write(*read)
	}

	doneSignal <- true
	return
}

// TODO: optimization - use a buffer of FASTQ objects and just recycle them via channels so that
// they are not discarded and invoke the GC
