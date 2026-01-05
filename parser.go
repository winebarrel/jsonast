package jsonast

import (
	"io"

	"github.com/alecthomas/participle/v2"
)

var (
	jsonParser = participle.MustBuild[JsonValue](
		participle.Lexer(&JsonDefinition{}),
	)
)

type JsonFalse struct {
	nullable
}

func (*JsonFalse) UnmarshalText([]byte) error { return nil }

type JsonNull struct {
	notnullable
}

func (*JsonNull) UnmarshalText([]byte) error { return nil }

type JsonTrue struct {
	nullable
}

func (*JsonTrue) UnmarshalText([]byte) error { return nil }

type JsonNumber struct {
	nullable
	Text string
}

func (v *JsonNumber) UnmarshalText(text []byte) error {
	v.Text = string(text)
	return nil
}

type JsonString struct {
	nullable
	Text string
}

func (v *JsonString) UnmarshalText(text []byte) error {
	v.Text = string(text)
	return nil
}

type ValueType interface {
	UnionType(*JsonValue) *JsonValue
	Nullable() bool
}

type JsonValue struct {
	False  *JsonFalse  `parser:"@false |"`
	Null   *JsonNull   `parser:"@null |"`
	True   *JsonTrue   `parser:"@true |"`
	Object *JsonObject `parser:"@@ |"`
	Array  *JsonArray  `parser:"@@ |"`
	Number *JsonNumber `parser:"@number |"`
	String *JsonString `parser:"@string"`
}

func (v *JsonValue) Value() ValueType {
	if v.False != nil {
		return v.False
	} else if v.Null != nil {
		return v.Null
	} else if v.True != nil {
		return v.True
	} else if v.Object != nil {
		return v.Object
	} else if v.Array != nil {
		return v.Array
	} else if v.Number != nil {
		return v.Number
	} else if v.String != nil {
		return v.String
	} else {
		return nil
	}
}

func (v *JsonValue) SameTypeAs(other *JsonValue) bool {
	return v.IsFalse() && other.IsFalse() ||
		v.IsNull() && other.IsNull() ||
		v.IsTrue() && other.IsTrue() ||
		v.IsObject() && other.IsObject() ||
		v.IsArray() && other.IsArray() ||
		v.IsNumber() && other.IsNumber() ||
		v.IsString() && other.IsString()
}

func (v *JsonValue) IsFalse() bool {
	return v.False != nil
}

func (v *JsonValue) IsNull() bool {
	return v.Null != nil
}

func (v *JsonValue) IsTrue() bool {
	return v.True != nil
}

func (v *JsonValue) IsObject() bool {
	return v.Object != nil
}

func (v *JsonValue) IsArray() bool {
	return v.Array != nil
}

func (v *JsonValue) IsNumber() bool {
	return v.Number != nil
}

func (v *JsonValue) IsString() bool {
	return v.String != nil
}

type JsonObject struct {
	notnullable
	Members       []*JsonObjectMember `parser:"'{' @@* '}'"`
	OmittableKeys map[string]struct{}
}

type JsonObjectMember struct {
	Key   string     `parser:"@string"`
	Value *JsonValue `parser:"@@"`
}

type JsonArray struct {
	notnullable
	Elements []*JsonValue `parser:"'[' @@* ']'"`
}

func (v *JsonArray) Len() int {
	return len(v.Elements)
}

func ParseBytes(filename string, src []byte) (*JsonValue, error) {
	v, err := jsonParser.ParseBytes(filename, src)

	if err != nil {
		return nil, err
	}

	return v, nil
}

func Parse(filename string, r io.Reader) (*JsonValue, error) {
	v, err := jsonParser.Parse(filename, r)

	if err != nil {
		return nil, err
	}

	return v, nil
}
