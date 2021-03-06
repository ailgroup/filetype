package filetype

var (
	TypeJpeg = NewType("jpg", "image/jpeg")
	TypePng  = NewType("png", "image/png")
	TypeGif  = NewType("gif", "image/gif")
	TypeWebp = NewType("webp", "image/webp")
	TypeCR2  = NewType("cr2", "image/x-canon-cr2")
	TypeTiff = NewType("tif", "image/tiff")
	TypeBmp  = NewType("bmp", "image/bmp")
	TypeJxr  = NewType("jxr", "image/vnd.ms-photo")
	TypePsd  = NewType("psd", "image/vnd.adobe.photoshop")
	TypeIco  = NewType("ico", "image/x-icon")
)

var Image = MapOfMatcherFunc{
	TypeJpeg: Jpeg,
	TypePng:  Png,
	TypeGif:  Gif,
	TypeWebp: Webp,
	TypeCR2:  CR2,
	TypeTiff: Tiff,
	TypeBmp:  Bmp,
	TypeJxr:  Jxr,
	TypePsd:  Psd,
	TypeIco:  Ico,
}

func Jpeg(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0xFF &&
		buf[1] == 0xD8 &&
		buf[2] == 0xFF
}

func Png(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x89 && buf[1] == 0x50 &&
		buf[2] == 0x4E && buf[3] == 0x47
}

func Gif(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x47 && buf[1] == 0x49 && buf[2] == 0x46
}

func Webp(buf []byte) bool {
	return len(buf) > 11 &&
		buf[8] == 0x57 && buf[9] == 0x45 &&
		buf[10] == 0x42 && buf[11] == 0x50
}

func CR2(buf []byte) bool {
	return len(buf) > 9 &&
		((buf[0] == 0x49 && buf[1] == 0x49 && buf[2] == 0x2A && buf[3] == 0x0) ||
			(buf[0] == 0x4D && buf[1] == 0x4D && buf[2] == 0x0 && buf[3] == 0x2A)) &&
		buf[8] == 0x43 && buf[9] == 0x52
}

func Tiff(buf []byte) bool {
	return len(buf) > 3 &&
		((buf[0] == 0x49 && buf[1] == 0x49 && buf[2] == 0x2A && buf[3] == 0x0) ||
			(buf[0] == 0x4D && buf[1] == 0x4D && buf[2] == 0x0 && buf[3] == 0x2A))
}

func Bmp(buf []byte) bool {
	return len(buf) > 1 &&
		buf[0] == 0x42 &&
		buf[1] == 0x4D
}

func Jxr(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x49 &&
		buf[1] == 0x49 &&
		buf[2] == 0xBC
}

func Psd(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x38 && buf[1] == 0x42 &&
		buf[2] == 0x50 && buf[3] == 0x53
}

func Ico(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x00 && buf[1] == 0x00 &&
		buf[2] == 0x01 && buf[3] == 0x00
}
