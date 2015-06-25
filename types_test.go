package switchblade

import (
	"fmt"
	"testing"
)

/*
type OrigRead struct {
	*bio.FASTQRead
	aligns3p     []Alignment
	finalSeq     bio.NucleotideSequence
	finalComplex bool
	//	aligns5p []Alignment
	//	barcode	[]rune
}
*/

// func NewOrigRead(f *bio.FASTQRead) (r OrigRead) {

func TestNewOrigRead(t *testing.T) {
	fmt.Println("testing NewOrigRead")

}

// func (r OrigRead) Trim() error {
func TestTrimOrigRead(t *testing.T) {
	fmt.Println("testing OrigRead.Trim()")
}

// func (r OrigRead) CalcComplex()
func TestCalcComplexOrigRead(t *testing.T) {
	fmt.Println("testing OrigRead.CalcComplex()")
}
