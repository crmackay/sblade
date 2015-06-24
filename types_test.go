package switchblade

import (
	"fmt"
	"testing"
)

//func newInProcessRead(r *bio.FASTQRead, t *bio.DNASequence) (i *inProcessRead) {}

func TestNewInProcessRead(t *testing.T) {
	fmt.Println("testing newInProcessRead()")
}

// func (r *inProcessRead) trim3p() (bio.FASTQRead, error) {}

func TestTrim3p(t *testing.T) {
	fmt.Println("testing *inProcessRead.trim3p()")

	// create some that should raise errors (ie, last trim.isLinker == true)

}
