package main

import (
	"fmt"
	"io"
	"os"
)

func solve(r io.Reader, w io.Writer) {
	var (
		n   int
		sum uint64
	)

	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		var a uint64
		fmt.Fscan(r, &a)
		sum += a
	}
	fmt.Fprint(w, sum)
}

func main() {
	solve(os.Stdin, os.Stdout)
}
