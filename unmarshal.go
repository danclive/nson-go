package nson

import (
	"fmt"
	"reflect"
	"time"
)

// Unmarshal 将 nson.Map 反序列化到结构体
func Unmarshal(m Map, v any) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Pointer {
		return fmt.Errorf("expected pointer, got %v", rv.Kind())
	}

	if rv.IsNil() {
		return fmt.Errorf("cannot unmarshal into nil pointer")
	}

	rv = rv.Elem()

	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("expected pointer to struct, got pointer to %v", rv.Kind())
	}

	return unmarshalStruct(m, rv)
}

// unmarshalStruct 将 Map 反序列化到结构体
func unmarshalStruct(m Map, rv reflect.Value) error {
	t := rv.Type()
	cache := getStructCache(t)

	for _, field := range cache.fields {
		val, has := m[field.nsonName]
		if !has {
			continue
		}

		fv := rv
		// 通过索引路径获取字段值
		for _, idx := range field.indices {
			fv = fv.Field(idx)
		}

		if !fv.CanSet() {
			continue
		}

		if err := unmarshalValue(val, fv); err != nil {
			return fmt.Errorf("field %s: %w", field.name, err)
		}
	}

	return nil
}

// unmarshalValue 将 nson.Value 反序列化到 reflect.Value
func unmarshalValue(val Value, rv reflect.Value) error {
	// 处理 nil 值
	if _, ok := val.(Null); ok {
		if rv.Kind() == reflect.Pointer {
			rv.Set(reflect.Zero(rv.Type()))
			return nil
		}
		// 对于非指针类型，保持零值
		return nil
	}

	// 处理指针
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		return unmarshalValue(val, rv.Elem())
	}

	switch rv.Kind() {
	case reflect.Bool:
		if v, ok := val.(Bool); ok {
			rv.SetBool(bool(v))
			return nil
		}
		return fmt.Errorf("expected Bool, got %T", val)

	case reflect.Int8:
		if v, ok := val.(I8); ok {
			rv.SetInt(int64(v))
			return nil
		}
		return fmt.Errorf("expected I8, got %T", val)

	case reflect.Int16:
		if v, ok := val.(I16); ok {
			rv.SetInt(int64(v))
			return nil
		}
		return fmt.Errorf("expected I16, got %T", val)

	case reflect.Int32, reflect.Int:
		if v, ok := val.(I32); ok {
			rv.SetInt(int64(v))
			return nil
		}
		return fmt.Errorf("expected I32, got %T", val)

	case reflect.Int64:
		switch v := val.(type) {
		case I64:
			rv.SetInt(int64(v))
		case Timestamp:
			rv.SetInt(int64(v))
		default:
			return fmt.Errorf("expected I64 or Timestamp, got %T", val)
		}
		return nil

	case reflect.Uint8:
		if v, ok := val.(U8); ok {
			rv.SetUint(uint64(v))
			return nil
		}
		return fmt.Errorf("expected U8, got %T", val)

	case reflect.Uint16:
		if v, ok := val.(U16); ok {
			rv.SetUint(uint64(v))
			return nil
		}
		return fmt.Errorf("expected U16, got %T", val)

	case reflect.Uint32, reflect.Uint:
		if v, ok := val.(U32); ok {
			rv.SetUint(uint64(v))
			return nil
		}
		return fmt.Errorf("expected U32, got %T", val)

	case reflect.Uint64:
		switch v := val.(type) {
		case U64:
			rv.SetUint(uint64(v))
		case Timestamp:
			rv.SetUint(uint64(v))
		default:
			return fmt.Errorf("expected U64 or Timestamp, got %T", val)
		}
		return nil

	case reflect.Float32:
		if v, ok := val.(F32); ok {
			rv.SetFloat(float64(v))
			return nil
		}
		return fmt.Errorf("expected F32, got %T", val)

	case reflect.Float64:
		if v, ok := val.(F64); ok {
			rv.SetFloat(float64(v))
			return nil
		}
		return fmt.Errorf("expected F64, got %T", val)

	case reflect.String:
		if v, ok := val.(String); ok {
			rv.SetString(string(v))
			return nil
		}
		return fmt.Errorf("expected String, got %T", val)

	case reflect.Slice:
		// 检查是否是 nson.Id 类型
		if rv.Type() == reflect.TypeOf(Id(nil)) {
			if v, ok := val.(Id); ok {
				rv.SetBytes([]byte(v))
				return nil
			}
			return fmt.Errorf("expected Id, got %T", val)
		}
		if rv.Type().Elem().Kind() == reflect.Uint8 {
			// []byte
			if v, ok := val.(Binary); ok {
				rv.SetBytes([]byte(v))
				return nil
			}
			return fmt.Errorf("expected Binary, got %T", val)
		}

		arr, ok := val.(Array)
		if !ok {
			return fmt.Errorf("expected Array, got %T", val)
		}

		slice := reflect.MakeSlice(rv.Type(), len(arr), len(arr))
		for i, item := range arr {
			if err := unmarshalValue(item, slice.Index(i)); err != nil {
				return fmt.Errorf("index %d: %w", i, err)
			}
		}
		rv.Set(slice)
		return nil

	case reflect.Array:
		arr, ok := val.(Array)
		if !ok {
			return fmt.Errorf("expected Array, got %T", val)
		}

		if len(arr) != rv.Len() {
			return fmt.Errorf("array length mismatch: expected %d, got %d", rv.Len(), len(arr))
		}

		for i, item := range arr {
			if err := unmarshalValue(item, rv.Index(i)); err != nil {
				return fmt.Errorf("index %d: %w", i, err)
			}
		}
		return nil

	case reflect.Struct:
		// 检查是否是 time.Time 类型
		if rv.Type() == reflect.TypeOf(time.Time{}) {
			if v, ok := val.(Timestamp); ok {
				// 从毫秒时间戳转换为 time.Time
				t := time.UnixMilli(int64(v))
				rv.Set(reflect.ValueOf(t))
				return nil
			}
			return fmt.Errorf("expected Timestamp for time.Time, got %T", val)
		}
		m, ok := val.(Map)
		if !ok {
			return fmt.Errorf("expected Map, got %T", val)
		}
		return unmarshalStruct(m, rv)

	case reflect.Map:
		// 检查是否是 nson.Map 类型
		if rv.Type() == reflect.TypeOf(Map{}) {
			m, ok := val.(Map)
			if !ok {
				return fmt.Errorf("expected Map, got %T", val)
			}
			rv.Set(reflect.ValueOf(m))
			return nil
		}

		if rv.Type().Key().Kind() != reflect.String {
			return fmt.Errorf("map key must be string")
		}

		m, ok := val.(Map)
		if !ok {
			return fmt.Errorf("expected Map, got %T", val)
		}

		mapVal := reflect.MakeMap(rv.Type())
		for key, item := range m {
			elemVal := reflect.New(rv.Type().Elem()).Elem()
			if err := unmarshalValue(item, elemVal); err != nil {
				return fmt.Errorf("key %s: %w", key, err)
			}
			mapVal.SetMapIndex(reflect.ValueOf(key), elemVal)
		}
		rv.Set(mapVal)
		return nil

	case reflect.Interface:
		// 对于 interface{}，直接设置对应的 Go 类型
		rv.Set(reflect.ValueOf(nsonToInterface(val)))
		return nil

	default:
		return fmt.Errorf("unsupported type: %v", rv.Type())
	}
}

// nsonToInterface 将 nson.Value 转换为 Go 原生类型
func nsonToInterface(val Value) any {
	switch v := val.(type) {
	case Bool:
		return bool(v)
	case I8:
		return int8(v)
	case I16:
		return int16(v)
	case I32:
		return int32(v)
	case I64:
		return int64(v)
	case U8:
		return uint8(v)
	case U16:
		return uint16(v)
	case U32:
		return uint32(v)
	case U64:
		return uint64(v)
	case F32:
		return float32(v)
	case F64:
		return float64(v)
	case String:
		return string(v)
	case Binary:
		return []byte(v)
	case Timestamp:
		return uint64(v)
	case Id:
		return []byte(v)
	case Array:
		arr := make([]any, len(v))
		for i, item := range v {
			arr[i] = nsonToInterface(item)
		}
		return arr
	case Map:
		m := make(map[string]any, len(v))
		for key, item := range v {
			m[key] = nsonToInterface(item)
		}
		return m
	case Null:
		return nil
	default:
		return nil
	}
}
