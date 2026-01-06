package jsonast

func MakeStrNullable(v *JsonString) {
	v.nullable = true
}

func MakeNumNullable(v *JsonNumber) {
	v.nullable = true
}

func MakeTrueNullable(v *JsonTrue) {
	v.nullable = true
}

func MakeFalseNullable(v *JsonFalse) {
	v.nullable = true
}

func MakeNullAny(v *JsonNull) {
	v.any = true
}
