package main

import (
	"bytes"
	"fmt"
	"testing"

	bio "github.com/crmackay/gobioinfo"
)

func TestProcessFile(t *testing.T) {
	fmt.Println("testing the main switchblade functionality")

	// set up the input io.Reader
	inFile := bytes.NewBuffer(inputData)
	inFile.Write(inputData)

	outFile := new(bytes.Buffer)

	processFile(inFile, outFile, 4)

	fmt.Println(outFile)

	results := bytes.NewBuffer(outputData)

	resultMap := make(map[string]bio.FASTQRead)
	resultReader := bio.NewFASTQScanner(results)
	for {
		nextRead, err := resultReader.NextRead()
		if err != nil {
			// if err != "EOF" {
			// 	fmt.Println(err)
			// }
			break
		}
		resultMap[nextRead.ID] = nextRead
	}

	outputMap := make(map[string]bio.FASTQRead)
	outputReader := bio.NewFASTQScanner(outFile)
	for {
		nextRead, err := outputReader.NextRead()
		if err != nil {
			// if err != "EOF" {
			// 	fmt.Println(err)
			// }
			break
		}
		outputMap[nextRead.ID] = nextRead
	}

	for k, v := range outputMap {
		if string(v.Sequence) != string(resultMap[k].Sequence) {
			t.Error(
				"expected: \t",
				string(resultMap[k].Sequence),
				"\n",
				"but got: \t\t",
				string(v.Sequence),
			)
		}
	}
}

// create FASTQ maps from output buffer and result buffer then check

var inputData = []byte(`@HWI-ST560:155:C574EACXX:3:1101:2214:1998 1:N:0:
AGCAGGGAGGACGATGCGGAACTGATGTCTAAGTACGCACGGCCGGTACAGTGAAACTGCGAATGGCTCGTGTCAGTCACTTCCAGCGGGCGTATGCCGT
+
CCCFFFFFHHHHHJJJJJJJJJJJJJJJJIIGHEIIJIGIGGIGFFCCDDECCDDDDDDDDDDDDCDCC-9?28A@>4@AAC3>4439B###########
@HWI-ST560:155:C574EACXX:3:1101:2300:1939 1:N:0:
CACAGGGAGGACGATGCGGGAGTGAGACCGTCTTGCTTACTTGTCCGATGAAATGAATGAAATAGAAAGTGGGAAAATAATGTGTCAGTCACTTCCAGCG
+
BBCFFFFFHHHHHJJJJJJJIJHIGIHHIJJJJJIGIIEHIIGHHHHFFFCEDDEDCCCDDDACDCDDD@CDDACBCCCDDCDEDDCDCCCCCDCC:@99
@HWI-ST560:155:C574EACXX:3:1101:2409:1942 1:N:0:
CACAGGGAGGACGATGCGGAAAAGAATGTGAATCATGGTGTTCTTGTGGTTGGCTATGGGACTCTTGATGGCAAAGATTACTGGCTTGTGAAAAAAGGGT
+
BBBFFFFFHHHHHFIJIJJJJJJJJIHIFIJJJJJJJJ=FGIJJJJHIJJJJHHHHFFFFFDEEEEDDDDDDDDDDDDDDDDDDDDD5?CCDCDD#####
@HWI-ST560:155:C574EACXX:3:1101:2433:1960 1:N:0:
AGCAGGGAGGACGATGCGGACAAGTCCCTGAGGAGCCCTTTGAGCCTGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAA
+
BCCFFFFFHHHHHJJJJJJJIJJJGJIJJJIJJJJJJJJJJGJIJIJHHHGFFFFFFEEEEEEEDDDBDDDDDDDCDDDDDDDDDDDDDDDDDDD@B@DB
@HWI-ST560:155:C574EACXX:3:1101:2381:1976 1:N:0:
AGCAGGGAGGACGATGCGGTGATGTTCACAGTGGCTAAGTTCCGCGGTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAAAAAAAAAAAAA
+
BCCFFFFFHHHHHIJJJJJJJJJIIHIIJJJFHJJJJJJJJJJJJJJHHFFFFFFEEEEEEEDDDDDDDDDDCDDBBBDDDDDDDDCDDDDDDD9BB>BD
@HWI-ST560:155:C574EACXX:3:1101:2403:1977 1:N:0:
GCTAGGGAGGACGATGCGGCTAAGTGGTTGGAACCCGATTGCCTCTCTGGAGCGTGTCAGTCACTTCCAGCGGGTGTCAGTCACTTCCAGCGGTCGTATG
+
@@@FFFFFHHGHHJJJJJJGIEFHFHGDHGIEGGHIIJIICHHIJHEFHGDDDCDD@BCCDDDDDDA@CDDDDD@><ACDDCCCCCC>CC?>B9@B>833
`)

var outputData = []byte(`@HWI-ST560:155:C574EACXX:3:1101:2214:1998 1:N:0:
ATGTCTAAGTACGCACGGCCGGTACAGTGAAACTGCGAATGGCTC
+AGC:AACT
JJJJJIIGHEIIJIGIGGIGFFCCDDECCDDDDDDDDDDDDCDCC
@HWI-ST560:155:C574EACXX:3:1101:2300:1939 1:N:0:
AGACCGTCTTGCTTACTTGTCCGATGAAATGAATGAAATAGAAAGTGGGAAAATAAT
+CAC:GAGT
GIHHIJJJJJIGIIEHIIGHHHHFFFCEDDEDCCCDDDACDCDDD@CDDACBCCCDD
@HWI-ST560:155:C574EACXX:3:1101:2409:1942 1:N:0:
AATGTGAATCATGGTGTTCTTGTGGTTGGCTATGGGACTCTTGATGGCAAAGATTACTGGCTTGTGAAAAAAGG
+CAC:AAAA
JIHIFIJJJJJJJJ=FGIJJJJHIJJJJHHHHFFFFFDEEEEDDDDDDDDDDDDDDDDDDDDD5?CCDCDD###
@HWI-ST560:155:C574EACXX:3:1101:2433:1960 1:N:0:
TCCCTGAGGAGCCCTTTGAGCCTG
+AGC:ACAA
GJIJJJIJJJJJJJJJJGJIJIJH
@HWI-ST560:155:C574EACXX:3:1101:2381:1976 1:N:0:
TTCACAGTGGCTAAGTTCCGCG
+AGC:TGAT
IHIIJJJFHJJJJJJJJJJJJJ
@HWI-ST560:155:C574EACXX:3:1101:2403:1977 1:N:0:
TGGTTGGAACCCGATTGCCTCTCTGGAGC
+GCT:CTAA
FHGDHGIEGGHIIJIICHHIJHEFHGDDD
`)
