package switchblade

import (
// "fmt"
//bio "github.com/crmackay/gobioinfo"
)

func alignAndTest3p(r *inProcessRead) {

	alignmentNum := len(r.threePTrims)

	// find where the last alignment started and set that to where this alignment ends
	alignTo := r.threePTrims[alignmentNum-1].alignment.SubjectStart

	// run the next alignment
	newAlignment := r.read.Sequence[:alignTo].Align(r.threePLinker.Sequence)

	// add the new alignment to the input struct
	r.threePTrims[alignmentNum].alignment = newAlignment

	// run the bayesian probability test on the new alignment, and record that value
	r.threePTrims[alignmentNum].isLinker = threePLinkerTest(r)
}

/*
func makeFinishedRead(t *trimmedRead) (finishedRead bio.FASTQRead) {

	trimFrom := t.threePTrims[len(t.threePTrims)-2].alignment.SubjectStart

	finishedRead = trimFASTQ(t.read, trimFrom)

	return
}

func trim3p(t *inProcessRead) bio.FASTQRead {

	var finishedRead bio.FASTQRead

	for t.threePTrims[len(t.threePTrims)-1].isLinker == true {
		alignAndTest3p(t)
	}

	if t.threePTrims[len(t.threePTrims)-1].isLinker == false {
		finishedRead = makeFinishedRead(t)
	}

	return finishedRead
}
*/
