package gozbar

import (
	"errors"
)

// #include <zbar.h>
import "C"

var NoMemoryError = errors.New("zbar: Out of memory")
var InternalError = errors.New("zbar: Internal library error")
var UnsupportedError = errors.New("zbar: Unsupported request")
var InvalidError = errors.New("zbar: Invalid request")
var SystemError = errors.New("zbar: System error")
var LockingError = errors.New("zbar: Locking error")
var BusyError = errors.New("zbar: All resources busy")
var XDisplayError = errors.New("zbar: X11 display error")
var XProtoError = errors.New("zbar: X11 protocol error")
var ClosedError = errors.New("zbar: Output window is closed")
var WinAPIError = errors.New("zbar: Windows system error")

var zbarCodeToError = map[int]error{
	C.ZBAR_ERR_NOMEM:       NoMemoryError,
	C.ZBAR_ERR_INTERNAL:    InternalError,
	C.ZBAR_ERR_UNSUPPORTED: UnsupportedError,
	C.ZBAR_ERR_INVALID:     InvalidError,
	C.ZBAR_ERR_SYSTEM:      SystemError,
	C.ZBAR_ERR_LOCKING:     LockingError,
	C.ZBAR_ERR_BUSY:        BusyError,
	C.ZBAR_ERR_XDISPLAY:    XDisplayError,
	C.ZBAR_ERR_XPROTO:      XProtoError,
	C.ZBAR_ERR_CLOSED:      ClosedError,
	C.ZBAR_ERR_WINAPI:      WinAPIError,
}

func errorCodeToError(errorCode int) error {
	if errorCode == 0 {
		return nil
	} else if errorCode >= C.ZBAR_ERR_NUM {
		return errors.New("zbar: Unknown error code from zbar library")
	} else {
		return zbarCodeToError[errorCode]
	}
}
