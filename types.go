package switchblade

import (
	//"errors"
	"fmt"
	bio "github.com/crmackay/gobioinfo"
)

type ThreePTrim struct {
	Alignment bio.PairWiseAlignment
	IsLinker  bool
}

type InProcessRead struct {
	Read         *bio.FASTQRead   //query
	ThreePLinker *bio.DNASequence //subject
	ThreePTrims  []ThreePTrim
	CutFrom      int
	// TODO: add fivePTrims
}

func NewInProcessRead(r *bio.FASTQRead, t *bio.DNASequence) (i *InProcessRead) {
	i = &InProcessRead{
		Read:         r,
		ThreePLinker: t,
		ThreePTrims:  make([]ThreePTrim, 0), // TODO: is this okay? <--
	}
	return i
}

func (r *InProcessRead) Trim3p() (bio.FASTQRead, error) {

	numTrims := len(r.ThreePTrims)
	var trimFrom int

	if r.ThreePTrims[numTrims-1].IsLinker == true {
		fmt.Println("trim list is only 1:", r.ThreePTrims[numTrims-1].Alignment.GappedQuery, "starts at:", r.ThreePTrims[numTrims-1].Alignment.QueryStart)
		trimFrom = r.ThreePTrims[numTrims-1].Alignment.QueryStart
		if trimFrom < 23 {
			trimFrom = len(r.Read.Sequence)
		}
	}

	if numTrims > 1 {
		for i, elem := range r.ThreePTrims {
			fmt.Println("trim list no:", i, "is:", elem.Alignment.GappedQuery, "starts at:", elem.Alignment.QueryStart)
		}
		trimFrom = r.ThreePTrims[numTrims-2].Alignment.QueryStart
		if trimFrom > 23 {
			if numTrims > 2 {
				trimFrom = r.ThreePTrims[numTrims-3].Alignment.QueryStart
			}
		} else {
			trimFrom = len(r.Read.Sequence)

		}
	}
	fmt.Println(trimFrom)
	p := bio.FASTQRead{
		DNASequence: bio.DNASequence{Sequence: r.Read.Sequence[:trimFrom]},
		ID:          r.Read.ID,
		Misc:        r.Read.Misc,
		PHRED: bio.PHRED{
			Encoded:  r.Read.PHRED.Encoded[:trimFrom],
			Decoded:  r.Read.PHRED.Decoded[:trimFrom],
			Encoding: r.Read.PHRED.Encoding,
		},
	}

	return p, nil
}
