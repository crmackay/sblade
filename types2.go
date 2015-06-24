package switchblade

import (
	//"errors"
	"fmt"
	bio "github.com/crmackay/gobioinfo"
)

type Alignment struct {
	bio.PairWiseAlignment
	isLinker bool
}

type DNARead struct {
	bio.FASTQRead
	aligns3p []Alignment
	//	aligns5p []Alignment
	//	barcode	[]rune
	finalSeq []uint8
	complexScore
}
