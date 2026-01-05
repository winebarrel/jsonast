package jsonast

var (
	nullValue = &JsonValue{Null: &JsonNull{}}
)

func (v *JsonValue) UnionType(other *JsonValue) *JsonValue {
	return v.Value().UnionType(other)
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
