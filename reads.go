package switchblade

import (
	"errors"
	"fmt"
	bio "github.com/crmackay/gobioinfo"
)

type threePTrim struct {
	alignment bio.PairWiseAlignment
	isLinker  bool
}

type inProcessRead struct {
	read         *bio.FASTQRead   //query
	threePLinker *bio.DNASequence //subject
	threePTrims  []threePTrim
	cutFrom      int
	// TODO: add fivePTrims
}

func newInProcessRead(r *bio.FASTQRead, t *bio.DNASequence) (i *inProcessRead) {
	i = &inProcessRead{
		read:         r,
		threePLinker: t,
		threePTrims:  make([]threePTrim, 0), // TODO: is this okay? <--
	}
	return i
}

func (r *inProcessRead) trim3p() (bio.FASTQRead, error) {

	numTrims := len(r.threePTrims)
	var trimFrom int

	if r.threePTrims[numTrims-1].isLinker == true {
		return bio.FASTQRead{}, errors.New("there was a problem with trimming")
	}

	if numTrims > 1 {
		for i, elem := range r.threePTrims {
			fmt.Println("trim list no:", i, "is:", elem.alignment.GappedQuery, "starts at:", elem.alignment.QueryStart)
		}
		trimFrom = r.threePTrims[numTrims-2].alignment.QueryStart
		if trimFrom < 23 {
			if numTrims > 2 {
				trimFrom = r.threePTrims[numTrims-3].alignment.QueryStart
			} else {
				trimFrom = len(r.read.Sequence)
			}

		}
	} else {
		trimFrom = len(r.read.Sequence)
	}

	p := bio.FASTQRead{
		DNASequence: bio.DNASequence{Sequence: r.read.Sequence[:trimFrom]},
		ID:          r.read.ID,
		Misc:        r.read.Misc,
		PHRED: bio.PHRED{
			Encoded:  r.read.PHRED.Encoded[:trimFrom],
			Decoded:  r.read.PHRED.Decoded[:trimFrom],
			Encoding: r.read.PHRED.Encoding,
		},
	}

	return p, nil
}
