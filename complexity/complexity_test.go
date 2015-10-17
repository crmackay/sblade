package complexity

import (
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	"testing"
)

func TestTestComplexity(t *testing.T) {
	fmt.Println("testing testComplexity")

	type testPair struct {
		seq    bio.NucleotideSequence
		result bool
	}

	tests := []testPair{
		testPair{
			seq:    bio.NucleotideSequence([]rune("ATG")),
			result: true,
		},
		testPair{
			seq:    bio.NucleotideSequence([]rune("AAAAAAAAAAAAAAAAA")),
			result: false,
		},
	}

	for _, elem := range tests {
		if IsComplex(elem.seq) != elem.result {
			t.Error(string(elem.seq), elem.result)
		}
	}

}
