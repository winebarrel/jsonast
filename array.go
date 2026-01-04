package jsonast

func (v *JsonArray) Len() int {
	return len(v.Elements)
}

func (v *JsonArray) IsFalseArray() bool {
	if len(v.Elements) == 0 {
		return false
	}

	for _, e := range v.Elements {
		if !e.IsFalse() {
			return false
		}
	}

	return true
}

func (v *JsonArray) FalseArray() []*JsonFalse {
	if !v.IsFalseArray() {
		return nil
	}

	objs := make([]*JsonFalse, 0, len(v.Elements))

	for _, e := range v.Elements {
		objs = append(objs, e.False)
	}

	return objs
}

func (v *JsonArray) IsNullArray() bool {
	if len(v.Elements) == 0 {
		return false
	}

	for _, e := range v.Elements {
		if !e.IsNull() {
			return false
		}
	}

	return true
}

func (v *JsonArray) NullArray() []*JsonNull {
	if !v.IsNullArray() {
		return nil
	}

	objs := make([]*JsonNull, 0, len(v.Elements))

	for _, e := range v.Elements {
		objs = append(objs, e.Null)
	}

	return objs
}

func (v *JsonArray) IsTrueArray() bool {
	if len(v.Elements) == 0 {
		return false
	}

	for _, e := range v.Elements {
		if !e.IsTrue() {
			return false
		}
	}

	return true
}

func (v *JsonArray) TrueArray() []*JsonTrue {
	if !v.IsTrueArray() {
		return nil
	}

	objs := make([]*JsonTrue, 0, len(v.Elements))

	for _, e := range v.Elements {
		objs = append(objs, e.True)
	}

	return objs
}

func (v *JsonArray) IsObjectArray() bool {
	if len(v.Elements) == 0 {
		return false
	}

	for _, e := range v.Elements {
		if !e.IsObject() {
			return false
		}
	}

	return true
}

func (v *JsonArray) ObjectArray() []*JsonObject {
	if !v.IsObjectArray() {
		return nil
	}

	objs := make([]*JsonObject, 0, len(v.Elements))

	for _, e := range v.Elements {
		objs = append(objs, e.Object)
	}

	return objs
}

func (v *JsonArray) IsArrayArray() bool {
	if len(v.Elements) == 0 {
		return false
	}

	for _, e := range v.Elements {
		if !e.IsArray() {
			return false
		}
	}

	return true
}

func (v *JsonArray) ArrayArray() []*JsonArray {
	if !v.IsArrayArray() {
		return nil
	}

	objs := make([]*JsonArray, 0, len(v.Elements))

	for _, e := range v.Elements {
		objs = append(objs, e.Array)
	}

	return objs
}

func (v *JsonArray) IsNumberArray() bool {
	if len(v.Elements) == 0 {
		return false
	}

	for _, e := range v.Elements {
		if !e.IsNumber() {
			return false
		}
	}

	return true
}

func (v *JsonArray) NumberArray() []*JsonNumber {
	if !v.IsNumberArray() {
		return nil
	}

	objs := make([]*JsonNumber, 0, len(v.Elements))

	for _, e := range v.Elements {
		objs = append(objs, e.Number)
	}

	return objs
}

func (v *JsonArray) IsStringArray() bool {
	if len(v.Elements) == 0 {
		return false
	}

	for _, e := range v.Elements {
		if !e.IsString() {
			return false
		}
	}

	return true
}

func (v *JsonArray) StringArray() []*JsonString {
	if !v.IsStringArray() {
		return nil
	}

	objs := make([]*JsonString, 0, len(v.Elements))

	for _, e := range v.Elements {
		objs = append(objs, e.String)
	}

	return objs
}
