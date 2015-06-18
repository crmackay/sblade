# switch-blade
a semi-global alignment and bayesian probability based high-throughput sequence adaptor remover, using golang to implement concurrency and parallelism.
```

  Semi-global
  w
  i
  t
  c
  h

   B ayesian
   L inker
   A lignment
an D
   E xcision
```

## Installation:

### find your binaries here:

### compile the `go` code yourself:

This assumes you have a set up GO environment. If not, you can see this [tutorial](test.com).

Get the SwitchBlade code:

`$ go get https://github.com/crmackay/SwitchBlade`

Build the binary for your system and into your $GOPATH/bin:

`$ go build $GOPATH/github.com/crmackay/SwitchBlade/`

the `switchblade` binary should now be in $GOPATH/bin. If your GO environment is set up correctly the `switchblade` command should now be accessible to you via the command line (terminal, command prompt, etc.).

To test this enter:
`$ switchblade -h` which should bring up the switchblade help screen

## How to Use:

`$ switch-blade -l 'GTGTCAGTGATCGAT' -i rawreads.fastq -o output/trimmedreads.fastq -n 20`



## Overview of Approach
