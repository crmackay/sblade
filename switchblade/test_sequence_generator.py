#create test data:

##

import random
import math

class FASTQ(object):
    def __init__(self, sequence, quality, name = None, misc = None)
        self.sequence = sequence
        self.quality = quality
        self.name = name
        self.misc = misc
        
    def toPrint(self):
        return '\n'.join(['@'+self.name, self.sequence, '+'+self.misc, self.quality])+'\n'
    

def create_sequence():
    sequence = ''

    linker_position = random.randrange(20, 90)

    for i in xrange(linker_position):
        sequence += random.choice(['A', 'T','C','G'])

    ins = 0
    dels = 0
    position_in_linker = 0
    for i in xrange(100-linker_position):
        position_in_linker = position_in_linker - dels + ins
        
        if position_in_linker < len(linker):
            if random.random() > 0.05:
                sequence += linker[position_in_linker]
            else:
                to_add = random.choice(["ins","del","mut"])
                if to_add == "ins":
                    ins+=1
                    sequence += random.choice(['A', 'T','C','G']) 
                elif to_add == "del":
                    dels += 1
                    # add nothing to the sequence
                elif to_add == "mut":
                    while True:
                        to_add = random.choice(['A', 'T','C','G'])
                        if to_add != linker[position_in_linker]:
                            sequence += to_add
                            break
        else: 
            position_in_linker = -1
            ins = 0
            dels = 0
        position_in_linker += 1
                            
    return sequence

def create_quality():
    quality = ''
    
    dist_mean = float(30)
    
    for i in xrange(100):
        if i != 0:
            dist_mean = dist_mean - log(i)
        print 
        random.gaussian(30, 2)
    
    return quality

linker = "GTGTCAGTCACTTCCAGCGGTCGTATGCCGTCTTGCTTG"

#random 61 nucleotides

#sequence with random indels and insertions

linker = "


for i in xrange(1000):
    fastq = FASTQ(sequence = create_sequence(), quality = create_quality())
    
    with open(out_path, 'a') as out:
        out.write(fastq.toPrint())


