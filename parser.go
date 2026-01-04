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

type JsonValue struct {
	False  *string     `parser:"@false |"`
	Null   *string     `parser:"@null |"`
	True   *string     `parser:"@true |"`
	Object *JsonObject `parser:"@@ |"`
	Array  *JsonArray  `parser:"@@ |"`
	Number *string     `parser:"@number |"`
	String *string     `parser:"@string"`
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

type JsonObject struct {
	Members []*JsonObjectMember `parser:"'{' @@* '}'"`
}

type JsonObjectMember struct {
	Key   string     `parser:"@string"`
	Value *JsonValue `parser:"@@"`
}

type JsonArray struct {
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
