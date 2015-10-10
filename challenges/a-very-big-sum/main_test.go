package main

import (
	"bytes"
	"strings"
	"testing"
)

var tests = [][]string{
	{
		`5
1000000001 1000000002 1000000003 1000000004 1000000005`,
		`5000000015`,
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
