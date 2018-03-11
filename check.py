#!/usr/bin/env python2

import re, sys

def an(n): return (re.compile('a?'*n+'a'*n), 'a'*n)

n = int(sys.argv[1])
r, s = an(n)
idxs = r.match(s).span()
print 'n:', n, 'match:', '[%s]' % ' '.join(map(str, idxs))
