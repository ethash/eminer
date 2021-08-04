package nvml

/*
#include "nvmlbridge.h"
*/
import "C"

func genCStringBuffer(size uint) *C.char {
	buf := make([]byte, size)
	return C.CString(string(buf))
}

func strndup(cs *C.char, len uint) string {
	return C.GoStringN(cs, C.int(C.strnlen(cs, C.size_t(len))))
}
