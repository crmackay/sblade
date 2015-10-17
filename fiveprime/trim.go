package fiveprime

import (
	//"fmt"
	sw "github.com/crmackay/switchblade/types"
	// "github.com/crmackay/switchblade/config"
)

// Process5p takes the supplied read pointer and fills in the barcode sequence, degenerate bases sequence, and the 5pLinker end point
func Process5p(r *sw.Read) {
	find5pLinker(r)
}
