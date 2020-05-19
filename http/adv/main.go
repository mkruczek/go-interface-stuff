// This program demonstrates why, if you're reading text files one line
// at a time, and the lines may be 64k or longer, you need to check for
// a too-long line after the scanner returns false, or your program will
// silently ignore some of its input. This program also shows how to
// increase the default 64k buffer size (using scan.Buffer()).
// (https://github.com/LarryRuane)

package main

import (
	"bufio"
	"fmt"
)

type (
	state struct {
		lineLen int // does not include newline
		n       int // number of 'a's we've returned
	}
	longLine struct{ s *state }
)

// Return a text line of exactly lineLen 'a' characters
func (l longLine) Read(p []byte) (n int, err error) {
	var i int
	for i = 0; i < len(p); i++ {
		if l.s.n == l.s.lineLen {
			p[i] = byte('\n')
			return i + 1, nil
		}
		p[i] = byte('a')
		l.s.n++
	}
	return i, nil
}

func main() {
	// lines with lengths 65535 are okay
	{
		s := state{lineLen: 64*1024 - 1}
		scan := bufio.NewScanner(longLine{&s})
		for scan.Scan() {
			fmt.Println("should read line length 64k-1:", len(scan.Text()))
			break
		}
		if scan.Err() != nil {
			fmt.Println("unexpected failure:", scan.Err())
		}
	}

	// lines with lengths 65536 don't work (need one byte for newline)
	{
		s := state{lineLen: 64 * 1024}
		scan := bufio.NewScanner(longLine{&s})
		for scan.Scan() {
			fmt.Println("not expected to reach here")
			break
		}
		fmt.Print("expecting bufio.Scanner: token too long: >>>> ")
		if scan.Err() != nil {
			fmt.Print(scan.Err())
		} else {
			fmt.Print("ERROR did not happen!")
		}
		fmt.Println(" <<<<")
	}

	// but if we create a buffer, lines with lengths 65536 are okay
	{
		s := state{lineLen: 64 * 1024}
		scan := bufio.NewScanner(longLine{&s})
		var buf []byte
		scan.Buffer(buf, 64*1024+1)
		for scan.Scan() {
			fmt.Println("custom buffer, should read line length 64k:", len(scan.Text()))
			break
		}
		if scan.Err() != nil {
			fmt.Println("unexpected failure:", scan.Err())
		}
	}
}
