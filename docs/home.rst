============
Switch-Blade
============

A high-throughput sequencing read trimmer designed to properly trim highly amplified sequencing libraries, allowing for the greatest accuracy and preservation of the most number of useable sequences despite PCR and sequencing errors in the linker sequence (generally at the 3' end of a read)

Switch-Blade is unique in a few ways:

- it uses a alignment algorithm to find the best linker match regardless of mutations of indels introduced by sequencer misscalls or PCR errors (similar to Cutadapt)
- it uses a Bayesian Probability model (based upon the approach used by Vince Buffalo with Scythe) to then determine whether the found sequence is a linker or just part of a read
- it is written in ``golang`` to run concurrently and when run on multipe-core systems, can be run in a massively parrellel manner (has been tested on 4-100 core systems...with Switch-Blade and its ``-n`` command-line flag, more cores = more speed)

