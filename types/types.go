package types

import (
	bio "github.com/crmackay/gobioinfo"
	"github.com/crmackay/switchblade/complexity"
)

// Alignment is a wrapper around gobioinfo.PairWiseAlignment, which adds a bool
// field for indicating whether the aliged region is considered a contaminant
type Alignment struct {
	bio.PairWiseAlignment
	IsContam bool
}

// Read is a wrapper around gobioinfo.FASTQRead that add attributes that
// also includes arrays of end alignments, a final sequence after trimming, and
// a bool field indicated whether the final read passed the complexity filter
type Read struct {
	bio.FASTQRead
	// Aligns3p       []Alignment
	End3p          int
	End5p          int
	Barcode        string
	DegenBases     string
	Done           bool
	FinalSeq       bio.NucleotideSequence
	IsFinalComplex bool
	// Align5p      Alignment
}

// NewRead takes a *gobioninfo.FASTQRead and packages it into a Read
// with the additional fields defined as nil
func NewRead(f *bio.FASTQRead) (r *Read) {
	r = &Read{FASTQRead: *f}
	return r
}

// CalcComplex applies the complexity algorithm to the Read.FinalSeq, calculates
// whether the final read is complex or not and stores that result as a bool in
// Read.FinalComplex.
func (r Read) CalcComplex() {

	//isNotComplex(s bio.DNASequence)

	r.IsFinalComplex = complexity.IsComplex(r.Sequence)

}
