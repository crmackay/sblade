/*
Bayesian probabiliy test for linker contamination

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
	//sw "github.com/crmackay/switchblade"
	conf "github.com/crmackay/switchblade/config"
	"math"
)

/*
threePLinkerTest takes a trimmedRead struct and returns a boolean of whether the
last alignment (x.threePTrims[len(x)-1]) is judged to be a contaminant
after factoring in the sequencing quality score of the query sequence and the
degree of alignment to the known contaminant
*/

type threePQuerySet struct {
	alignment *bio.PairWiseAlignment
	read      *bio.FASTQRead
	testNum   int
}

func threePLinkerTest(s threePQuerySet) bool {

	// get the segment of the read that is under consideration for this test:
	// this the segment of the read that is part of the alignment

	testStart := s.alignment.QueryStart
	// fmt.Println(testStart)
	var hasLinker bool

	// TODO: this should get cleaned up...perhaps by creating a collection of functions, see: http://jordanorelli.com/post/42369331748/function-types-in-go-golang

	// calculates the probability that a base is a contaminant given that it is
	// an alignment match. It takes the PHRED score of the base of interest as
	// input and returns the probability
	probContamGivenMatch := func(phred uint8) float64 {

		var probMiscall, probCorrcall, prob float64

		phred64 := float64(phred)
		//	// fmt.Println("phred: ", phred64)
		probMiscall = math.Pow(10, (-phred64 / 10))
		//	// fmt.Println("probMiscall: ", probMiscall)
		probCorrcall = 1 - probMiscall
		//	// fmt.Println("probCorrcall: ", probCorrcall)
		prob = (probCorrcall * (1 - conf.PCRError)) +
			((float64(2) / 3) * conf.PCRError * probMiscall) +
			((float64(1) / 3) * conf.PCRError * probCorrcall)
		// fmt.Println("probContamGivenMatch: ", prob)
		return (prob)
	}

	// calculates the probability that a base is a contaminant given that it is
	// an alignment mismatch. It takes the PHRED score of the base as input and
	// returns the probability
	probContamGivenMismatch := func(phred uint8) float64 {

		var probMiscall, probCorrcall, prob float64
		// fmt.Println("pcr error: ", conf.PCRError)
		phred64 := float64(phred)
		// fmt.Println("phred: ", phred64)
		probMiscall = math.Pow(10, (-phred64 / 10))
		// fmt.Println("probMiscall: ", probMiscall)
		probCorrcall = 1 - probMiscall
		// fmt.Println("probCorrcall: ", probCorrcall)

		// fmt.Println("probMiscall * (1 - conf.PCRError): ", probMiscall*(1-conf.PCRError))

		// fmt.Println("(1 / 3) * conf.PCRError * probCorrcall: ", (float64(1)/3)*conf.PCRError*probCorrcall)

		// fmt.Println("(2 / 3) * conf.PCRError * probMiscall: ", (float64(2)/3)*conf.PCRError*probMiscall)

		prob = (probMiscall * (1 - conf.PCRError) * 3) +
			((float64(1) / 3) * conf.PCRError * probCorrcall) +
			((float64(1) / 3) * (1 - conf.PCRError) * probMiscall)

		// fmt.Println("probContamGivenMismatch: ", prob)
		return (prob)
	}

	probMiscall := func(phred uint8) float64 {

		phred64 := float64(phred)

		return math.Pow(10, (-phred64 / 10))

	}

	// calculates the probability that a base is a contaminant given that is is an alignment InDel
	// the PHRED score of the sequenced base does not matter here, and in fact might not exists
	probContamGivenIndel := func() float64 {

		prob := conf.PCRError

		return (prob)
	}

	// calculate P(Sequence|Linker)

	// calculate P(Sequence|Chance)

	// parse CIGAR string and calculate the probability of the sequence given that is a contaminant

	var probSeqGivenContam = 1.0

	var probSeqGivenChance = 1.0

	for i, elem := range s.alignment.ExpandedCIGAR {

		// track position along query string, especially to keep track in indels
		queryPosition := testStart + i
		// fmt.Println(queryPosition)

		switch {
		case elem == "m":
			probSeqGivenContam *= probContamGivenMatch(s.read.PHRED.Decoded[queryPosition])
			//		// fmt.Println("probSeqGivenChance", probSeqGivenChance)
			probSeqGivenChance *= 0.25
			//		// fmt.Println("probSeqGivenChance", probSeqGivenChance)
			queryPosition++

		case elem == "x":
			probSeqGivenContam *= probContamGivenMismatch(s.read.PHRED.Decoded[queryPosition])

			probSeqGivenChance *= 0.75

			queryPosition++

			//		// fmt.Println("at a mismatch:")
			//		// fmt.Printf("probSeqGivenContam: %20.400f", probSeqGivenContam)
			//		// fmt.Println()
			//		// fmt.Printf("probSeqGivenChance: %20.400f", probSeqGivenChance)
			//		// fmt.Println()

		case elem == "n":
			probSeqGivenContam *= probMiscall(s.read.PHRED.Decoded[queryPosition])
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

		// fmt.Println("probSeqGivenContam", probSeqGivenContam)
		// fmt.Println("probSeqGivenChance", probSeqGivenChance)
	}

	// calculate P(Linker|Sequence)

	/*
		probContam is the *a priori* contaminantion frequency (P(contam))
		which equals the number of alignments that should have a linker, divided by
		the number of total alignments.

		by default we assume that the proportion of linker-contaminaned reads is 80%.

		Since every read is aligned until a negative probability test is found, this
		means that 8 out of 10 reads have a linker, but 8 out of 18 alignments have
		a linker. Therefore the default value here is 8/18 aprox = 0.444
	*/

	probContam := float64(8) / (math.Pow(10, float64(s.testNum)) +
		math.Pow(10, float64(s.testNum-1))*8 +
		math.Pow(8, float64(s.testNum-1)))

	// fmt.Println(probContam)
	probContamGivenSeq := (probContam * probSeqGivenContam) /
		((probSeqGivenContam * probContam) + (probSeqGivenChance * (1 - probContam)))

		// calculate P(Chance|Sequence)

	probChanceGivenSeq := ((1 - probContam) * probSeqGivenChance) /
		((probSeqGivenContam * probContam) + (probSeqGivenChance * (1 - probContam)))

	//// fmt.Printf("probChanceGivenSeq: %20.400f", probChanceGivenSeq)
	// fmt.Println("probChanceGivenSeq", probChanceGivenSeq)
	//// fmt.Printf("probContamGivenSeq: %20.400f", probContamGivenSeq)
	// fmt.Println("probContamGivenSeq", probContamGivenSeq)
	//// fmt.Println()
	// test P(L|S) > P(C|S)

	if probContamGivenSeq > probChanceGivenSeq {
		hasLinker = true
	} else {
		hasLinker = false
	}

	return (hasLinker)

}
