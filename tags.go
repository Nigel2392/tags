package tags

import (
	"reflect"
	"strings"
)

// A map of tags, where each tag can have multiple values
//
// This would be the equivalent of the following struct:
//
//	type TagMap struct {
//		Tagfield []string `tag:"tag1=value1,value2,value3;tag2=value1,value2;"`
//	}
//
// You will get the KV pairs from the struct by calling:
//
//	var m = FromStruct(TagMap{}, "tag", "=", ";", ",")
//
// You can also parse a string directly:
//
//	var m = ParseTags("tag1=value1,value2,value3;tag2=value1,value2;")
type TagMap map[string][]string

// Returns a map of fieldnames, to their respective TagMaps.
func FromStruct(s any, tag string, delimKV, delimK, delimV string) map[string]TagMap {
	var m = make(map[string]TagMap)
	var t = reflect.TypeOf(s)
	var n = t.NumField()
	for i := 0; i < n; i++ {
		var f = t.Field(i)
		if !f.IsExported() || f.Anonymous {
			continue
		}
		var t = f.Tag.Get(tag)
		if t == "" || t == "-" {
			continue
		}
		m[f.Name] = ParseWithDelimiter(t, delimKV, delimK, delimV)
	}
	return m
}

func ParseTags(tag string) TagMap {
	// KEY=VALUE, VALUE, VALUE; KEY=VALUE, VALUE, VALUE; KEY=VALUE, VALUE, VALUE
	return ParseWithDelimiter(tag, ";", "=", ",")
}

func ParseWithDelimiter(tag string, delimiterKV string, delimiterK string, delimiterV string) TagMap {
	var m = make(TagMap)
	var parts = strings.Split(tag, delimiterKV)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		var kv = strings.Split(part, delimiterK)
		if len(kv) == 2 {
			var v = strings.Split(kv[1], delimiterV)
			m[kv[0]] = v
		} else if len(kv) == 1 {
			m[kv[0]] = []string{}
		}
	}
	return m
}

func (t TagMap) GetOK(key string) ([]string, bool) {
	var v, ok = t[key]
	if !ok {
		v = make([]string, 0)
	}
	return v, ok
}

func (t TagMap) GetSingle(key string, def ...string) string {
	var v, ok = t[key]
	if !ok || len(v) == 0 {
		if len(def) > 0 {
			return def[0]
		}
		return ""
	}
	return v[0]
}

func (t TagMap) Exists(key string) bool {
	_, ok := t[key]
	return ok
}
