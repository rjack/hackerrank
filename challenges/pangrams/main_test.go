package main

import (
	"bytes"
	"strings"
	"testing"
)

var tests = [][]string{
	{
		`We promptly judged antique ivory buckles for the next prize    `,
		`pangram`,
	},
	{
		`We promptly judged antique ivory buckles for the prize    `,
		`not pangram`,
	},
}

func TestSolve(t *testing.T) {
	var b bytes.Buffer
	for i, test := range tests {
		solve(strings.NewReader(test[0]), &b)
		if b.String() != test[1] {
			t.Fatalf("Errore test #%v\nAtteso:\n-----\n%s\n-----\nRicevuto:\n-----\n%s\n-----\n", i, test[1], b.String())
		}
		b.Reset()
	}
}
