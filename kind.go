package filetype

// Image tries to match a file as image type
func MatchImage(buf []byte) (Typed, error) {
	return doMatchMap(buf, Image)
}

// IsImage checks if the given buffer is an image type
func IsImage(buf []byte) bool {
	kind, _ := MatchImage(buf)
	return kind != Unknown
}

// Audio tries to match a file as audio type
func MatchAudio(buf []byte) (Typed, error) {
	return doMatchMap(buf, Audio)
}

// IsAudio checks if the given buffer is an audio type
func IsAudio(buf []byte) bool {
	kind, _ := MatchAudio(buf)
	return kind != Unknown
}

// Video tries to match a file as video type
func MatchVideo(buf []byte) (Typed, error) {
	return doMatchMap(buf, Video)
}

// IsVideo checks if the given buffer is a video type
func IsVideo(buf []byte) bool {
	kind, _ := MatchVideo(buf)
	return kind != Unknown
}

// Font tries to match a file as text font type
func MatchFont(buf []byte) (Typed, error) {
	return doMatchMap(buf, Font)
}

// IsFont checks if the given buffer is a font type
func IsFont(buf []byte) bool {
	kind, _ := MatchFont(buf)
	return kind != Unknown
}

// Archive tries to match a file as generic archive type
func MatchArchive(buf []byte) (Typed, error) {
	return doMatchMap(buf, Archive)
}

// IsArchive checks if the given buffer is an archive type
func IsArchive(buf []byte) bool {
	kind, _ := MatchArchive(buf)
	return kind != Unknown
}

func doMatchMap(buf []byte, mmapFunc MapOfMatcherFunc) (Typed, error) {
	kind := MatchMap(buf, mmapFunc)
	if kind != Unknown {
		return kind, nil
	}
	return kind, ErrUnknownBuffer
}
