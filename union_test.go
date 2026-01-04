package jsonast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/jsonast"
)

func TestStringUnionType(t *testing.T) {
	tests := []struct {
		name     string
		value    *jsonast.JsonString
		other    *jsonast.JsonValue
		expected *jsonast.JsonValue
	}{
		{
			name:     "string <=> string",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{String: pstr("s")},
			expected: &jsonast.JsonValue{String: pstr("s")},
		},
		{
			name:     "string <=> number",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{Number: pnum("1")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "string <=> true",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{True: ptrue("true")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "string <=> false",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{False: pfalse("false")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "string <=> null",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{Null: pnull("null")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "string <=> array",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "string <=> object",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			union := tt.value.UnionType(tt.other)
			assert.Equal(t, tt.expected, union)
		})
	}
}

func TestNumberUnionType(t *testing.T) {
	tests := []struct {
		name     string
		value    *jsonast.JsonNumber
		other    *jsonast.JsonValue
		expected *jsonast.JsonValue
	}{
		{
			name:     "number <=> string",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{String: pstr("s")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "number <=> number",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{Number: pnum("1")},
			expected: &jsonast.JsonValue{Number: pnum("1")},
		},
		{
			name:     "number <=> true",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{True: ptrue("true")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "number <=> false",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{False: pfalse("false")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "number <=> null",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{Null: pnull("null")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "number <=> array",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "number <=> object",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			union := tt.value.UnionType(tt.other)
			assert.Equal(t, tt.expected, union)
		})
	}
}

func TestTrueUnionType(t *testing.T) {
	tests := []struct {
		name     string
		value    *jsonast.JsonTrue
		other    *jsonast.JsonValue
		expected *jsonast.JsonValue
	}{
		{
			name:     "true <=> string",
			value:    ptrue("true"),
			other:    &jsonast.JsonValue{String: pstr("s")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "true <=> number",
			value:    ptrue("true"),
			other:    &jsonast.JsonValue{Number: pnum("1")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "true <=> true",
			value:    ptrue("true"),
			other:    &jsonast.JsonValue{True: ptrue("true")},
			expected: &jsonast.JsonValue{True: ptrue("true")},
		},
		{
			name:     "true <=> false",
			value:    ptrue("true"),
			other:    &jsonast.JsonValue{False: pfalse("false")},
			expected: &jsonast.JsonValue{True: ptrue("true")},
		},
		{
			name:     "true <=> null",
			value:    ptrue("true"),
			other:    &jsonast.JsonValue{Null: pnull("null")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "true <=> array",
			value:    ptrue("true"),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "true <=> object",
			value:    ptrue("true"),
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			union := tt.value.UnionType(tt.other)
			assert.Equal(t, tt.expected, union)
		})
	}
}

func TestFalseUnionType(t *testing.T) {
	tests := []struct {
		name     string
		value    *jsonast.JsonFalse
		other    *jsonast.JsonValue
		expected *jsonast.JsonValue
	}{
		{
			name:     "false <=> string",
			value:    pfalse("false"),
			other:    &jsonast.JsonValue{String: pstr("s")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "false <=> number",
			value:    pfalse("false"),
			other:    &jsonast.JsonValue{Number: pnum("1")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "false <=> true",
			value:    pfalse("false"),
			other:    &jsonast.JsonValue{True: ptrue("true")},
			expected: &jsonast.JsonValue{False: pfalse("false")},
		},
		{
			name:     "false <=> false",
			value:    pfalse("false"),
			other:    &jsonast.JsonValue{False: pfalse("false")},
			expected: &jsonast.JsonValue{False: pfalse("false")},
		},
		{
			name:     "false <=> null",
			value:    pfalse("false"),
			other:    &jsonast.JsonValue{Null: pnull("null")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "false <=> array",
			value:    pfalse("false"),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "false <=> object",
			value:    pfalse("false"),
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			union := tt.value.UnionType(tt.other)
			assert.Equal(t, tt.expected, union)
		})
	}
}

func TestNullUnionType(t *testing.T) {
	tests := []struct {
		name     string
		value    *jsonast.JsonNull
		other    *jsonast.JsonValue
		expected *jsonast.JsonValue
	}{
		{
			name:     "null <=> string",
			value:    pnull("null"),
			other:    &jsonast.JsonValue{String: pstr("s")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "null <=> number",
			value:    pnull("null"),
			other:    &jsonast.JsonValue{Number: pnum("1")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "null <=> true",
			value:    pnull("null"),
			other:    &jsonast.JsonValue{True: ptrue("true")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "null <=> false",
			value:    pnull("null"),
			other:    &jsonast.JsonValue{False: pfalse("false")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "null <=> null",
			value:    pnull("null"),
			other:    &jsonast.JsonValue{Null: pnull("null")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "null <=> array",
			value:    pnull("null"),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name:     "null <=> object",
			value:    pnull("null"),
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			union := tt.value.UnionType(tt.other)
			assert.Equal(t, tt.expected, union)
		})
	}
}

func TestArrayUnionType(t *testing.T) {
	tests := []struct {
		name     string
		value    *jsonast.JsonArray
		other    *jsonast.JsonValue
		expected *jsonast.JsonValue
	}{
		{
			name: "array <=> true",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}},
			other:    &jsonast.JsonValue{True: ptrue("true")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "array <=> false",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}},
			other:    &jsonast.JsonValue{False: pfalse("false")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "array <=> string",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}},
			other:    &jsonast.JsonValue{String: pstr("s")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "array <=> number",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}},
			other:    &jsonast.JsonValue{Number: pnum("1")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "array <=> null",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}},
			other:    &jsonast.JsonValue{Null: pnull("null")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "array <=> object",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}},
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "array <=> string array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s2")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}}},
		},
		{
			name: "array <=> number array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Number: pnum("1")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Number: pnum("2")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Number: pnum("1")},
			}}},
		},
		{
			name: "array <=> true array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{True: ptrue("true")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{True: ptrue("true")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{True: ptrue("true")},
			}}},
		},
		{
			name: "array <=> false array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{False: pfalse("false")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{False: pfalse("false")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{False: pfalse("false")},
			}}},
		},
		{
			name: "array <=> null array 1",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: pnull("null")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: pnull("null")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: pnull("null")},
			}}},
		},
		{
			name: "array <=> null array 2",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: pnull("null")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: pnull("null")},
			}}},
		},
		{
			name: "array <=> null array 3",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: pnull("null")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: pnull("null")},
			}}},
		},
		{
			name: "array <=> empty array 1",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}},
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{}}},
		},
		{
			name:  "array <=> empty array 2",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{}}},
		},
		{
			name: "array <=> composite array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
				{Number: pnum("1")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s2")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: pnull("null")},
			}}},
		},
		{
			name: "array <=> nested array 1",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{String: pstr("s")},
				}}},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{String: pstr("s2")},
				}}},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{String: pstr("s")},
				}}},
			}}},
		},
		{
			name: "array <=> nested array 2",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{String: pstr("s")},
				}}},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{Number: pnum("1")},
				}}},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{Null: pnull("null")},
				}}},
			}}},
		},
		{
			name: "array <=> object array 1",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
				}}},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s2")}},
				}}},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{
					Members: []*jsonast.JsonObjectMember{
						{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
						{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s2")}},
					},
					OmittableKeys: map[string]struct{}{"str": {}, "str2": {}},
				}},
			}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			union := tt.value.UnionType(tt.other)
			assert.Equal(t, tt.expected, union)
		})
	}
}

