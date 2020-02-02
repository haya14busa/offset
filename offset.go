package offset

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Position describes an arbitrary position including the file, line, and
// column location.
type Position struct {
	File   string `json:"file,omitempty"`   // file name, if any
	Offset int    `json:"offset,omitempty"` // offset, starting at 0
	Line   int    `json:"line,omitempty"`   // line number, starting at 1
	Column int    `json:"column,omitempty"` // column number, starting at 1 (byte count)
}

// String returns a string in one of several forms:
//
//	file:line:column
//	file:line
//	line:column
//	line
//	file
func (pos Position) String() string {
	var b strings.Builder
	b.WriteString(pos.File)
	if pos.Line > 0 {
		if pos.File != "" {
			b.WriteRune(':')
		}
		b.WriteString(fmt.Sprintf("%d", pos.Line))
		if pos.Column != 0 {
			b.WriteString(fmt.Sprintf(":%d", pos.Column))
		}
	}
	return b.String()
}

func FromOffset(r io.Reader, offset int) (Position, error) {
	rem := offset // remaining bytes
	scanner := bufio.NewScanner(r)
	line := 0
	col := 0
	valid := false
	for scanner.Scan() {
		line++
		l := len(scanner.Bytes()) + 1
		if rem <= l {
			col = rem
			valid = true
			break
		}
		rem -= l
	}
	if valid {
		return Position{Offset: offset, Line: line, Column: col}, nil
	}
	return Position{}, fmt.Errorf("invalid offset=%d", offset)
}

func FromFilename(filename string, offset int) (Position, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Position{}, err
	}
	defer file.Close()
	return FromOffset(file, offset)
}
