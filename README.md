# sblade
[![Build Status](https://travis-ci.org/crmackay/sblade.svg)](https://travis-ci.org/crmackay/sblade) [![Coverage Status](https://coveralls.io/repos/crmackay/sblade/badge.svg?branch=master&service=github)](https://coveralls.io/github/crmackay/sblade?branch=master)

<u>**s**</u>emi-global <u>**B**</u>ayesian <u>**L**</u>inker <u>**A**</u>lignment an<u>**D**</u> <u>**E**</u>xcision

a semi-global alignment and bayesian probability based high-throughput sequence adaptor remover, using golang to allow for parallel processing and easier auditability/reproducability

## What does it do:

- removes malformed, mutated, concatenated, or otherwise hard to find linkers from the 3' end of your small RNA reads using a very accurate probabilistic approach
- removes 5' linkers, and sorts reads by 5' barcodes, using a similar probabilistic approach, 
- allows you to keep more of your reads for downstream processing, since removal of even the most malformed linkers at the 3' end (where sequence quality is generally low) allows you to retain the rest of the read instead of throwing it away


`sblade` was greatly inspired by [scythe](https://github.com/vsbuffalo/scythe) by Vince Buffalo, but instead of only allowing for mutations in the 3'-linker, `sblade` also allows for insertions and deltions in the 3'-linker. `sblade` also can remove concatenated 3' linkers, which are common in small RNA sequencing. 





## more to come (eg To Dos): 
- installation instructions
- precompiled binaries for linux, OSX, and W
- documentation on how to use
- documentation on the underlying algorithms and methodology
