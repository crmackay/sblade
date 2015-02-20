/*Bayesian test for linker contamination

Here we take a pairwise alignment structure produced by alignment.go, and
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

package main

import (
	"fmt"
	bio "github.com/crmackay/gobioinfo"
)

func baysTest(alignment PairWiseAlignment) bool {
	var isLinker bool

	//parse CIGAR string

	//calculate

	isLinker = true

	return (isLinker)
}

func Test() {

	//SET GLOBAL PCR VARIABLES for testing

	rt_error_rate := config_details["RT_ERROR_RATE"]

	dnapol_error_rate = config_details["DNAPol_ERROR_RATE"]

	pcr_cyles = config_details["PCR_CYCLES"]

	fastq = config_details["FASTQ"]

	prob_of_pcr_error = rt_error_rate + (dnapol_error_rate * pcr_cyles)

	prob_of_seq_given_adapter = 1

	query = alignment.query

	subject = alignment.subject

	quality = alignment.query.quality

	testCIGAR := []string{"m", "m", "m", "m", "i", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m", "m"}
	alignment := PairWiseAlignment{Subject: "GTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTG",
		Query:                   "GCTAGGGAGGACGATGCGGTGGTGATGCTGCCACATACACTAAGAAGGTCCTGGACGCGTGTAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTGAA",
		ExpandedCIGAR:           testCIGAR,
		SubjectStart:            0,
		QueryStart:              58,
		SubjectAlignLen:         40,
		QueryAlignLen:           41,
		GappedSubject:           "GTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTG",
		GappedQuery:             "GTGT-AGTCACTTCCAGCGGTCGTATGCCGTCTTCTGCTTG",
		AlignmentRepresentation: "|||| ||||||||||||||||||||||||||||||||||||"}
	result := baysTest(alignment)
	fmt.Println(result)

}
