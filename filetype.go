package filetype

import (
	"errors"
)

// ErrEmptyBuffer represents an empty buffer error
var ErrEmptyBuffer = errors.New("Empty buffer")

// ErrUnknownBuffer represents a unknown buffer error
var ErrUnknownBuffer = errors.New("Unknown buffer type")

// Is checks if a given buffer matches with the given file type extension
func Is(buf []byte, ext string) bool {
	kind, ok := TypedMap[ext]
	if ok {
		return IsType(buf, kind)
	}
	return false
}

// IsExtension semantic alias to Is()
func IsExtension(buf []byte, ext string) bool {
	return Is(buf, ext)
}

// IsType checks if a given buffer matches with the given file type
func IsType(buf []byte, kind Typed) bool {
	matcher := Matchers[kind]
	if matcher == nil {
		return false
	}
	return matcher(buf) != Unknown
}

// IsMIME checks if a given buffer matches with the given MIME type
func IsMIME(buf []byte, mime string) bool {
	for _, kind := range TypedMap {
		if kind.MIME.Value == mime {
			matcher := Matchers[kind]
			return matcher(buf) != Unknown
		}
	}
	return false
}

// IsSupported checks if a given file extension is supported
func IsSupported(ext string) bool {
	for name := range TypedMap {
		if name == ext {
			return true
		}
	}
	return false
}

// IsMIMESupported checks if a given MIME type is supported
func IsMIMESupported(mime string) bool {
	for _, m := range TypedMap {
		if m.MIME.Value == mime {
			return true
		}
	}
	return false
}
