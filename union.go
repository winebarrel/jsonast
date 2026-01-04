package jsonast

import (
	"github.com/kazamori/orderedmap"
)

var (
	nullValue = func() *JsonValue {
		null := JsonNull("null")
		return &JsonValue{Null: &null}
	}()
)

func (v *JsonValue) UnionType(other *JsonValue) *JsonValue {
	value := v.Value().(interface{ UnionType(*JsonValue) *JsonValue })
	return value.UnionType(other)
}

func (v *JsonTrue) UnionType(other *JsonValue) *JsonValue {
	if other.IsTrue() || other.IsFalse() {
		return &JsonValue{True: v}
	} else {
		return nullValue
	}
}

func (v *JsonFalse) UnionType(other *JsonValue) *JsonValue {
	if other.IsTrue() || other.IsFalse() {
		return &JsonValue{False: v}
	} else {
		return nullValue
	}
}

func (v *JsonNull) UnionType(other *JsonValue) *JsonValue {
	return &JsonValue{Null: v}
}

func (v *JsonNumber) UnionType(other *JsonValue) *JsonValue {
	if other.IsNumber() {
		return &JsonValue{Number: v}
	} else {
		return nullValue
	}
}

func (v *JsonString) UnionType(other *JsonValue) *JsonValue {
	if other.IsString() {
		return &JsonValue{String: v}
	} else {
		return nullValue
	}
}

func (v *JsonArray) UnionType(other *JsonValue) *JsonValue {
	if other != nil && !other.IsArray() {
		return nullValue
	}

	if len(v.Elements) == 0 || other != nil && len(other.Array.Elements) == 0 {
		return &JsonValue{Array: &JsonArray{Elements: []*JsonValue{}}}
	}

	if other == nil {
		other = &JsonValue{Array: &JsonArray{Elements: []*JsonValue{}}}
	}

	elems := make([]*JsonValue, 0, len(v.Elements)+len(other.Array.Elements))
	elems = append(elems, v.Elements...)
	elems = append(elems, other.Array.Elements...)
	union := elems[0]
	elems = elems[1:]

	for _, e := range elems {
		if union.IsNull() {
			break
		}

		union = union.UnionType(e)
	}

	return &JsonValue{Array: &JsonArray{Elements: []*JsonValue{union}}}
}

func (v *JsonObject) UnionType(other *JsonValue) *JsonValue {
	if !other.IsObject() {
		return nullValue
	}

	members := orderedmap.New[string, *JsonObjectMember]()
	keys := map[string]int{}

	for _, m := range v.Members {
		members.Set(m.Key, m)
		keys[m.Key] += 1
	}

	for _, omem := range other.Object.Members {
		k := omem.Key
		var newm *JsonObjectMember
		vmem, ok := members.Get(k)
		keys[k] += 1

		if ok {
			union := vmem.Value.UnionType(omem.Value)
			newm = &JsonObjectMember{Key: k, Value: union}
		} else {
			newm = omem
		}

		members.Set(k, newm)
	}

	entries := []*JsonObjectMember{}

	for _, pair := range members.Pairs() {
		entries = append(entries, pair.Value)
	}

	omittableKeys := map[string]struct{}{}

	for k, n := range keys {
		if n == 1 {
			omittableKeys[k] = struct{}{}
		}
	}

	return &JsonValue{
		Object: &JsonObject{
			Members:       entries,
			OmittableKeys: omittableKeys,
		},
	}
}
