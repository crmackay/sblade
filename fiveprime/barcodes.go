package fiveprime

import (
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	sw "github.com/crmackay/switchblade"
	conf "github.com/crmackay/switchblade/config"
	"math"
	//"strings"
)

type linkerFeat struct {
	start    int
	end      int
	sequence string
}

func newLinkerFeat(st, en int, seq string) linkerFeat {
	return linkerFeat{
		start:    st, // start of the feature in the query sequence
		end:      en, // end of the feature in the query sequence
		sequence: seq,
	}
}

type parsed5pLinker struct {
	sequence   string
	barcode    linkerFeat
	degenBases linkerFeat
}

// parse user supplied linker
// func NewParesed5pLinker(s string) parsed5pLinker {
// 	var feats []linkerFeat
// 	linker := parsed5pLinker{}
// 	currFeat := ""
// 	inFeat := false
// 	for i, elem := range s {
// 		letter := string(elem)
// 		switch {
// 		case letter == "X":
// 			inFeat = true
// 			currFeat += letter
// 		case letter == "N":
//
// 		default:
//
// 		}
// 	}
//
// 	return linker
//
// }

type alignment5p struct {
	bio.PairWiseAlignment
	mapStoQ map[int]int
}

//
//
// CGACGATCXXXAGGGAGGACGATGCGGNNNNG[...Read...]GTGTCAGTCACTTCC
//         ⇪⇪⇪                ⇪⇪⇪⇪⇪
//

// look for exact match at the beginning of sequence
// 	- if there put barcode degen bases and read start position in read

// TODO: if 5p is not exact match: align --> create StoQ map --> find positions for barcode and degen bases

func barcodeInSet(b string) bool {
	found := false

	for k := range conf.Barcodes {
		if b == k {
			found = true
		}
	}

	return found
}

func find5pLinker(r *sw.Read) {

	var barcode, degen string
	var end5p int

	// look for linker

	//lenBC := 3
	//lenLink := 16
	//lenDegen := 4

	// checks and sees if the string is 5' end is a perfect match
	if string(r.Sequence)[3:19] == "AGGGAGGACGATGCGG" && string(r.Sequence[23]) == "G" {
		barcode = string(r.Sequence[0:3])
		degen = string(r.Sequence[19:23])
		end5p = 24
	} // else {
	// TODO: put a function here to perform sg alignment and find linker, barcode and degen sequence
	// bio.
	//}

	if barcodeInSet(barcode) {
		r.Barcode = barcode
	} else {
		bcQual := r.PHRED.Decoded[0:3]
		barcode = findBarcode(barcode, bcQual)
	}

	r.Barcode = barcode
	r.DegenBases = degen
	r.End5p = end5p
}

func findBarcode(b string, q []uint8) string {

	probBaseGivenMatch := func(phred uint8) float64 {

		var probMiscall, probCorrcall, prob float64

		phred64 := float64(phred)
		//	// fmt.Println("phred: ", phred64)
		probMiscall = math.Pow(10, (-phred64 / 10))
		//	// fmt.Println("probMiscall: ", probMiscall)
		probCorrcall = 1 - probMiscall
		//	// fmt.Println("probCorrcall: ", probCorrcall)
		prob = (probCorrcall * (float64(1) - conf.PCRError)) +
			(probMiscall * conf.PCRError)

		// fmt.Println("probContamGivenMatch: ", prob)
		return (prob)
	}

	probBaseGivenMismatch := func(phred uint8) float64 {

		var probMiscall, probCorrcall, prob float64

		phred64 := float64(phred)

		probMiscall = math.Pow(10, (-phred64 / 10))

		probCorrcall = 1 - probMiscall

		prob = ((float64(1) / 3) * probMiscall * (float64(1) - conf.PCRError)) +
			(float64(2)/9)*conf.PCRError*probMiscall +
			(float64(1)/3)*conf.PCRError*probCorrcall

		// fmt.Println("probContamGivenMismatch: ", prob)
		return (prob)
	}

	seqProbs := make(map[string]float64)

	for k := range conf.Barcodes {
		fmt.Println(k)
		probSeq := float64(1)
		bcode := []rune(k)
		for i, elem := range b {
			if string(elem) == string(bcode[i]) {

				probSeq *= probBaseGivenMatch(q[i])
			} else {
				probSeq *= probBaseGivenMismatch(q[i])
			}

		}
		fmt.Println(probSeq)
		seqProbs[k] = probSeq
	}

	var denominator float64
	// TODO: THIS SHOULD BE 0.3 AND 0.2 NOT 0.25
	for k, v := range seqProbs {
		denominator += v * conf.BarcodeRatios[k]

	}
	fmt.Println("denominator: ", denominator)
	bcProbs := make(map[string]float64)
	for k, v := range seqProbs {
		// TODO change this to the correct ration (0.2 and 0.3)
		probBarcodeGivenSeq := (v * conf.BarcodeRatios[k]) / denominator
		bcProbs[k] = probBarcodeGivenSeq

	}
	fmt.Println(bcProbs)

	trueBarcode := maxProbBarcode(bcProbs)

	return trueBarcode
}

func maxProbBarcode(m map[string]float64) string {
	var maxV float64
	var maxK string
	i := 0
	for k, v := range m {
		if i == 0 {
			maxV = v
			maxK = k
		} else {
			if v > maxV {
				maxV = v
				maxK = k
			}
		}
		i++
	}
	fmt.Println(maxK, maxV)
	return maxK

}

//
//
//
