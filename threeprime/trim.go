package threeprime

import (
	//"errors"
	"fmt"
	//bio "github.com/crmackay/gobioinfo"
	sw "github.com/crmackay/switchblade"
	//data "github.com/crmackay/switchblade/research"
	"github.com/crmackay/switchblade/config"
)

func next3pAlign(r *sw.Read) (bool, int) {

	num3pAligns := len(r.Aligns3p)
	alignTo := len(r.Sequence)

	var isContam bool
	var pos int

	if num3pAligns != 0 {
		alignTo = r.Aligns3p[num3pAligns-1].QueryStart
		fmt.Println("alignTo", alignTo)
	}

	alignment := r.Sequence[:alignTo-1].Align(config.Linker3p.Sequence)
	// func threePLinkerTest(a *bio.PairWiseAlignment, r *bio.FASTQRead, testNum int) bool {
	isContam = threePLinkerTest(alignment, r.FASTQRead, num3pAligns+1)
	r.Aligns3p = append(r.Aligns3p, sw.Alignment{PairWiseAlignment: alignment, IsContam: isContam})
	pos = alignment.QueryStart

	return isContam, pos

}

func trim3p(r *sw.Read) {

	var trimTo int

	numAligns := len(r.Aligns3p)

	lastAlign := r.Aligns3p[numAligns-1]
	lastAlignPos := lastAlign.QueryStart
	trimLimit := len(config.Linker5p)

	switch {
	case (lastAlign.IsContam == false && numAligns > 1), (numAligns > 1 && lastAlignPos < trimLimit):
		trimTo = r.Aligns3p[numAligns-2].QueryStart
	case (lastAlign.IsContam == false && numAligns == 1), (lastAlign.IsContam == true && lastAlignPos < trimLimit):
		trimTo = len(r.Sequence)

	}

	r.FinalSeq = r.Sequence[:trimTo]
}

func process3p(r *sw.Read) {

	still3pContam := true

	pos3p := len(r.Sequence)
	for still3pContam && pos3p > len(config.Linker5p) {
		still3pContam, pos3p = next3pAlign(r)
		fmt.Println(still3pContam, pos3p)
	}

	trim3p(r)
}

//
// 	var atEnd bool
//
// 	alignmentNum := len(r.ThreePTrims)
// 	var newAlignment bio.PairWiseAlignment
//
// 	if alignmentNum == 0 {
// 		//run the first alignment
// 		newAlignment = r.Read.Sequence.Align(r.ThreePLinker.Sequence)
//
// 	} else {
//
// 		// find where the last alignment started and set that to where this alignment ends
// 		alignTo := r.ThreePTrims[alignmentNum-1].Alignment.QueryStart
// 		if alignTo > 23 {
// 			// run next alignment
// 			newAlignment = r.Read.Sequence[:alignTo].Align(r.ThreePLinker.Sequence)
// 		} else {
// 			atEnd = true
// 			return atEnd
// 		}
//
// 	}
// 	// add the new alignment to the input struct
// 	r.ThreePTrims = append(r.ThreePTrims, sw.ThreePTrim{Alignment: newAlignment})
// 	atEnd = false
// 	return atEnd
// }
//
// // takes a read and tests the last alignment for a contaminant
// func next3pAlignTest(r *sw.Read) bool {
//
// 	numTrims := len(r.ThreePTrims)
// 	// run the bayesian probability test on the new alignment, and record that value
// 	newAlignment := &r.ThreePTrims[numTrims-1].Alignment
//
// 	result := threePLinkerTest(threePQuerySet{alignment: newAlignment, read: r.Read, testNum: numTrims})
//
// 	r.ThreePTrims[len(r.ThreePTrims)-1].IsLinker = result
//
// 	return result
//
// }
//
