package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type matrixLayers struct {
	layers   [][]uint64
	heads    []int
	tails    []int
	rows     int
	cols     int
	halfRows int
	halfCols int
	nlayers  int
}

func (l *matrixLayers) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%v layers:\n", len(l.layers))
	for i, v := range l.layers {
		fmt.Fprintf(&b, "%v: [", i)
		for j, w := range v {
			if j == l.heads[i] {
				fmt.Fprint(&b, "h")
			}
			if j == l.tails[i] {
				fmt.Fprint(&b, "t")
			}
			fmt.Fprintf(&b, "%v ", w)
		}
		fmt.Fprintf(&b, "]\n")
	}
	return b.String()
}

func (l *matrixLayers) layerLength(layer int) int {
	return 2 * ((l.rows - layer*2) + (l.cols - 2 - layer*2))
}

func newMatrixLayers(rows, cols int) *matrixLayers {

	l := &matrixLayers{
		rows:     rows,
		cols:     cols,
		halfRows: rows / 2,
		halfCols: cols / 2,
	}

	nl := cols
	if rows < cols {
		nl = rows
	}
	nl = nl / 2

	l.layers = make([][]uint64, nl)
	l.heads = make([]int, nl)
	l.tails = make([]int, nl)
	for i := 0; i < nl; i++ {
		llen := l.layerLength(i)
		l.layers[i] = make([]uint64, llen)
		l.tails[i] = llen - 1
	}
	l.nlayers = nl

	return l
}

func (l *matrixLayers) layerIndex(row, col int) int {
	if row >= l.halfRows {
		row = l.rows - 1 - row
	}
	if col >= l.halfCols {
		col = l.cols - 1 - col
	}
	min := row
	if col < min {
		min = col
	}
	return min
}

func (l *matrixLayers) rotate(rotations uint64) {
	for i := 0; i < l.nlayers; i++ {
		llen := len(l.layers[i])
		head := int(rotations % uint64(llen))
		tail := (head - 1)
		if tail < 0 {
			tail = llen - 1
		}
		l.heads[i] = head
		l.tails[i] = tail
	}
}

func (l *matrixLayers) coords(row, col int) (int, int) {
	i := l.layerIndex(row, col)
	var head bool
	if l.rows > l.cols {
		if col < l.halfCols {
			if row <= col {
				head = true
			} else {
				head = false
			}
		} else {
			if l.rows-row > l.cols-col {
				head = true
			} else {
				head = false
			}
		}
	} else {
		if row < l.halfRows {
			if row <= col {
				head = true
			} else {
				head = false
			}
		} else {
			if l.rows-row > l.cols-col {
				head = true
			} else {
				head = false
			}
		}
	}
	if head {
		h := l.heads[i]
		l.heads[i] = (l.heads[i] + 1) % len(l.layers[i])
		return i, h
	}
	t := l.tails[i]
	l.tails[i]--
	if l.tails[i] < 0 {
		l.tails[i] = len(l.layers[i]) - 1
	}
	return i, t
}

func (l *matrixLayers) set(row, col int, a uint64) {
	i, j := l.coords(row, col)
	l.layers[i][j] = a
}

func (l *matrixLayers) get(row, col int) (a uint64) {
	i, j := l.coords(row, col)
	return l.layers[i][j]
}

func solve(r io.Reader, w io.Writer) {
	var rows, cols int
	var rotations uint64

	fmt.Fscan(r, &rows, &cols, &rotations)

	layers := newMatrixLayers(rows, cols)

	//	fmt.Fprintln(debug, "JUST CREATED")
	//	fmt.Fprintln(debug, layers)
	//	fmt.Fprintln(debug, "READING INPUT")
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			var a uint64
			fmt.Fscan(r, &a)
			//			fmt.Fprintf(debug, "read: %v (%v,%v)\n", a, row, col)
			layers.set(row, col, a)
			//			fmt.Fprintln(debug, layers)
		}
	}

	layers.rotate(rotations)
	//	fmt.Fprintln(debug, "JUST ROTATED BY", rotations)
	//	fmt.Fprintln(debug, layers)
	//	fmt.Fprintln(debug, "WRITING OUTPUT")
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			a := layers.get(row, col)
			fmt.Fprint(w, a)
			if col < cols-1 {
				fmt.Fprint(w, " ")
			}
			//			fmt.Fprintln(debug, layers)
		}
		if row < rows-1 {
			fmt.Fprint(w, "\n")
		}
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
