package jsonast

type nullable bool

func (v nullable) Nullable() bool {
	return bool(v)
}

func (v nullable) Or(other bool) nullable {
	return nullable(bool(v) || other)
}

func (v *JsonValue) Nullable() bool {
	return v.Value().Nullable()
}

type notnullable struct{}

func (v notnullable) Nullable() bool {
	return false
}
