package nson

import (
	"fmt"
	"reflect"
	"time"
)

// Marshal 将结构体序列化为 nson.Map
func Marshal(v any) (Map, error) {
	rv := reflect.ValueOf(v)

	// 处理指针
	for rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil, fmt.Errorf("cannot marshal nil pointer")
		}
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %v", rv.Kind())
	}

	return marshalStruct(rv)
}

// marshalStruct 将结构体序列化为 Map
func marshalStruct(rv reflect.Value) (Map, error) {
	t := rv.Type()
	cache := getStructCache(t)

	m := make(Map, len(cache.fields))

	for _, field := range cache.fields {
		fv := rv
		// 通过索引路径获取字段值
		for _, idx := range field.indices {
			fv = fv.Field(idx)
		}

		// 处理 omitempty
		if field.omitEmpty && isEmptyValue(fv) {
			continue
		}

		val, err := marshalValue(fv)
		if err != nil {
			return nil, fmt.Errorf("field %s: %w", field.name, err)
		}

		if val != nil {
			m[field.nsonName] = val
		}
	}

	return m, nil
}

// marshalValue 将 reflect.Value 转换为 nson.Value
func marshalValue(rv reflect.Value) (Value, error) {
	// 处理指针
	for rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return Null{}, nil
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Bool:
		return Bool(rv.Bool()), nil

	case reflect.Int8:
		return I8(rv.Int()), nil

	case reflect.Int16:
		return I16(rv.Int()), nil

	case reflect.Int32, reflect.Int:
		return I32(rv.Int()), nil

	case reflect.Int64:
		return I64(rv.Int()), nil

	case reflect.Uint8:
		return U8(rv.Uint()), nil

	case reflect.Uint16:
		return U16(rv.Uint()), nil

	case reflect.Uint32, reflect.Uint:
		return U32(rv.Uint()), nil

	case reflect.Uint64:
		return U64(rv.Uint()), nil

	case reflect.Float32:
		return F32(rv.Float()), nil

	case reflect.Float64:
		return F64(rv.Float()), nil

	case reflect.String:
		return String(rv.String()), nil

	case reflect.Slice:
		if rv.Type().Elem().Kind() == reflect.Uint8 {
			// []byte
			return Binary(rv.Bytes()), nil
		}
		return marshalSlice(rv)

	case reflect.Array:
		// 检查是否是 nson.Id 类型（[12]byte）
		if rv.Type() == reflect.TypeFor[Id]() {
			var id Id
			reflect.Copy(reflect.ValueOf(&id).Elem(), rv)
			return id, nil
		}
		return marshalArray(rv)

	case reflect.Struct:
		// 检查是否是 time.Time 类型
		if rv.Type() == reflect.TypeFor[time.Time]() {
			t := rv.Interface().(time.Time)
			// 转换为毫秒时间戳
			return Timestamp(t.UnixMilli()), nil
		}
		return marshalStruct(rv)

	case reflect.Map:
		return marshalMap(rv)

	case reflect.Interface:
		if rv.IsNil() {
			return Null{}, nil
		}
		// 检查是否实现了 nson.Value 接口
		if rv.Type().Implements(reflect.TypeFor[Value]()) {
			// 直接返回 nson.Value
			if val, ok := rv.Interface().(Value); ok {
				return val, nil
			}
		}
		return marshalValue(rv.Elem())

	default:
		return nil, fmt.Errorf("unsupported type: %v", rv.Type())
	}
}

// marshalSlice 序列化切片
func marshalSlice(rv reflect.Value) (Array, error) {
	arr := make(Array, 0, rv.Len())

	for i := 0; i < rv.Len(); i++ {
		val, err := marshalValue(rv.Index(i))
		if err != nil {
			return nil, fmt.Errorf("index %d: %w", i, err)
		}
		arr = append(arr, val)
	}

	return arr, nil
}

// marshalArray 序列化数组
func marshalArray(rv reflect.Value) (Array, error) {
	arr := make(Array, 0, rv.Len())

	for i := 0; i < rv.Len(); i++ {
		val, err := marshalValue(rv.Index(i))
		if err != nil {
			return nil, fmt.Errorf("index %d: %w", i, err)
		}
		arr = append(arr, val)
	}

	return arr, nil
}

// marshalMap 序列化 map
func marshalMap(rv reflect.Value) (Map, error) {
	if rv.Type().Key().Kind() != reflect.String {
		return nil, fmt.Errorf("map key must be string")
	}

	m := make(Map, rv.Len())

	iter := rv.MapRange()
	for iter.Next() {
		key := iter.Key().String()
		val, err := marshalValue(iter.Value())
		if err != nil {
			return nil, fmt.Errorf("key %s: %w", key, err)
		}
		m[key] = val
	}

	return m, nil
}

// isEmptyValue 检查值是否为空（用于 omitempty）
func isEmptyValue(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return rv.IsNil()
	}
	return false
}
