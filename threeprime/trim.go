package switchblade

import (
	//"errors"
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	sw "github.com/crmackay/switchblade"
	data "github.com/crmackay/switchblade/research"
)

func next3pAlign(r *sw.InProcessRead) bool {

	var atEnd bool

	alignmentNum := len(r.ThreePTrims)
	var newAlignment bio.PairWiseAlignment

	if alignmentNum == 0 {
		//run the first alignment
		newAlignment = r.Read.Sequence.Align(r.ThreePLinker.Sequence)

	} else {

		// find where the last alignment started and set that to where this alignment ends
		alignTo := r.ThreePTrims[alignmentNum-1].Alignment.QueryStart
		if alignTo > 23 {
			// run next alignment
			newAlignment = r.Read.Sequence[:alignTo].Align(r.ThreePLinker.Sequence)
		} else {
			atEnd = true
			return atEnd
		}

	}
	// add the new alignment to the input struct
	r.ThreePTrims = append(r.ThreePTrims, sw.ThreePTrim{Alignment: newAlignment})
	atEnd = false
	return atEnd
}

// takes a read and tests the last alignment for a contaminant
func next3pAlignTest(r *sw.InProcessRead) bool {

	// run the bayesian probability test on the new alignment, and record that value
	newAlignment := &r.ThreePTrims[len(r.ThreePTrims)-1].Alignment

	result := threePLinkerTest(newAlignment, r.Read)

	r.ThreePTrims[len(r.ThreePTrims)-1].IsLinker = result

	return result

}

func process3p(r *sw.InProcessRead) (bio.FASTQRead, []string) {

	hasContam := true
	atEnd := false
	for hasContam && !atEnd {
		atEnd = next3pAlign(r)

		hasContam = next3pAlignTest(r)

	}

	f, err := r.Trim3p()
	if err != nil {
		fmt.Println(err)
	}

	csvData := data.GetDataCSV(r)

	return f, csvData
}
