// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"unicode/utf8"
)

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//!-Marshal

var length int = 0

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf *bytes.Buffer, v reflect.Value) error {

	if checkIsNotZero(v) == false {
		return nil
	}

	switch v.Kind() {
	case reflect.Invalid:

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
				fmt.Fprintf(buf, "\n")
			}
			if i != 0 {
				for i := 0; i < length; i++ {
					buf.WriteString(" ")
				}
			}

			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // ((name value) ...)

		if v.NumField() == 0 {
			return nil
		}

		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			// ここから
			if checkIsNotZero(v.Field(i)) == false {
				continue
			}

			if i != 0 {
				fmt.Fprintf(buf, "\n")
			}
			fmt.Fprintf(buf, "\"%s\":", v.Type().Field(i).Name)
			str := v.Type().Field(i).Name
			length = utf8.RuneCountInString(str) + 3
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			length = 0

			if i != v.NumField()-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if i != 0 {
				fmt.Fprintf(buf, "\n")
				for i := 0; i < length; i++ {
					buf.WriteString(" ")
				}
			}

			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			if v.Len()-1 != i {
				buf.WriteByte(',')
			}
		}
		buf.UnreadByte()
		buf.WriteByte('}')

		// ここから追加
	case reflect.Bool:
		val := v.Bool()

		if val == true {
			buf.WriteString("t")
		} else {
			buf.WriteString("null")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		cmp := v.Complex()

		fmt.Fprintf(buf, "#C(%f %f)", real(cmp), imag(cmp))

	case reflect.Interface:
		fmt.Fprintf(buf, "('%s' (%s))", v.Type(), v.Type().String())

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

// ゼロ値の場合は0を返す.
func checkIsNotZero(v reflect.Value) bool {

	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Ptr:
		if v.IsNil() == true {
			return false
		}
	case reflect.String:
		if v.String() == "" {
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if v.Uint() == 0 {
			return false
		}
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		if v.Int() == 0 {
			return false
		}
	}

	return true
}

//!-encode