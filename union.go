package jsonast

import "fmt"

func (v *JsonValue) UnionType(other *JsonValue) *JsonValue {
	return v.Value().UnionType(other)
}

func (v *JsonTrue) UnionType(other *JsonValue) *JsonValue {
	if other.IsTrue() || other.IsFalse() || other.IsNull() {
		newval := &JsonTrue{}
		newval.nullable = v.Or(other.IsNull() || other.Nullable())
		return &JsonValue{True: newval}
	} else {
		return &JsonValue{Null: &JsonNull{any: true}}
	}
}

func (v *JsonFalse) UnionType(other *JsonValue) *JsonValue {
	if other.IsTrue() || other.IsFalse() || other.IsNull() {
		newval := &JsonFalse{}
		newval.nullable = v.Or(other.IsNull() || other.Nullable())
		return &JsonValue{False: newval}
	} else {
		return &JsonValue{Null: &JsonNull{any: true}}
	}
}

func (v *JsonNull) UnionType(other *JsonValue) *JsonValue {
	switch o := other.Value().(type) {
	case *JsonFalse:
		newval := &JsonFalse{}
		newval.nullable = true
		return &JsonValue{False: newval}
	case *JsonTrue:
		newval := &JsonTrue{}
		newval.nullable = true
		return &JsonValue{True: newval}
	case *JsonObject:
		newval := &JsonObject{Members: o.Members}
		return &JsonValue{Object: newval}
	case *JsonArray:
		newval := &JsonArray{Elements: o.Elements}
		return &JsonValue{Array: newval}
	case *JsonNumber:
		newval := &JsonNumber{Text: o.Text}
		newval.nullable = true
		return &JsonValue{Number: newval}
	case *JsonString:
		newval := &JsonString{Text: o.Text}
		newval.nullable = true
		return &JsonValue{String: newval}
	case *JsonNull:
		newval := &JsonNull{any: v.any || o.any}
		return &JsonValue{Null: newval}
	default:
		panic(fmt.Sprintf("unexpected type: %+v", o))
	}
}

func (v *JsonNumber) UnionType(other *JsonValue) *JsonValue {
	if other.IsNumber() || other.IsNull() {
		newval := &JsonNumber{Text: v.Text}
		newval.nullable = v.Or(other.IsNull() || other.Nullable())
		return &JsonValue{Number: newval}
	} else {
		return &JsonValue{Null: &JsonNull{any: true}}
	}
}

func (v *JsonString) UnionType(other *JsonValue) *JsonValue {
	if other.IsString() || other.IsNull() {
		newval := &JsonString{Text: v.Text}
		newval.nullable = v.Or(other.IsNull() || other.Nullable())
		return &JsonValue{String: newval}
	} else {
		return &JsonValue{Null: &JsonNull{any: true}}
	}
}

func (v *JsonArray) UnionType(other *JsonValue) *JsonValue {
	if other != nil {
		if other.IsNull() {
			newval := &JsonArray{Elements: v.Elements}
			return &JsonValue{Array: newval}
		} else if !other.IsArray() {
			return &JsonValue{Null: &JsonNull{any: true}}
		}
	}

	if other == nil {
		other = &JsonValue{Array: &JsonArray{Elements: []*JsonValue{}}}
	}

	if len(v.Elements) == 0 && len(other.Array.Elements) == 0 {
		return &JsonValue{Array: &JsonArray{Elements: []*JsonValue{}}}
	}

	elems := make([]*JsonValue, 0, len(v.Elements)+len(other.Array.Elements))
	elems = append(elems, v.Elements...)
	elems = append(elems, other.Array.Elements...)
	union := elems[0]
	elems = elems[1:]

	for _, e := range elems {
		if union.IsNull() && union.Null.any {
			break
		}

		union = union.UnionType(e)
	}

	return &JsonValue{
		Array: &JsonArray{
			Elements: []*JsonValue{union},
		},
	}
}

func (v *JsonObject) UnionType(other *JsonValue) *JsonValue {
	if other.IsNull() {
		newval := &JsonObject{
			Members:       v.Members,
			OmittableKeys: v.OmittableKeys,
		}
		return &JsonValue{Object: newval}
	} else if !other.IsObject() {
		return &JsonValue{Null: &JsonNull{any: true}}
	}

	type entry struct {
		member *JsonObjectMember
		order  int
		keycnt int
	}

	members := make(map[string]*entry, len(v.Members)+len(other.Object.Members))

	for i, m := range v.Members {
		members[m.Key] = &entry{member: m, order: i, keycnt: 1}
	}

	midx := len(members)

	for _, omem := range other.Object.Members {
		k := omem.Key
		var ent *entry

		if e, ok := members[k]; ok {
			union := e.member.Value.UnionType(omem.Value)
			e.member = &JsonObjectMember{Key: k, Value: union}
			e.keycnt += 1
			ent = e
		} else {
			ent = &entry{
				member: omem,
				order:  midx,
				keycnt: 1,
			}
			midx++
		}

		members[k] = ent
	}

	entries := make([]*JsonObjectMember, len(members))
	omittableKeys := map[string]struct{}{}

	for k, e := range members {
		entries[e.order] = e.member

		if e.keycnt == 1 {
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
