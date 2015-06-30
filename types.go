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
	IsContam bool
}

// Read is a wrapper around gobioinfo.FASTQRead that add attributes that
// also includes arrays of end alignments, a final sequence after trimming, and
// a bool field indicated whether the final read passed the complexity filter
type Read struct {
	bio.FASTQRead
	Aligns3p     []Alignment
	End3p        int
	End5p        int
	Barcode      string
	DegenBases   string
	FinalSeq     bio.NucleotideSequence
	FinalComplex bool
	//	aligns5p []Alignment
	//	barcode	[]rune
}

// NewRead takes a *gobioninfo.FASTQRead and packages it into a Read
// with the additional fields defined as nil
func NewRead(f bio.FASTQRead) (r Read) {
	r = Read{FASTQRead: f}
	return r
}

// Trim uses that three-p alignments and five-p alignments found within the given struct to
// determine at which base the 3p contaminant starts, the 5p contaminante ends, and the read starts and ends. It then
// trims the original sequence to those positions, and stored the final read as
// Read.FinalSeq. Trim returns an error if the alignment arrays are misformed in someway.
func (r Read) Trim() error {

	//find the last alignment
	// trim 3p
	var err error
	var cutFrom3p int

	numAligns3p := len(r.Aligns3p)
	lastAlignContam := r.Aligns3p[numAligns3p-1].IsContam

	switch {
	case numAligns3p == 0:
		cutFrom3p = len(r.Sequence)
		err = errors.New("no alignments were found, something went wrong with the read processing before attempting to trim it")
	case lastAlignContam == true:
		cutFrom3p = r.Aligns3p[numAligns3p-1].QueryStart
	case numAligns3p > 1 && lastAlignContam == false:
		cutFrom3p = r.Aligns3p[numAligns3p-2].QueryStart
	}

	for i, elem := range r.Aligns3p {
		fmt.Println("3paligns list no: ", i, "is: ", elem.QueryStart)

		r.FinalSeq = r.Sequence[:cutFrom3p]

	}
	return err
}

// CalcComplex applies the complexity algorithm to the Read.FinalSeq, calculates
// whether the final read is complex or not and stores that result as a bool in
// Read.FinalComplex.
func (r Read) CalcComplex() {

	//isNotComplex(s bio.DNASequence)

	r.FinalComplex = complexity.IsComplex(r.Sequence)

}
