package threeprime

import (
	//"errors"
	"fmt"
	//bio "github.com/crmackay/gobioinfo"
	sw "github.com/crmackay/switchblade/types"
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

	alignment := r.Sequence[:alignTo].Align(config.Linker3p.Sequence)
	// func threePLinkerTest(a *bio.PairWiseAlignment, r *bio.FASTQRead, testNum int) bool {
	isContam = threePLinkerTest(alignment, r.FASTQRead, num3pAligns+1)
	r.Aligns3p = append(r.Aligns3p, sw.Alignment{PairWiseAlignment: alignment, IsContam: isContam})
	pos = alignment.QueryStart

	return isContam, pos

}

func trim3p(r *sw.Read) {

	var trimTo int

	numAligns := len(r.Aligns3p)

	lastAlign := r.Aligns3p[numAligns]
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

	pos3p := len(r.Sequence)
	for still3pContam && pos3p > len(config.Linker5p) {
		still3pContam, pos3p = next3pAlign(r)
		fmt.Println(still3pContam, pos3p)
	}

	trim3p(r)
}
