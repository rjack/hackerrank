package main

import (
	"fmt"
	"io"
	"os"
)

func solve(r io.Reader, w io.Writer) {
	var n, sum int

	fmt.Fscan(r, &n)

	for col := 0; col < n; col++ {
		for row := 0; row < n; row++ {
			var a int
			fmt.Fscan(r, &a)
			if row == col {
				sum += a
			}
			if n-1-col == row {
				sum -= a
			}
		}
	}
	if sum < 0 {
		sum *= -1
	}
	fmt.Fprint(w, sum)
}

func main() {
	solve(os.Stdin, os.Stdout)
}
