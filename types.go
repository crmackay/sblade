package switchblade

import (
	"errors"
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	"github.com/crmackay/switchblade/complexity"
)

// Alignment is a wrapper around gobioinfo.PairWiseAlignment, which adds a bool
// field for indicating whether the aliged region is considered a contaminant
type Alignment struct {
	bio.PairWiseAlignment
	isContam bool
}

// OrigRead is a wrapper around gobioinfo.FASTQRead that add attributes that
// also includes arrays of end alignments, a final sequence after trimming, and
// a bool field indicated whether the final read passed the complexity filter
type OrigRead struct {
	*bio.FASTQRead
	aligns3p     []Alignment
	finalSeq     bio.NucleotideSequence
	finalComplex bool
	//	aligns5p []Alignment
	//	barcode	[]rune
}

// NewOrigRead takes a *gobioninfo.FASTQRead and packages it into a OrigRead
// with the additional fields defined as nil
func NewOrigRead(f *bio.FASTQRead) (r OrigRead) {
	r = OrigRead{FASTQRead: f}
	return r
}

// Trim uses that three-p alignments and five-p alignments found within the given struct to
// determine at which base the 3p contaminant starts, the 5p contaminante ends, and the read starts and ends. It then
// trims the original sequence to those positions, and stored the final read as
// OrigRead.finalSeq. Trim returns an error if the alignment arrays are misformed in someway.
func (r OrigRead) Trim() error {

	//find the last alignment
	// trim 3p
	var err error
	var cutFrom3p int

	numAligns3p := len(r.aligns3p)
	lastAlignContam := r.aligns3p[numAligns3p-1].isContam

	switch {
	case numAligns3p == 0:
		cutFrom3p = len(r.Sequence)
		err = errors.New("no alignments were found, something went wrong with the read processing before attempting to trim it")
	case lastAlignContam == true:
		cutFrom3p = r.aligns3p[numAligns3p-1].QueryStart
	case numAligns3p > 1 && lastAlignContam == false:
		cutFrom3p = r.aligns3p[numAligns3p-2].QueryStart
	}

	for i, elem := range r.aligns3p {
		fmt.Println("3paligns list no: ", i, "is: ", elem.QueryStart)

		r.finalSeq = r.Sequence[:cutFrom3p]

	}
	return err
}

// CalcComplex applies the complexity algorithm to the OrigRead.finalSeq, calculates
// whether the final read is complex or not and stores that result as a bool in
// OrigRead.finalComplex.
func (r OrigRead) CalcComplex() {

	//isNotComplex(s bio.DNASequence)

	r.finalComplex = complexity.IsComplex(r.Sequence)

}
