package main

import (
	"bytes"
	"strings"
	"testing"
)

var tests = [][]string{
	{"2 3", "5"},
	{"0 0", "0"},
	{"1000 1000", "2000"},
}

func TestSolve(t *testing.T) {
	var b bytes.Buffer
	for i, test := range tests {
		solve(strings.NewReader(test[0]), &b)
		if b.String() != test[1] {
			t.Fatal("Errore test ", i, test[1], b.String())
		}
		b.Reset()
	}
}
