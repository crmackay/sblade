package research

import (
	//"encoding/csv"
	//"fmt"
	sw "github.com/crmackay/switchblade/types"
	"strconv"
	"strings"
)

/* fields that we want to collect per read:

- number of cuts conducted
- final length
- length of each cut
- avg PHRED score before and after

fields that we want to collect per file:
- total number of reads
- total number of linkers removed
- avg final read length
- avg before PHRED score
- avg final PHRED score

*/

type readData struct {
	numCuts        int
	finalLen       int
	cutLens        []int
	avgPHREDBefore float64
	avgPHREDAfter  float64
}

type fileData struct {
	numReads    int
	numLinkers  int
	avgFinalLen float64
}

func newReadData(p *sw.Read) readData {
	var numCuts, finalLen int
	var cutLens []int
	var avgPHREDBefore, avgPHREDAfter float64

	for _, elem := range p.Aligns3p {
		if elem.IsContam == true {
			cutLens = append(cutLens, len(elem.Query)-elem.QueryStart)
		}
	}
	numCuts = len(cutLens)

	if numCuts > 1 {
		finalLen = p.Aligns3p[len(p.Aligns3p)-2].QueryStart
	} else {
		finalLen = len(p.Sequence)
	}

	avgPHREDBefore = avgPHRED(p.PHRED.Decoded)
	avgPHREDBefore = avgPHRED(p.PHRED.Decoded[:finalLen])

	newData := readData{
		numCuts:        numCuts,
		finalLen:       finalLen,
		cutLens:        cutLens,
		avgPHREDBefore: avgPHREDBefore,
		avgPHREDAfter:  avgPHREDAfter,
	}
	return newData
}

func avgPHRED(s []uint8) float64 {
	var i, sum int
	for i, elem := range s {
		sum += int(elem)
		i++
	}
	avg := float64(sum) / float64(i)
	return avg
}

// GetDataCSV does that
func GetDataCSV(p *sw.Read) []string {
	newData := newReadData(p)

	toCSV := []string{
		strconv.Itoa(newData.numCuts),
		strconv.Itoa(newData.finalLen),
		arrayToString(newData.cutLens),
		strconv.FormatFloat(newData.avgPHREDBefore, 'f', 5, 64),
		strconv.FormatFloat(newData.avgPHREDAfter, 'f', 5, 64),
	}

	return toCSV

}

func arrayToString(a []int) (s string) {

	newStrings := make([]string, len(a))

	for i := range a {
		newStrings[i] = strconv.Itoa(int(a[i]))
	}

	s = strings.Join(newStrings, ",")
	return s
}
