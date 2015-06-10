package switchblade

import (
	//"errors"
	"fmt"
	bio "github.com/crmackay/gobioinfo"
)

func next3pAlign(r *inProcessRead) {
	alignmentNum := len(r.threePTrims)

	var newAlignment bio.PairWiseAlignment

	if alignmentNum == 0 {

		//run the first alignment
		newAlignment = r.read.Sequence.Align(r.threePLinker.Sequence)

	} else {

		// find where the last alignment started and set that to where this alignment ends
		alignTo := r.threePTrims[alignmentNum-1].alignment.QueryStart
		// run next alignment
		newAlignment = r.read.Sequence[:alignTo].Align(r.threePLinker.Sequence)

	}
	// add the new alignment to the input struct
	r.threePTrims[alignmentNum].alignment = newAlignment
}

// takes a read and tests the last alignment for a contaminant
func next3pAlignTest(r *inProcessRead) bool {

	// run the bayesian probability test on the new alignment, and record that value
	newAlignment := &r.threePTrims[len(r.threePTrims)-1].alignment

	result := threePLinkerTest(newAlignment, r.read)

	r.threePTrims[len(r.threePTrims)-1].isLinker = result

	return result

}

func process3p(r *inProcessRead) (bio.FASTQRead, []string) {

	hasContam := true

	for hasContam == true {

		next3pAlign(r)

		hasContam = next3pAlignTest(r)

	}

	f, err := r.trim3p()
	if err != nil {
		fmt.Println(err)
	}

	data := r.getDataCSV()

	return f, data
}
