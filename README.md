# sblade
a semi-global alignment and bayesian probability based high-throughput sequence adaptor remover, using golang to allow for parallel processing (whether on your laptop or a computer cluster) and easier auditability/reproducability

<u>**s**</u>emi-global <u>**B**</u>ayesian <u>**L**</u>inker <u>**A**</u>lignment an<u>**D**</u> <u>**E**</u>xcision

## What does it do:

- removes malformed, mutated, concatenated, or otherwise hard to find linkers from the 3' end of your small RNA reads using a very accurate probabilistic approach
- removes 5' linkers, and sorts reads by 5' barcodes, using a similar probabilistic approach, 
- allows you to keep more of your reads for downstream processes, since removal of even the most malformed linkers at the 3' end (where sequence quality is generally low) allows you to retain the rest of the read instead of just arbitrarily throwing it away


sbald was greatly inspired by [scythe](https://github.com/vsbuffalo/scythe) by Vince Buffalo, but instead of only allowing for mutations in the 3'-linker, sblade also allows for insertions and deltions in the 3'-linker. sblade also can remove concatenated 3' linkers, which are common in small RNA sequencing. 





more to come including: 
- installation instructions
- precompiled binaries for linux, OSX, and W
- documentation on how to use
- documentation on the underlying algorithms and methodology
