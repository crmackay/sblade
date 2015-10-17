/*
Bayesian probabiliy test for linker contamination

Here we take a pairwise alignment structure produced by
http://github.com/crmackay/gobioinfo/Align calculate the probabilities of the
query sequence (the read) being present under two assumptions: (1) that it is a
contaminant and (2) thatit is a random sequence. These probabilities are
calculated by factoring insequencing quality data reported by the
Illumina sequencer, as well as user-defined values for Reverse-Transcriptase
error rates, and PCR error rates (errors per base-pair). Using Bayes' theorem,
we are then able to calculate (1) the probability of contamination given the
subject sequence present in the alignment and (2) the probability of a random
sequence given the subject sequence present in the alignment. The greater of
these two probabilities then is taken as being the most likely, and used to
determine whether the sequence in question is a contaminant or not.
*/

package threeprime

import (
	//"fmt"
	bio "github.com/crmackay/gobioinfo"
	//sw "github.com/crmackay/switchblade"
	"github.com/crmackay/switchblade/config"
	"math"
)

/*
threePLinkerTest takes a trimmedRead struct and returns a boolean of whether the
last alignment (x.threePTrims[len(x)-1]) is judged to be a contaminant
after factoring in the sequencing quality score of the query sequence and the
degree of alignment to the known contaminant
*/

func threePLinkerTest(a bio.PairWiseAlignment, r bio.FASTQRead, testNum int) bool {

	// get the segment of the read that is under consideration for this test:
	// this the segment of the read that is part of the alignment

	testStart := a.QueryStart
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
		prob = (probCorrcall * (float64(1) - config.PCRError)) +
			(probMiscall * config.PCRError)

		// fmt.Println("probContamGivenMatch: ", prob)
		return (prob)
	}

	// calculates the probability that a base is a contaminant given that it is
	// an alignment mismatch. It takes the PHRED score of the base as input and
	// returns the probability
	probContamGivenMismatch := func(phred uint8) float64 {

		var probMiscall, probCorrcall, prob float64
		// fmt.Println("pcr error: ", config.PCRError)
		phred64 := float64(phred)
		// fmt.Println("phred: ", phred64)
		probMiscall = math.Pow(10, (-phred64 / 10))
		// fmt.Println("probMiscall: ", probMiscall)
		probCorrcall = 1 - probMiscall
		// fmt.Println("probCorrcall: ", probCorrcall)

		// fmt.Println("probMiscall * (1 - config.PCRError): ", probMiscall*(1-config.PCRError))

		// fmt.Println("(1 / 3) * config.PCRError * probCorrcall: ", (float64(1)/3)*config.PCRError*probCorrcall)

		// fmt.Println("(2 / 3) * config.PCRError * probMiscall: ", (float64(2)/3)*config.PCRError*probMiscall)

		prob = ((float64(1) / 3) * probMiscall * (float64(1) - config.PCRError)) +
			(float64(2)/9)*config.PCRError*probMiscall +
			(float64(1)/3)*config.PCRError*probCorrcall

		// fmt.Println("probContamGivenMismatch: ", prob)
		return (prob)
	}

	// calculates the probability that a base is a contaminant given that is is an alignment InDel
	// the PHRED score of the sequenced base does not matter here, and in fact might not exists
	probContamGivenIndel := func() float64 {

		prob := config.PCRError

		return (prob)
	}

	// calculate P(Sequence|Linker)

	// calculate P(Sequence|Chance)

	// parse CIGAR string and calculate the probability of the sequence given that is a contaminant

	probMiscall := func(phred uint8) float64 {

		phred64 := float64(phred)

		return math.Pow(10, (-phred64 / 10))

	}

	var probSeqGivenContam = 1.0

	var probSeqGivenChance = 1.0

	for i, elem := range a.ExpandedCIGAR {

		// track position along query string, especially to keep track in indels
		queryPosition := testStart + i

		switch {
		case string(elem) == "m":
			probSeqGivenContam *= probContamGivenMatch(r.PHRED.Decoded[queryPosition])
			//		// fmt.Println("probSeqGivenChance", probSeqGivenChance)
			probSeqGivenChance *= 0.25
			//		// fmt.Println("probSeqGivenChance", probSeqGivenChance)
			queryPosition++

		case string(elem) == "x":
			probSeqGivenContam *= probContamGivenMismatch(r.PHRED.Decoded[queryPosition])

			probSeqGivenChance *= 0.75

			queryPosition++

			//		// fmt.Println("at a mismatch:")
			//		// fmt.Printf("probSeqGivenContam: %20.400f", probSeqGivenContam)
			//		// fmt.Println()
			//		// fmt.Printf("probSeqGivenChance: %20.400f", probSeqGivenChance)
			//		// fmt.Println()

		case string(elem) == "n":
			probSeqGivenContam *= probMiscall(r.PHRED.Decoded[queryPosition])
			queryPosition++

		case string(elem) == "i":
			probSeqGivenContam *= probContamGivenIndel()
			// in the case of a calculated deletion in the query seqeuence, we
			// do not increment queryPosition, since we are effectively in a
			// "gap" in the query string

		case string(elem) == "j":
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

	probContam := float64(8) / (math.Pow(10, float64(testNum)) +
		math.Pow(10, float64(testNum-1))*8 +
		math.Pow(8, float64(testNum-1)))

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
