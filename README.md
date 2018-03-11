regex_bomb
----------

This set of scripts demonstrates "regex bomb" - issue with matching
certain types of regular expressions in popular implmenentations.

Specificially, implementations based on PCRE, suffer from exponential
time when matching certain valid expressions.

This experiment shows execution times for matching regular expression
 *a?<sup>n</sup>a<sup>n</sup>* against input string of *a<sup>n</sup>*.
(*n* denotes repetition, that is,
*a?<sup>2</sup>a<sup>2</sup>* is *a?a?aa*.

The outcome is that Go's library is free from this exponential
time characteristincs.

Please also see rsc's [regexp1] article.

## usage

`./run N1 N2`

runs checks in range N1...N2. Interesting things start above N=25.

[regexp1]: https://swtch.com/~rsc/regexp/regexp1.html
