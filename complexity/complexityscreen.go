/*

determining read complexity via the DustMask algorithm (http://kodomo.fbb.msu.ru/FBB/year_10/ppt/DUST.pdf)

*/

package complexity

import (
	bio "github.com/crmackay/gobioinfo"
	conf "github.com/crmackay/switchblade/config"
)

func incTriplet(m map[string]int, t string) {
	m[t]++
}

func isNotComplex(s bio.DNASequence) bool {
	seq := s.Sequence
	triplets := make(map[string]int)
	for i := 0; i < (len(s.Sequence) - 2); i++ {
		trip := string(seq[i : i+3])
		incTriplet(triplets, trip)
	}

	numerator := float64(0)
	for _, v := range triplets {

		numerator += (float64(v) * (float64(v) - 1)) / 2
	}
	score := numerator / (float64(len(s.Sequence)) - 2)

	var notComplex bool

	if score > conf.ComplexityThreshold {
		notComplex = true
	} else {
		notComplex = false
	}

	return notComplex

}
