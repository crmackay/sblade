/*

Bayesian test for linker contamination

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
	//"fmt"
	bio "github.com/crmackay/gobioinfo"
	"math"
)

/*
threePLinkerTest takes a trimmedRead struct and returns a boolean of whether the
last alignment (x.threePTrims[len(x)-1]) is judged to be a contaminant
after factoring in the sequencing quality score of the query sequence and the
degree of alignment to the known contaminant
*/
func threePLinkerTest(alignment *bio.PairWiseAlignment, read *bio.FASTQRead) bool {

	// get the segment of the read that is under consideration for this test:
	// this the segment of the read that is part of the alignment

	testStart := alignment.QueryStart

	var hasLinker bool

	// TODO: this should get cleaned up...perhaps by creating a collection of functions, see: http://jordanorelli.com/post/42369331748/function-types-in-go-golang

	// calculates the probability that a base is a contaminant given that it is
	// an alignment match. It takes the PHRED score of the base of interest as
	// input and returns the probability
	probContamGivenMatch := func(phred uint8) float64 {

		var probMiscall, probCorrcall, prob float64

		phred64 := float64(phred)

		probMiscall = math.Pow(10, (-phred64 / 10))

		probCorrcall = 1 - probMiscall

		prob = (pcrError * probMiscall) + (probCorrcall * (1 - pcrError)) + ((1 / 3) * probMiscall * pcrError)

		return (prob)
	}

	// calculates the probability that a base is a contaminant given that it is
	// an alignment mismatch. It takes the PHRED score of the base as input and
	// returns the probability
	probContamGivenMismatch := func(phred uint8) float64 {

		var probMiscall, probCorrcall, prob float64

		phred64 := float64(phred)

		probMiscall = math.Pow(10, (-phred64 / 10))

		probCorrcall = 1 - probMiscall

		prob = ((1 - pcrError) * probMiscall) + (probCorrcall * pcrError) + ((2 / 3) * probMiscall * pcrError)

		return (prob)
	}

	probMiscall := func(phred uint8) float64 {

		phred64 := float64(phred)

		return math.Pow(10, (-phred64 / 10))

	}

	//calculates the probability that a base is a contaminant given that is is an alignment InDel
	//the PHRED score of the sequenced base does not matter here, and in fact might not exists
	probContamGivenIndel := func() float64 {

		prob := pcrError

		return (prob)
	}

	// calculate P(Sequence|Linker)

	// calculate P(Sequence|Chance)

	//parse CIGAR string and calculate the probability of the sequence given that is a contaminant

	var probSeqGivenContam float64 = 1

	var probSeqGivenChance float64 = 1

	for _, elem := range alignment.ExpandedCIGAR {

		// track position along query string, especially to keep track in indels
		queryPosition := testStart

		switch {
		case elem == "m":
			probSeqGivenContam *= probContamGivenMatch(read.PHRED.Decoded[queryPosition])

			probSeqGivenChance *= (1 / 4)

			queryPosition++

		case elem == "x":
			probSeqGivenContam *= probContamGivenMismatch(read.PHRED.Decoded[queryPosition])

			probSeqGivenChance *= (3 / 4)

			queryPosition++

		case elem == "n":
			probSeqGivenContam *= probMiscall(read.PHRED.Decoded[queryPosition])
			queryPosition++

		case elem == "i":
			probSeqGivenContam *= probContamGivenIndel()
			// in the case of a calculated deletion in the query seqeuence, we
			// do not increment queryPosition, since we are effectively in a
			// "gap" in the query string

		case elem == "j":
			probSeqGivenContam *= probContamGivenIndel()
			queryPosition++

		}
	}

	// calculate P(Linker|Sequence)

	var probContam float64 = 0.8

	probContamGivenSeq := (probContam * probSeqGivenContam) /
		((probSeqGivenContam * probContam) + (probSeqGivenChance * (1 - probContam)))

		// calculate P(Chance|Sequence)

	probChanceGivenSeq := ((1 - probContam) * probSeqGivenChance) /
		((probSeqGivenContam * probContam) + (probSeqGivenChance * (1 - probContam)))

		//fmt.Printf("probChanceGivenSeq: %20.400f", probChanceGivenSeq)

		// test P(L|S) > P(C|S)

	if probContamGivenSeq > probChanceGivenSeq {
		hasLinker = true
	} else {
		hasLinker = false
	}

	return (hasLinker)

}
