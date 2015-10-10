package main

import (
	"fmt"
	"io"
	"os"
)

func solve(r io.Reader, w io.Writer) {
	var a, b int

	fmt.Fscan(r, &a, &b)
	fmt.Fprint(w, a+b)
}

func main() {
	solve(os.Stdin, os.Stdout)
}
