package filetype

import (
	"testing"
)

func TestKindImage(t *testing.T) {
	var cases = []struct {
		buf []byte
		ext string
	}{
		{[]byte{0xFF, 0xD8, 0xFF}, "jpg"},
		{[]byte{0x89, 0x50, 0x4E, 0x47}, "png"},
		{[]byte{0x89, 0x0, 0x0}, "unknown"},
	}

	for _, test := range cases {
		kind, _ := MatchImage(test.buf)
		if kind.Extension != test.ext {
			t.Errorf("Invalid match: %s != %s", kind.Extension, test.ext)
		}
	}
}

func TestKindIsImage(t *testing.T) {
	var cases = []struct {
		buf   []byte
		match bool
	}{
		{[]byte{0xFF, 0xD8, 0xFF}, true},
		{[]byte{0x89, 0x50, 0x4E, 0x47}, true},
		{[]byte{0x89, 0x0, 0x0}, false},
	}

	for _, test := range cases {
		if IsImage(test.buf) != test.match {
			t.Errorf("Invalid match: %t", test.match)
		}
	}
}