func TestObjectUnionType(t *testing.T) {
	tests := []struct {
		name     string
		value    *jsonast.JsonObject
		other    *jsonast.JsonValue
		expected *jsonast.JsonValue
	}{
		{
			name: "object <=> true",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
			}},
			other:    &jsonast.JsonValue{True: ptrue("true")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "object <=> false",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
			}},
			other:    &jsonast.JsonValue{False: pfalse("false")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "object <=> string",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
			}},
			other:    &jsonast.JsonValue{String: pstr("s")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "object <=> number",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
			}},
			other:    &jsonast.JsonValue{Number: pnum("1")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "object <=> null",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
			}},
			other:    &jsonast.JsonValue{Null: pnull("null")},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "object <=> array",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
			}},
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: pnull("null")},
		},
		{
			name: "object <=> object 1",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s")}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
					{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s")}},
				},
				OmittableKeys: map[string]struct{}{"str": {}, "str2": {}},
			}},
		},
		{
			name: "object <=> object 2",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: pstr("s2")}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
				},
				OmittableKeys: map[string]struct{}{},
			}},
		},
		{
			name: "object <=> object 3",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
				{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s2")}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s2'")}},
				{Key: "str3", Value: &jsonast.JsonValue{String: pstr("s3")}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
					{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s2")}},
					{Key: "str3", Value: &jsonast.JsonValue{String: pstr("s3")}},
				},
				OmittableKeys: map[string]struct{}{"str": {}, "str3": {}},
			}},
		},
		{
			name: "object <=> object 4",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: pstr("s")}},
				},
				OmittableKeys: map[string]struct{}{"str": {}},
			}},
		},
		{
			name:  "object <=> object 5",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s")}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s")}},
				},
				OmittableKeys: map[string]struct{}{"str2": {}},
			}},
		},
		{
			name: "object <=> array object 1",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: pstr("s")}}}}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: pstr("s2")}}}}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: pstr("s")}}}}},
				},
				OmittableKeys: map[string]struct{}{},
			}},
		},
		{
			name: "object <=> array object 2",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: pstr("s")}}}}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{Number: pnum("1")}}}}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{Null: pnull("null")}}}}},
				},
				OmittableKeys: map[string]struct{}{},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			union := tt.value.UnionType(tt.other)
			assert.Equal(t, tt.expected, union)
		})
	}
}
