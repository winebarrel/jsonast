package jsonast

func NullStr(v *JsonString) {
	v.nullable = true
}

func NullNum(v *JsonNumber) {
	v.nullable = true
}

func NullTrue(v *JsonTrue) {
	v.nullable = true
}

func NullFalse(v *JsonFalse) {
	v.nullable = true
}
