import math
import random

def create_quality():
    quality = ''

    dist_mean = float(30)

    for i in xrange(100):
        if i != 0:

            mean = dist_mean - math.log(i**4)
        else:
            mean = dist_mean
        print 'mean: ', mean
        print random.gauss(mean, 3)


    return quality

create_quality()
