package main

import (
	"fmt"
	"io"
	"os"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func newAlphabetMap() map[rune]struct{} {
	m := make(map[rune]struct{})
	for _, x := range alphabet {
		m[x] = struct{}{}
	}
	return m
}

func solve(r io.Reader, w io.Writer) {
	var s string

	m := newAlphabetMap()

	for {
		_, err := fmt.Fscan(r, &s)
		if err != nil {
			break
		}

		for _, x := range s {
			delete(m, x)
		}
	}
	fmt.Println(m)
	if len(m) > 0 {
		fmt.Fprint(w, "not pangram")
	} else {
		fmt.Fprint(w, "pangram")
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
