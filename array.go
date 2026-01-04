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

	a := make([]*JsonFalse, 0, len(v.Elements))

	for _, e := range v.Elements {
		a = append(a, e.False)
	}

	return a
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

	a := make([]*JsonNull, 0, len(v.Elements))

	for _, e := range v.Elements {
		a = append(a, e.Null)
	}

	return a
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

	a := make([]*JsonTrue, 0, len(v.Elements))

	for _, e := range v.Elements {
		a = append(a, e.True)
	}

	return a
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

	a := make([]*JsonObject, 0, len(v.Elements))

	for _, e := range v.Elements {
		a = append(a, e.Object)
	}

	return a
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

	a := make([]*JsonArray, 0, len(v.Elements))

	for _, e := range v.Elements {
		a = append(a, e.Array)
	}

	return a
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

	a := make([]*JsonNumber, 0, len(v.Elements))

	for _, e := range v.Elements {
		a = append(a, e.Number)
	}

	return a
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

	a := make([]*JsonString, 0, len(v.Elements))

	for _, e := range v.Elements {
		a = append(a, e.String)
	}

	return a
}
