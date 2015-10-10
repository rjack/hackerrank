package main

import (
	"fmt"
	"io"
	"math/big"
	"os"
)

func solve(r io.Reader, w io.Writer) {
	var n int64

	fmt.Fscan(r, &n)

	x := big.NewInt(n)
	fact := x.MulRange(2, n)

	fmt.Fprint(w, fact.String())
}

func main() {
	solve(os.Stdin, os.Stdout)
}
