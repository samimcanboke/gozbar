package gozbar

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

// Version returns the major and minor version numbers of the
// underlying zbar library.
func Version() (major, minor, patch uint) {
	var raw_major, raw_minor, raw_patch C.uint
	C.zbar_version(&raw_major, &raw_minor, &raw_patch)
	return uint(raw_major), uint(raw_minor), uint(raw_patch)
}

// SetVerbosity sets the library debug level.  Higher values spew more
// debug information.
func SetVerbosity(verbosity int) {
	C.zbar_set_verbosity(C.int(verbosity))
}

// Increases the library debug level.
func IncreaseVerbosity() {
	C.zbar_increase_verbosity()
}
