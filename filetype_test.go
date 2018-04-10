package filetype

import (
	"testing"
)

func TestIs(t *testing.T) {
	cases := []struct {
		buf   []byte
		ext   string
		match bool
	}{
		{[]byte{0xFF, 0xD8, 0xFF}, "jpg", true},
		{[]byte{0xFF, 0xD8, 0x00}, "jpg", false},
		{[]byte{0x89, 0x50, 0x4E, 0x47}, "png", true},
	}

	for _, test := range cases {
		if Is(test.buf, test.ext) != test.match {
			t.Fatalf("Invalid match: %s", test.ext)
		}
	}
}

func TestIsType(t *testing.T) {
	cases := []struct {
		buf   []byte
		kind  Typed
		match bool
	}{
		{[]byte{0xFF, 0xD8, 0xFF}, GetByExtension("jpg"), true},
		{[]byte{0xFF, 0xD8, 0x00}, GetByExtension("jpg"), false},
		{[]byte{0x89, 0x50, 0x4E, 0x47}, GetByExtension("png"), true},
	}

	for _, test := range cases {
		if IsType(test.buf, test.kind) != test.match {
			t.Fatalf("Invalid match: %s", test.kind.Extension)
		}
	}
}

func TestIsMIME(t *testing.T) {
	cases := []struct {
		buf   []byte
		mime  string
		match bool
	}{
		{[]byte{0xFF, 0xD8, 0xFF}, "image/jpeg", true},
		{[]byte{0xFF, 0xD8, 0x00}, "image/jpeg", false},
		{[]byte{0x89, 0x50, 0x4E, 0x47}, "image/png", true},
	}

	for _, test := range cases {
		if IsMIME(test.buf, test.mime) != test.match {
			t.Fatalf("Invalid match: %s", test.mime)
		}
	}
}

func TestIsSupported(t *testing.T) {
	cases := []struct {
		ext   string
		match bool
	}{
		{"jpg", true},
		{"jpeg", false},
		{"abc", false},
		{"png", true},
		{"mp4", true},
		{"", false},
	}

	for _, test := range cases {
		if IsSupported(test.ext) != test.match {
			t.Fatalf("Invalid match: %s", test.ext)
		}
	}
}

func TestIsMIMESupported(t *testing.T) {
	cases := []struct {
		mime  string
		match bool
	}{
		{"image/jpeg", true},
		{"foo/bar", false},
		{"image/png", true},
		{"video/mpeg", true},
	}

	for _, test := range cases {
		if IsMIMESupported(test.mime) != test.match {
			t.Fatalf("Invalid match: %s", test.mime)
		}
	}
}

func TestAddType(t *testing.T) {
	NewType("foo", "foo/foo")

	if !IsSupported("foo") {
		t.Fatalf("Not supported extension")
	}

	if !IsMIMESupported("foo/foo") {
		t.Fatalf("Not supported MIME type")
	}
}

func TestGetType(t *testing.T) {
	jpg := GetByExtension("jpg")
	if jpg == Unknown {
		t.Fatalf("Type should be supported")
	}

	invalid := GetByExtension("invalid")
	if invalid != Unknown {
		t.Fatalf("Type should not be supported")
	}
}
