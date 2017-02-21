package circuChk

import (
	"reflect"
	"unsafe"
)

// 循環あり true
// 循環なし false

func CirculationCheck(v reflect.Value) bool {
	seen := make(map[comparison]bool)
	result := circulationCheck(v, seen)
	return result
}

type comparison struct {
	x unsafe.Pointer
	t reflect.Type
}

func circulationCheck(v reflect.Value, seen map[comparison]bool) bool {

	if v.CanAddr() {
		vptr := unsafe.Pointer(v.UnsafeAddr())
		c := comparison{vptr, v.Type()}

		if seen[c] {
			// 循環あり
			return true
		}
		seen[c] = true
	}

	result := false
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			result = circulationCheck(v.Field(i), seen)
			if result == true {
				break
			}
		}
	}

	return result

}
