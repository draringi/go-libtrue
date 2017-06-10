package truth

// #cgo CFLAGS: -I../libtrue
// #cgo LDFLAGS: -L../libtrue -ltrue
// #include <stdbool.h>
// #include <true.h>
//
// int bool_to_int(bool b) {
//     return b;
// }
import "C"

// GetTrue returns the truth
// You can't handle the truth.
func GetTrue() bool {
	val := C.get_true()
	converted := C.bool_to_int(val)
	if converted == 0 {
		return false
	}
	return true
}
