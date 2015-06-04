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
	//"fmt"
	bio "github.com/crmackay/gobioinfo"
	"math"
)

type linkerTestSet struct {
	alignment  bio.PairWiseAlignment
	query      bio.FASTQRead
	subject    bio.NucleotideSequence
	pcrDetails map[string]float64
}

/* takes a pairwise alignment struct and returns a boolean of whether the
aligned sequence is judged to be a contaminant based after factoring in the
sequencing quality score and the degree of alignment to the known contaminant */
func LinkerTest(input linkerTestSet) bool {

	var hasLinker bool

	errorPCR := input.pcrDetails["RTError"] +
		(input.pcrDetails["DNAPolError"] * input.pcrDetails["NumPCRCycles"])

	// TODO: this should get cleaned up...perhaps by creating a collection of functions, see: http://jordanorelli.com/post/42369331748/function-types-in-go-golang

	//calculates the probability that a base is a contaminant given that it is an alignment match
	//takes the PHRED score of the base as input and returns the probability
	probContamGivenMatch := func(phred uint8) float64 {

		var probMiscall, probCorrcall, prob float64

		phred64 := float64(phred)

		probMiscall = math.Pow(10, (-phred64 / 10))

		probCorrcall = 1 - probMiscall

		prob = (errorPCR * probMiscall) + (probCorrcall * (1 - errorPCR)) + ((1 / 3) * probMiscall * errorPCR)

		return (prob)
	}

	//calculates the probability that a base is a contaminant given that it is an alignment mismatch
	//takes the PHRED score of the base as input and returns the probability
	probContamGivenMismatch := func(phred uint8) float64 {

		var probMiscall, probCorrcall, prob float64

		phred64 := float64(phred)

		probMiscall = math.Pow(10, (-phred64 / 10))

		probCorrcall = 1 - probMiscall

		prob = ((1 - errorPCR) * probMiscall) + (probCorrcall * errorPCR) + ((2 / 3) * probMiscall * errorPCR)

		return (prob)
	}

	probMiscall := func(phred uint8) float64 {

		phred64 := float64(phred)

		return math.Pow(10, (-phred64 / 10))

	}

	//calculates the probability that a base is a contaminant given that is is an alignment InDel
	//the PHRED score of the sequenced base does not matter here, and in fact might not exists
	probContamGivenIndel := func() float64 {

		prob := errorPCR

		return (prob)
	}

	// calculate P(Sequence|Linker)

	// calculate P(Sequence|Chance)

	//parse CIGAR string and calculate the probability of the sequence given that is a contaminant

	var probSeqGivenContam float64 = 1

	var probSeqGivenChance float64 = 1

	for _, elem := range input.alignment.ExpandedCIGAR {

		// track position along query string, especially to keep track in indels
		queryPosition := 0

		switch {
		case elem == "m":
			probSeqGivenContam *= probContamGivenMatch(input.query.PHRED[queryPosition])

			probSeqGivenChance *= (1 / 4)

			queryPosition += 1
		case elem == "x":
			probSeqGivenContam *= probContamGivenMismatch(input.query.PHRED[queryPosition])

			probSeqGivenChance *= (3 / 4)

			queryPosition += 1
		case elem == "n":
			probSeqGivenContam *= probMiscall(input.query.PHRED[queryPosition])
			queryPosition += 1
		case elem == "i":
			probSeqGivenContam *= probContamGivenIndel()

		case elem == "j":
			probSeqGivenContam *= probContamGivenIndel()
			queryPosition += 1
		}
	}

	// calculate P(Linker|Sequence)

	var probContam float64 = 0.8

	probContamGivenSeq := (probContam * probSeqGivenContam) /
		((probSeqGivenContam * probContam) + (probSeqGivenChance * (1 - probContam)))

		// calculate P(Chance|Sequence)

	probChanceGivenSeq := ((1 - probContam) * probSeqGivenChance) /
		((probSeqGivenContam * probContam) + (probSeqGivenChance * (1 - probContam)))

		// test P(L|S) > P(C|S)

	if probContamGivenSeq > probChanceGivenSeq {
		hasLinker = true
	} else {
		hasLinker = false
	}

	return (hasLinker)

}
