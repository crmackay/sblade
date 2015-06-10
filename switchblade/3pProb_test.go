package switchblade

import (
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	//"math"
	"bufio"
	"bytes"
	"testing"
)

// func threePLinkerTest(alignment *bio.PairWiseAlignment, read *bio.FASTQRead) bool {}
func TestThreePLinkerTest(t *testing.T) {
	fmt.Println("testing threePLinkerTest()")

	numRead := 5
	rawTestData := bytes.NewBufferString("@HWI-ST560:155:C574EACXX:3:1101:2409:1942 1:N:0:\nCACAGGGAGGACGATGCGGAAAAGAATGTGAATCATGGTGTTCTTGTGGTTGGCTATGGGACTCTTGATGGCAAAGATTACTGGCTTGTGAAAAAAGGGT\n+\nBBBFFFFFHHHHHFIJIJJJJJJJJIHIFIJJJJJJJJ=FGIJJJJHIJJJJHHHHFFFFFDEEEEDDDDDDDDDDDDDDDDDDDDD5?CCDCDD#####\n@HWI-ST560:155:C574EACXX:3:1101:2433:1960 1:N:0:\nAGCAGGGAGGACGATGCGGACAAGTCCCTGAGGAGCCCTTTGAGCCTGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAA\n+\nBCCFFFFFHHHHHJJJJJJJIJJJGJIJJJIJJJJJJJJJJGJIJIJHHHGFFFFFFEEEEEEEDDDBDDDDDDDCDDDDDDDDDDDDDDDDDDD@B@DB\n@HWI-ST560:155:C574EACXX:3:1101:2381:1976 1:N:0:\nAGCAGGGAGGACGATGCGGTGATGTTCACAGTGGCTAAGTTCCGCGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAA\n+\nBCCFFFFFHHHHHIJJJJJJJJJIIHIIJJJFHJJJJJJJJJJJJJJHHFFFFFFEEEEEEEDDDDDDDDDDCDDBBBDDDDDDDDCDDDDDDD9BB>BD\n@HWI-ST560:155:C574EACXX:3:1101:2403:1977 1:N:0:\nGCTAGGGAGGACGATGCGGCTAAGTGGTTGGAACCCGATTGCCTCTCTGGAGCGTGTCAGTCACTTCCAGCGGGTGTCAGTCACTTCCAGCGGTCGTATG\n+\n@@@FFFFFHHGHHJJJJJJGIEFHFHGDHGIEGGHIIJIICHHIJHEFHGDDDCDD@BCCDDDDDDA@CDDDDD@><ACDDCCCCCC>CC?>B9@B>833\n@HWI-ST560:155:C574EACXX:3:1101:2425:1982 1:N:0:\nGTGAGGGAGGACGATGCGGTTGTGTGAGAACTGAATTCCATAGGCTGTGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAA\n+\nCCCFFFFFHHFHHIJJJJJIJIGIIJJJFGIIIIIJJJG<FHEHIJIGIHGHJFIHHHHHHFFFFEDDDB@=?CD@ABBBDDDDCCDDC@DDDDDDB@BB")

	testReads := make([]bio.FASTQRead, numRead)
	scanner := bio.FASTQScanner{Scanner: bufio.NewScanner(rawTestData), File: nil}

	for i := 0; i < numRead; i++ {
		testReads[i] = scanner.NextRead()
	}

	subject := bio.NewDNASequence("GTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTGCTTG")

	type testGroup struct {
		alignment      bio.PairWiseAlignment
		read           bio.FASTQRead
		expectedResult bool
	}

	expectedResults := []bool{false, true, true, true, true}

	testSet := make([]testGroup, numRead)

	for i, read := range testReads {
		testSet[i].read = read

		testSet[i].alignment = read.Sequence.Align(subject.Sequence)
		testSet[i].expectedResult = expectedResults[i]
	}

	for _, elem := range testSet {
		result := threePLinkerTest(&elem.alignment, &elem.read)
		if result != elem.expectedResult {
			t.Error(
				"got:",
				result,
				", expected:",
				elem.expectedResult,
				", here is the alignment: ", "\n",
				elem.alignment.GappedQuery, "\n",
				elem.alignment.AlignmentRepresentation, "\n",
				elem.alignment.GappedSubject, "\n",
			)
		}
	}

}
