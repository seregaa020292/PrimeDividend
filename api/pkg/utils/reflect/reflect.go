package reflect

import (
	"reflect"
	"strings"
)

// Value returns reflect value underlying given value, unwrapping pointer and slice
func Value(v interface{}) reflect.Value {
	rv, ok := v.(reflect.Value)
	if !ok {
		rv = reflect.ValueOf(v)
	}
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Slice {
		rv = rv.Elem()
	}
	return rv
}

// Type returns reflect type underlying given value, unwrapping pointer and slice
func Type(v interface{}) reflect.Type {
	var rt reflect.Type
	if rvt, ok := v.(reflect.Type); ok {
		rt = rvt
	} else {
		rv, ok := v.(reflect.Value)
		if !ok {
			rv = reflect.ValueOf(v)
		}
		rt = rv.Type()
	}
	for rt.Kind() == reflect.Ptr || rt.Kind() == reflect.Slice {
		rt = rt.Elem()
	}
	return rt
}

// FieldByPath returns StructField under a given path in a given value,
// unwrapping pointer and slice.
// If path has levels ("Parent.Field") then it will unwrap all levels and return type of the
// leaf field.
func FieldByPath(v interface{}, path string) (reflect.StructField, bool) {
	subkeys := strings.Split(path, ".")
	var subfield reflect.StructField
	for _, key := range subkeys {
		var ok bool
		if subfield.Type == nil {
			subfield, ok = Type(v).FieldByName(key)
		} else {
			if subfield.Type.Kind() == reflect.Ptr {
				subfield.Type = subfield.Type.Elem()
			}
			subfield, ok = subfield.Type.FieldByName(key)
		}
		if !ok {
			return reflect.StructField{}, false
		}
	}
	return subfield, true
}

// ValueByPath returns reflect value under a given path in a given value,
// unwrapping pointer and slice.
// If path has levels ("Parent.Field") then it will unwrap all levels and return value of the
// leaf field.
func ValueByPath(v interface{}, path string) reflect.Value {
	subkeys := strings.Split(path, ".")
	subfield := Value(v)
	for _, key := range subkeys {
		if subfield != (reflect.Value{}) {
			if subfield.Kind() == reflect.Ptr {
				subfield = subfield.Elem()
			}
			subfield = subfield.FieldByName(key)
		} else {
			return reflect.Value{}
		}
	}
	return subfield
}
