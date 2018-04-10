package filetype

import "strings"

// MIME stores the file MIME type values
type MIME struct {
	Type    string
	Subtype string
	Value   string
}

// Typed represents a file MIME type and its extension
type Typed struct {
	MIME      MIME
	Extension string
}

/*
Are maps cheaper to reference or make?
*/
// TypedMap of signs:Typed
//type TypedMap map[string]Typed
var TypedMap = make(map[string]Typed)

// Unknown for cases where file type can't be determined
var Unknown = NewType("unknown", "")

/*
// TypedMap of signs:Typed
type TypedMapType map[string]Typed
type UnknownType TypedMapType
var (
	Unknown   = make(UnknownType)
	TypedMake = make(TypedMapType)
)
//Unknown := NewType("unknown", "")
//Typed{MIME:      NewMIME(mime),Extension: ext}
*/

func splitMime(s string) (string, string) {
	x := strings.Split(s, "/")
	if len(x) > 1 {
		return x[0], x[1]
	}
	return x[0], ""
}

// Creates a new MIME type
func NewMIME(mime string) MIME {
	kind, subtype := splitMime(mime)
	return MIME{Type: kind, Subtype: subtype, Value: mime}
}

// Add registers a new type in the package
func Add(t Typed) Typed {
	TypedMap[t.Extension] = t
	return t
}

// GetByExtension retrieves a Type by extension
func GetByExtension(ext string) Typed {
	kind := TypedMap[ext]
	if kind.Extension != "" {
		return kind
	}
	return Unknown
}

// NewType creates a new Type
func NewType(ext, mime string) Typed {
	t := Typed{
		MIME:      NewMIME(mime),
		Extension: ext,
	}
	return Add(t)
}
