package main

import (
	"io"
	"os"
	"strings"
)

// Applying ROT13 to a piece of text merely requires examining its alphabetic characters and
// replacing each one by the letter 13 places further along in the alphabet, wrapping back
// to the beginning if necessary.
// A becomes N, B becomes O, and so on up to M, which becomes Z, then the sequence continues
// at the beginning of the alphabet: N becomes A, O becomes B, and so on to Z, which becomes M.
type rot13Reader struct {
	in io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	l, err := r.in.Read(b)
	if err != nil {
		return 0, err
	}
	for i, x := range b {
		if (x >= 'A' && x <= 'M') || (x >= 'a' && x <= 'm') {
			b[i] += 13
		} else if (x >= 'N' && x <= 'Z') || (x >= 'n' && x <= 'z') {
			b[i] -= 13
		}
	}
	return l, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
