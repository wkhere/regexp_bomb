///bin/true; exec go run "$0" "$@"
package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func an(n int) (r *regexp.Regexp, s string) {
	var buf1, buf2 bytes.Buffer
	for i := 0; i < n; i++ {
		fmt.Fprint(&buf1, "a?")
		fmt.Fprint(&buf2, "a")
	}
	s = buf2.String()
	fmt.Fprint(&buf1, s)
	r = regexp.MustCompile(buf1.String())
	return
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	r, s := an(n)
	fmt.Println("n:", n, "match:", r.FindStringIndex(s))
}
