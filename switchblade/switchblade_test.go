package switchblade

import (
	//"fmt"
	"testing"
)

func TestLinkerTest(t *testing.T) {

	//SET GLOBAL PCR VARIABLES for testing
    
    var PCRDetails map[string]float64

	PCRDetails["RTError"] = 0.0000003

	PCRDetails["DNAPolError"] = 0.000000001 

	PCRDetails["NumPCRCycles"] = 20
	
	// 
	
    //create series of reads and linkers and bools

}

func TestTrimRead(t *testing.T) {

    // create series of (alignment, bool, and trimmed alignments)
    // trimmedRead := TrimRead(in bio.PairWiseAlignment, hasLinker bool)
    // test trimmedRead
    
}

func TestCLI(t *testing.T) {
}

