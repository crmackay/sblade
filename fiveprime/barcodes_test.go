package fiveprime

import (
	"bufio"
	"bytes"
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	sw "github.com/crmackay/switchblade"
	"testing"
)

var scanner = bio.FASTQScanner{Scanner: bufio.NewScanner(testFASTQ), File: nil}

var testSuite []checkBarcode

var nextRead, _ = scanner.NextRead()

// func barcodeInSet(b string) bool {
func TestBarcodeInSet(t *testing.T) {
	fmt.Println("testing TestBarcodeInSet")
}

// func find5pLinker(r *sw.Read, l parsed5pLinker) {
func TestFind5pLinker(t *testing.T) {
	fmt.Println("testing TestFind5pLinker")

	for nextRead.Sequence != nil {
		testSuite = append(testSuite, newCheckBarcode(sw.NewRead(nextRead)))
		nextRead, _ = scanner.NextRead()
	}

	for _, elem := range testSuite {
		read := &elem.read
		find5pLinker(read)
		if read.Barcode != elem.result {
			t.Error(
				"\texpected: \t",
				elem.result,
				"\n\t\tbut got: \t",
				read.Barcode,
			)
		}
	}
}

// func findBarcode(b string, q []uint8) string {
func TestFindBarcode(t *testing.T) {
	fmt.Println("testing TestFindBarcode")
}

// func maxProbBarcode(m map[string]float64) string {
func TestMaxProbBarcode(t *testing.T) {
	fmt.Println("testing TestMaxProbBarcode")
}

type checkBarcode struct {
	read   sw.Read
	result string
}

func newCheckBarcode(r sw.Read) checkBarcode {
	return checkBarcode{
		read:   r,
		result: r.Misc[1:],
	}
}

func init() {

}

// "AGC": "sample1",
// "CAC": "sample2",
// "GCT": "sample3",
// "GTG": "sample4",

// fastaq raw data with the third line that containes the extected barcode result:
var testFASTQ = bytes.NewBufferString(`@HWI-ST560:155:C574EACXX:3:1101:2214:1998 1:N:0:
AGCAGGGAGGACGATGCGGAACTGATGTCTAAGTACGCACGGCCGGTACAGTGAAACTGCGAATGGCTCGTGTCAGTCACTTCCAGCGGGCGTATGCCGT
+AGC
CCCFFFFFHHHHHJJJJJJJJJJJJJJJJIIGHEIIJIGIGGIGFFCCDDECCDDDDDDDDDDDDCDCC-9?28A@>4@AAC3>4439B###########
@HWI-ST560:155:C574EACXX:3:1101:2300:1939 1:N:0:
CACAGGGAGGACGATGCGGGAGTGAGACCGTCTTGCTTACTTGTCCGATGAAATGAATGAAATAGAAAGTGGGAAAATAATGTGTCAGTCACTTCCAGCG
+CAC
BBCFFFFFHHHHHJJJJJJJIJHIGIHHIJJJJJIGIIEHIIGHHHHFFFCEDDEDCCCDDDACDCDDD@CDDACBCCCDDCDEDDCDCCCCCDCC:@99
@HWI-ST560:155:C574EACXX:3:1101:2409:1942 1:N:0:
GACAGGGAGGACGATGCGGAAAAGAATGTGAATCATGGTGTTCTTGTGGTTGGCTATGGGACTCTTGATGGCAAAGATTACTGGCTTGTGAAAAAAGGGT
+CAC
BBBFFFFFHHHHHFIJIJJJJJJJJIHIFIJJJJJJJJ=FGIJJJJHIJJJJHHHHFFFFFDEEEEDDDDDDDDDDDDDDDDDDDDD5?CCDCDD#####
@HWI-ST560:155:C574EACXX:3:1101:2433:1960 1:N:0:
ACCAGGGAGGACGATGCGGACAAGTCCCTGAGGAGCCCTTTGAGCCTGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAA
+AGC
BCCFFFFFHHHHHJJJJJJJIJJJGJIJJJIJJJJJJJJJJGJIJIJHHHGFFFFFFEEEEEEEDDDBDDDDDDDCDDDDDDDDDDDDDDDDDDD@B@DB
@HWI-ST560:155:C574EACXX:3:1101:2381:1976 1:N:0:
TGCAGGGAGGACGATGCGGTGATGTTCACAGTGGCTAAGTTCCGCGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAA
+AGC
BCCFFFFFHHHHHIJJJJJJJJJIIHIIJJJFHJJJJJJJJJJJJJJHHFFFFFFEEEEEEEDDDDDDDDDDCDDBBBDDDDDDDDCDDDDDDD9BB>BD
@HWI-ST560:155:C574EACXX:3:1101:2403:1977 1:N:0:
GATAGGGAGGACGATGCGGCTAAGTGGTTGGAACCCGATTGCCTCTCTGGAGCGTGTCAGTCACTTCCAGCGGGTGTCAGTCACTTCCAGCGGTCGTATG
+GCT
@@@FFFFFHHGHHJJJJJJGIEFHFHGDHGIEGGHIIJIICHHIJHEFHGDDDCDD@BCCDDDDDDA@CDDDDD@><ACDDCCCCCC>CC?>B9@B>833
`)
