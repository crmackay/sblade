package research

import (
	//"bufio"
	//"bytes"
	"fmt"
	//bio "github.com/crmackay/gobioinfo"
	"testing"
)

//func arrayToString(a []int) (s string)
func TestArrayToString(t *testing.T) {

	fmt.Println("testing arrayToString")

	type testPair struct {
		values []int
		result string
	}

	testSuite := []testPair{
		{[]int{20, 10, 14, 0, 2}, "20,10,14,0,2"},
		{[]int{14, 11, 1, 6, 4}, "14,11,1,6,4"},
		{[]int{400, 30, 2, 98, 4}, "400,30,2,98,4"},
		{[]int{20, 23, 13, 65, 90}, "20,23,13,65,90"},
		{[]int{2, 1, 14, 0, 9}, "2,1,14,0,9"},
	}

	for _, elem := range testSuite {
		if arrayToString(elem.values) != elem.result {
			t.Error("arrayToString is not working properly, got ", arrayToString(elem.values), "instead of ", elem.result)
		}
	}

}

// func (p *inProcessRead) getDataCSV() []string {}
func TestGetDataCSV(t *testing.T) {

	//first test align, and prob testing

	fmt.Println("testing getDataCSV")
	/*
		type testPair struct {
			input  *inProcessRead
			output []string
			predicted []
		}

		rawTestData := bytes.NewBufferString("@HWI-ST560:155:C574EACXX:3:1101:2409:1942 1:N:0:\nCACAGGGAGGACGATGCGGAAAAGAATGTGAATCATGGTGTTCTTGTGGTTGGCTATGGGACTCTTGATGGCAAAGATTACTGGCTTGTGAAAAAAGGGT\n+\nBBBFFFFFHHHHHFIJIJJJJJJJJIHIFIJJJJJJJJ=FGIJJJJHIJJJJHHHHFFFFFDEEEEDDDDDDDDDDDDDDDDDDDDD5?CCDCDD#####\n@HWI-ST560:155:C574EACXX:3:1101:2433:1960 1:N:0:\nAGCAGGGAGGACGATGCGGACAAGTCCCTGAGGAGCCCTTTGAGCCTGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAA\n+\nBCCFFFFFHHHHHJJJJJJJIJJJGJIJJJIJJJJJJJJJJGJIJIJHHHGFFFFFFEEEEEEEDDDBDDDDDDDCDDDDDDDDDDDDDDDDDDD@B@DB\n@HWI-ST560:155:C574EACXX:3:1101:2381:1976 1:N:0:\nAGCAGGGAGGACGATGCGGTGATGTTCACAGTGGCTAAGTTCCGCGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAA\n+\nBCCFFFFFHHHHHIJJJJJJJJJIIHIIJJJFHJJJJJJJJJJJJJJHHFFFFFFEEEEEEEDDDDDDDDDDCDDBBBDDDDDDDDCDDDDDDD9BB>BD\n@HWI-ST560:155:C574EACXX:3:1101:2403:1977 1:N:0:\nGCTAGGGAGGACGATGCGGCTAAGTGGTTGGAACCCGATTGCCTCTCTGGAGCGTGTCAGTCACTTCCAGCGGGTGTCAGTCACTTCCAGCGGTCGTATG\n+\n@@@FFFFFHHGHHJJJJJJGIEFHFHGDHGIEGGHIIJIICHHIJHEFHGDDDCDD@BCCDDDDDDA@CDDDDD@><ACDDCCCCCC>CC?>B9@B>833\n@HWI-ST560:155:C574EACXX:3:1101:2425:1982 1:N:0:\nGTGAGGGAGGACGATGCGGTTGTGTGAGAACTGAATTCCATAGGCTGTGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAA\n+\nCCCFFFFFHHFHHIJJJJJIJIGIIJJJFGIIIIIJJJG<FHEHIJIGIHGHJFIHHHHHHFFFFEDDDB@=?CD@ABBBDDDDCCDDC@DDDDDDB@BB")

		scanner := bio.FASTQScanner{Scanner: bufio.NewScanner(rawTestData), File: nil}

		testSuite := make([]testPair, 5)

		for i := 0; i < 5; i++ {
			fastq := scanner.NextRead()
			dna := bio.NewDNASequence("GTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTGCTTG")
			object := newInProcessRead(&(fastq), &dna)
			result := object.getDataCSV()
			testSuite[i] = testPair{object, result}
		}

		for i := range testSuite {
			fmt.Println(i)
		}
	*/
}

func TestAvgPHRED(t *testing.T) {
	fmt.Println("testing avgPHRED")

}

func TestnewReadData(t *testing.T) {
	fmt.Println("testing newReadData")

}
