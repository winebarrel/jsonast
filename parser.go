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

type JsonFalse string
type JsonNull string
type JsonTrue string
type JsonNumber string

func (v *JsonNumber) String() string {
	return string(*v)
}

type JsonString string

func (v *JsonString) String() string {
	return string(*v)
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

func (v *JsonValue) Value() any {
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

func (v *JsonValue) IsPrimitive() bool {
	return v.False != nil ||
		v.Null != nil ||
		v.True != nil ||
		v.Number != nil ||
		v.String != nil
}

type JsonObject struct {
	Members       []*JsonObjectMember `parser:"'{' @@* '}'"`
	OmittableKeys map[string]struct{}
}

type JsonObjectMember struct {
	Key   string     `parser:"@string"`
	Value *JsonValue `parser:"@@"`
}

type JsonArray struct {
	Elements []*JsonValue `parser:"'[' @@* ']'"`
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
