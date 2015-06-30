package workers

import (
	"fmt"
	bio "github.com/crmackay/gobioinfo"
	five "github.com/crmackay/switchblade/fiveprime"
	three "github.com/crmackay/switchblade/threeprime"
)

func Process(r Read) bio.DNASequence {

	r.Trim()

	csvData = data.GetDataCSV(r)

}
