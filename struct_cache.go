package nson

import (
	"reflect"
	"sync"
)

// structCache 缓存结构体的字段信息以提高性能
type structCache struct {
	fields []fieldInfo
}

type fieldInfo struct {
	indices   []int // 字段索引路径（支持嵌入字段）
	name      string
	nsonName  string
	typ       reflect.Type
	omitEmpty bool
}

var (
	structCacheMutex sync.RWMutex
	structCacheMap   = make(map[reflect.Type]*structCache)
)

// getStructCache 获取或创建结构体缓存
func getStructCache(t reflect.Type) *structCache {
	structCacheMutex.RLock()
	cache, ok := structCacheMap[t]
	structCacheMutex.RUnlock()

	if ok {
		return cache
	}

	structCacheMutex.Lock()
	defer structCacheMutex.Unlock()

	// 双重检查
	if cache, ok := structCacheMap[t]; ok {
		return cache
	}

	cache = buildStructCache(t)
	structCacheMap[t] = cache
	return cache
}

// buildStructCache 构建结构体缓存
func buildStructCache(t reflect.Type) *structCache {
	cache := &structCache{
		fields: make([]fieldInfo, 0, t.NumField()),
	}

	buildFieldsRecursive(t, nil, &cache.fields)

	return cache
}

// buildFieldsRecursive 递归构建字段列表（支持嵌入字段）
func buildFieldsRecursive(t reflect.Type, indexPrefix []int, fields *[]fieldInfo) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// 跳过未导出的字段
		if !field.IsExported() {
			continue
		}

		indices := append(indexPrefix, i)

		// 处理匿名嵌入字段
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			// 递归处理嵌入结构体
			buildFieldsRecursive(field.Type, indices, fields)
			continue
		}

		// 获取 nson tag
		tag := field.Tag.Get("nson")
		if tag == "-" {
			continue
		}

		nsonName := tag
		omitEmpty := false

		// 解析 tag，支持 "name,omitempty"
		if tag != "" {
			for j, part := range []byte(tag) {
				if part == ',' {
					nsonName = tag[:j]
					if len(tag) > j+1 && tag[j+1:] == "omitempty" {
						omitEmpty = true
					}
					break
				}
			}
		}

		if nsonName == "" {
			nsonName = field.Name
		}

		*fields = append(*fields, fieldInfo{
			indices:   indices,
			name:      field.Name,
			nsonName:  nsonName,
			typ:       field.Type,
			omitEmpty: omitEmpty,
		})
	}
}
