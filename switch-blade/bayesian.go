/*Bayesian test for linker contamination

Here we take a pairwise alignment structure produced by gobioinfo.align, and
calculate the probabilities of the query sequence (the read) being present
under two assumptions: (1) that it is a contaminant and (2) that
it is a random sequence. These probabilities are calculated by factoring in
sequencing quality data reported by the Illumina sequencer, as well as
user-defined values for Reverse-Transcriptase error rates, and PCR error rates
(errors per base-pair). Using Bayes' theorem, we are then able to calculate
(1) the probability of contamination given the subject sequence present in the
alignment and (2) the probability of a random sequence given the subject
sequence present in the alignment. The greater of these two probabilities then
is taken as being the most likely, and used to determine whether the sequence in
question is a contaminant or not.

*/

package switchblade

import (
	"fmt"
	"math"
	bio "github.com/crmackay/gobioinfo"
	"testing"
)

/* takes a pairwise alignment struct and returns a boolean of whether the 
aligned sequence is judged to be a contaminant based after factoring in the 
sequencing quality score and the degree of alignment to the known contaminant */
func LinkerTest(alignment bio.PairWiseAlignment, PCRDetails map[string]float64) bool {
	var hasLinker bool

    errorPCR := PCRDetails["RTError"] + 
                (PCRDetails["DNAPolError"] * PCRDetails["NumPCRCycles"])

    // TODO: this should get cleaned up...perhaps by creating a collection of functions, see: http://jordanorelli.com/post/42369331748/function-types-in-go-golang

    //calculates the probability that a base is a contaminant given that it is an alignment match
    //takes the PHRED score of the base as input and returns the probability
    probContamGivenMatch := func (phred float64) float64 {
        
        probMiscall := math.Pow(10, (-phred/10))
        
        probCorrcall := 1 - probMiscall
        
        prob := (errorPCR * probMiscall) + (probCorrcall * (1 - errorPCR)) + ((1/3) * probMiscall * errorPCR)
        
        return(prob)
    }
    
    //calculates the probability that a base is a contaminant given that it is an alignment mismatch
    //takes the PHRED score of the base as input and returns the probability
    probContamGivenMismatch := func (phred float64) float64 {
        
        probMiscall := math.Pow(10, (-phred/10))
        
        probCorrcall := 1 - probMiscall
        
        prob := ((1 - errorPCR) * probMiscall) + (probCorrcall * errorPCR) + ((2/3) * probMiscall * errorPCR)
        
        return(prob)
    }
    
    //calculates the probability that a base is a contaminant given that is is an alignment InDel
    //the PHRED score of the sequenced base does not matter here, and in fact might not exists
    probContamGivenIndel:= func () float64 {
        
        prob := errorPCR
        
        return(prob)
    }
    
    
	//parse CIGAR string


    // calculate P(Sequence|Linker)
    
    // calculate P(Sequence|Chance)
    
    // calculate P(Linker|Sequence)
    
    // calculate P(Chance|Sequence)
    
    // test P(L|S) > P(C|S)
    
    // return true of false

	hasLinker = true

	return (hasLinker)
}

func Test() {

	//SET GLOBAL PCR VARIABLES for testing
    
    var PCRDetails map[string]float64

	PCRDetails["RTError"] = 0.0000003

	PCRDetails["DNAPolError"] = 0.000000001 

	PCRDetails["NumPCRCycles"] = 20

	prob_of_pcr_error := PCRDetails["RTError"] + (PCRDetails["DNAPolError"] * PCRDetails["NumPCRCycles"])

	prob_of_seq_given_adapter := 1

	testCIGAR := []string{"m", "m", "m", "m", "i", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m"}
	alignment := bio.PairWiseAlignment{Subject: "GTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTG",
		Query:                   "GCTAGGGAGGACGATGCGGTGGTGATGCTGCCACATACACTAAGAAGGTCCTGGACGCGTGTAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAA",
		ExpandedCIGAR:           testCIGAR,
		SubjectStart:            0,
		QueryStart:              58,
		SubjectAlignLen:         40,
		QueryAlignLen:           41,
		GappedSubject:           "GTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTG",
		GappedQuery:             "GTGT-AGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTG",
		AlignmentRepresentation: "|||| ||||||||||||||||||||||||||||||||||||"}
		
	query := alignment.Query

	subject := alignment.Subject

	quality := alignment.query.quality
	result := LinkerTest(alignment)
	fmt.Println(result)
}

func TestLinkerTest(t *testing.T) {
}
