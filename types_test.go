package switchblade

import (
	"fmt"
	"testing"
)

/*
type Read struct {
	*bio.FASTQRead
	aligns3p     []Alignment
	finalSeq     bio.NucleotideSequence
	finalComplex bool
	//	aligns5p []Alignment
	//	barcode	[]rune
}
*/

// func NewRead(f *bio.FASTQRead) (r Read) {

func TestNewRead(t *testing.T) {
	fmt.Println("testing NewRead")

}

// func (r Read) Trim() error {
func TestTrimRead(t *testing.T) {
	fmt.Println("testing Read.Trim()")
}

// func (r Read) CalcComplex()
func TestCalcComplexRead(t *testing.T) {
	fmt.Println("testing Read.CalcComplex()")
}
