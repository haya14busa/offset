package offset

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFromOffset(t *testing.T) {
	const text = `MIT License

Copyright (c) 2020 haya14busa

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.`

	// Check `:<offset>go` in Vim. See :h go.
	tests := []struct {
		offset int
		want   Position
	}{
		{
			offset: 0,
			want:   Position{Offset: 0, Line: 1, Column: 0},
		},
		{
			offset: 1,
			want:   Position{Offset: 1, Line: 1, Column: 1},
		},
		{
			offset: 13,
			want:   Position{Offset: 13, Line: 2, Column: 1},
		},
		{
			offset: 14,
			want:   Position{Offset: 14, Line: 3, Column: 1},
		},
		{
			offset: 100,
			want:   Position{Offset: 100, Line: 5, Column: 56},
		},
	}

	for _, tt := range tests {
		got, err := FromOffset(strings.NewReader(text), tt.offset)
		if err != nil {
			t.Error(err)
			continue
		}
		if diff := cmp.Diff(got, tt.want); diff != "" {
			t.Error(diff)
		}
	}
}

func TestFromOffset_invalidOffset(t *testing.T) {
	if _, err := FromOffset(strings.NewReader("test offset"), 1000); err == nil {
		t.Error("want error, but got nil")
	}
}
