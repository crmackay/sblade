package switchblade

import (
	"errors"
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
		threePTrims:  make([]threePTrim, 5), // TODO: is this okay? <--
	}
	return i
}

func (r *inProcessRead) trim3p() (bio.FASTQRead, error) {

	numTrims := len(r.threePTrims)

	if r.threePTrims[numTrims-1].isLinker == true {
		return bio.FASTQRead{}, errors.New("there was a problem with trimming")
	}

	trimFrom := r.threePTrims[numTrims-2].alignment.QueryStart

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
