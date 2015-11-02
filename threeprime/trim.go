package threeprime

import (
	//"errors"

	//bio "github.com/crmackay/gobioinfo"

	"github.com/crmackay/switchblade/config"
	sw "github.com/crmackay/switchblade/types"
)

// REFACTORING:
// process the read from the 5p end (found first) to then end

func get3pEnd(r *sw.Read) {
	var isContam bool
	numTests := 1
	lastPos := len(r.Sequence)
	for {
		alignment := r.Sequence.SG3pAlign(config.Linker3p.Sequence)
		numTests++
		// take match and test if for a positive
		isContam = threePLinkerTest(alignment, r.FASTQRead, numTests)

		if isContam == true {
			// if YES, align and test again from before that linker
			// lastPos =
		} else {
			// if NO, save 3p end as the start of last match, or the end of read
			r.End3p = lastPos
			break
		}

		// TODO add break if read is too short or 3p is less than 5p or something
	}

}

func next3pAlign(r *sw.Read) (bool, int) {
	num3pAligns := len(r.Aligns3p)
	// alignTo defaults to the 3-prime end of the full sequence (len -1 )
	alignTo := len(r.Sequence)
	// alignFrom := config.LenFivePrimeLinker

	var isContam bool
	var pos int

	if num3pAligns != 0 {
		alignTo = r.Aligns3p[num3pAligns-1].QueryStart
	}

	// TODO: optimization - do a simply test for if the whole linker string is present, and if so
	//  just cut it off and move one without aligning or doing a probab

	alignment := r.Sequence[:alignTo].SG3pAlign(config.Linker3p.Sequence)

	// func threePLinkerTest(a *bio.PairWiseAlignment, r *bio.FASTQRead, testNum int) bool {}
	isContam = threePLinkerTest(alignment, r.FASTQRead, num3pAligns+1)
	//fmt.Println(isContam)
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

	r.End3p = trimTo
}

// Process3p does 3p aligning, and fills in 3p data to the supplied struct
func Process3p(r *sw.Read) {

	still3pContam := true

	for still3pContam {

		still3pContam, _ = next3pAlign(r)
	}

	trim3p(r)
}
