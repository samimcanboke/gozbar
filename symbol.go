package gozbar

import (
	"image"
)

// #include <zbar.h>
import "C"

type SymbolType int

const (
	None    SymbolType = C.ZBAR_NONE
	Partial SymbolType = C.ZBAR_PARTIAL
	EAN8    SymbolType = C.ZBAR_EAN8
	UPCE    SymbolType = C.ZBAR_UPCE
	ISBN10  SymbolType = C.ZBAR_ISBN10
	UPCA    SymbolType = C.ZBAR_UPCA
	EAN13   SymbolType = C.ZBAR_EAN13
	ISBN13  SymbolType = C.ZBAR_ISBN13
	I25     SymbolType = C.ZBAR_I25
	Code39  SymbolType = C.ZBAR_CODE39
	PDF417  SymbolType = C.ZBAR_PDF417
	QRCode  SymbolType = C.ZBAR_QRCODE
	Code128 SymbolType = C.ZBAR_CODE128
)

// Name returns the name of a given symbol encoding type, or "UNKNOWN"
// if the encoding is not recognized.
func (s SymbolType) Name() string {
	return C.GoString(C.zbar_get_symbol_name(s.toEnum()))
}

// Quick conversion function to deal with C functions that want an
// enum type.
func (s SymbolType) toEnum() C.zbar_symbol_type_t {
	return C.zbar_symbol_type_t(s)
}

// Symbol represents a scanned barcode.
type Symbol struct {
	Type SymbolType
	Data string

	// Quality is an unscaled, relative quantity which expresses the
	// confidence of the match.  These values are currently meaningful
	// only in relation to each other: a larger value is more
	// confident than a smaller one.
	Quality int

	// Boundary is a set of image.Point which define a polygon
	// containing the scanned symbol in the image.
	Boundary []image.Point
}
