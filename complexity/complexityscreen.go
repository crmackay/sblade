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

// IsComplex takes a gobioinfo.NucleotideSequence and applies the DustMask
// algorithm and the complexity threshold (set @ configuration) to determine
// if the sequence is complex or noncomplex. The result is a bool representing
// whether the sequence is complex (true) or noncomplex (false).
func IsComplex(s bio.NucleotideSequence) bool {

	length := len(s)

	triplets := make(map[string]int)
	for i := 0; i < (length - 2); i++ {
		t := string(s[i : i+3])
		incTriplet(triplets, t)
	}

	numerator := float64(0)
	for _, v := range triplets {

		numerator += (float64(v) * (float64(v) - 1)) / 2
	}
	score := numerator / (float64(length) - 2)

	var isComplex bool

	if score > conf.ComplexityThreshold {
		isComplex = false
	} else {
		isComplex = true
	}

	return isComplex

}
