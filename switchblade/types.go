package switchblade

import (
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
	// TODO: add fivePTrims
}

func newInProcessRead(r *bio.FASTQRead, t *bio.DNASequence) (i *inProcessRead) {
	i = &inProcessRead{
		read:         r,
		threePLinker: t,
		threePTrims:  nil,
	}
	return i
}

func (r *inProcessRead) trim3p() (p *bio.FASTQRead) {
	trimFrom := r.threePTrims[len(r.threePTrims)-2].alignment.QueryStart

	p = &bio.FASTQRead{
		DNASequence: bio.DNASequence{Sequence: r.read.Sequence[:trimFrom]},
		ID:          r.read.ID,
		Misc:        r.read.Misc,
		PHRED: bio.PHRED{
			Encoded:  r.read.PHRED.Encoded[:trimFrom],
			Decoded:  r.read.PHRED.Decoded[:trimFrom],
			Encoding: r.read.PHRED.Encoding,
		},
	}

	return p
}

var pcrDetails = map[string]float64{
	"RTError":      0.0000003,
	"DNAPolError":  0.000000001,
	"NumPCRCycles": 20,
}

var pcrError = pcrDetails["RTError"] +
	(pcrDetails["DNAPolError"] * pcrDetails["NumPCRCycles"])
