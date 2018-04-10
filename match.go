package filetype

import (
	"io"
	"os"
)

// Matcher function interface as type alias
type MatcherFunc func([]byte) bool

// Type interface to store pairs of type with its matcher function
type MapOfMatcherFunc map[Typed]MatcherFunc

// Type specific matcher function interface
type TypeMatcherFunc func([]byte) Typed

// Store registered file type matchers
var Matchers = make(map[Typed]TypeMatcherFunc)

func register(mmapFunc ...MapOfMatcherFunc) {
	for _, m := range mmapFunc {
		for kind, matchr := range m {
			NewMatcher(kind, matchr)
		}
	}
}

func init() {
	// Arguments order is intentional
	register(Image, Video, Audio, Font, Document, Archive)
}

// AddMatcher registers a new matcher type
//func AddMatcher(fileType Typed, mFunc MatcherFunc) TypeMatcherFunc {
//	return NewMatcher(fileType, mFunc)
//}

// Create and register a new type matcher function
func NewMatcher(t Typed, fn MatcherFunc) TypeMatcherFunc {
	matcher := func(buf []byte) Typed {
		if fn(buf) {
			return t
		}
		return Unknown
	}

	Matchers[t] = matcher
	return matcher
}

// Match infers the file type of a given buffer inspecting its magic numbers signature
func MatchBuffer(buf []byte) (Typed, error) {
	length := len(buf)
	if length == 0 {
		return Unknown, ErrEmptyBuffer
	}

	for _, checker := range Matchers {
		match := checker(buf)
		if match != Unknown && match.Extension != "" {
			return match, nil
		}
	}

	return Unknown, nil
}

// MatchReader is convenient wrapper to MatchBuffer() any Reader
func MatchReader(reader io.Reader) (Typed, error) {
	buffer := make([]byte, 512)

	_, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		return Unknown, err
	}

	return MatchBuffer(buffer)
}

// MatchFile infers a file type for a file
func MatchFile(filepath string) (Typed, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return Unknown, err
	}
	defer file.Close()

	return MatchReader(file)
}

// Matches checks if the given buffer matches with some supported file type
func Matches(buf []byte) bool {
	kind, _ := MatchBuffer(buf)
	return kind != Unknown
}

// MatchMap performs a file matching againts a map of match functions
func MatchMap(buf []byte, mmapFunc MapOfMatcherFunc) Typed {
	for kind, matcher := range mmapFunc {
		if matcher(buf) {
			return kind
		}
	}
	return Unknown
}
