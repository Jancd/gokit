package chars

import (
	"unsafe"
)

// BytesToString is a hack style transformer
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes is a hack style transformer
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
