#!/usr/bin/env ruby

def an(n)
	[Regexp.new("a?"*n + "a"*n), "a"*n]
end

n = ARGV[0].to_i
p, s = an(n)
idxs = p.match(s).offset(0)
puts "n: #{n} match: [#{idxs.join(' ')}]" 
