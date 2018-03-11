#!/usr/bin/env python2

import re, sys

def an(n): return ('a?'*n+'a'*n, 'a'*n)

n = int(sys.argv[1])
p, s = an(n)
idxs = re.match(p,s).span()
print 'n:', n, 'match:', '[%s]' % ' '.join(map(str, idxs))
