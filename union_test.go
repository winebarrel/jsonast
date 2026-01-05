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
			value:    vstr("s"),
			other:    &jsonast.JsonValue{String: vstr("s")},
			expected: &jsonast.JsonValue{String: vstr("s")},
		},
		{
			name:     "string <=> number",
			value:    vstr("s"),
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "string <=> true",
			value:    vstr("s"),
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "string <=> false",
			value:    vstr("s"),
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "string <=> null",
			value:    vstr("s"),
			other:    &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "string <=> array",
			value:    vstr("s"),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "string <=> object",
			value:    vstr("s"),
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: vnull()},
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
			value:    vnum("1"),
			other:    &jsonast.JsonValue{String: vstr("s")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "number <=> number",
			value:    vnum("1"),
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Number: vnum("1")},
		},
		{
			name:     "number <=> true",
			value:    vnum("1"),
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "number <=> false",
			value:    vnum("1"),
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "number <=> null",
			value:    vnum("1"),
			other:    &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "number <=> array",
			value:    vnum("1"),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "number <=> object",
			value:    vnum("1"),
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: vnull()},
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
			value:    vtrue(),
			other:    &jsonast.JsonValue{String: vstr("s")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "true <=> number",
			value:    vtrue(),
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "true <=> true",
			value:    vtrue(),
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{True: vtrue()},
		},
		{
			name:     "true <=> false",
			value:    vtrue(),
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{True: vtrue()},
		},
		{
			name:     "true <=> null",
			value:    vtrue(),
			other:    &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "true <=> array",
			value:    vtrue(),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "true <=> object",
			value:    vtrue(),
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: vnull()},
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
			value:    vfalse(),
			other:    &jsonast.JsonValue{String: vstr("s")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "false <=> number",
			value:    vfalse(),
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "false <=> true",
			value:    vfalse(),
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{False: vfalse()},
		},
		{
			name:     "false <=> false",
			value:    vfalse(),
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{False: vfalse()},
		},
		{
			name:     "false <=> null",
			value:    vfalse(),
			other:    &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "false <=> array",
			value:    vfalse(),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "false <=> object",
			value:    vfalse(),
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: vnull()},
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
			value:    vnull(),
			other:    &jsonast.JsonValue{String: vstr("s")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "null <=> number",
			value:    vnull(),
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "null <=> true",
			value:    vnull(),
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "null <=> false",
			value:    vnull(),
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "null <=> null",
			value:    vnull(),
			other:    &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "null <=> array",
			value:    vnull(),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "null <=> object",
			value:    vnull(),
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: vnull()},
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
				{String: vstr("s")},
			}},
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "array <=> false",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}},
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "array <=> string",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}},
			other:    &jsonast.JsonValue{String: vstr("s")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "array <=> number",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}},
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "array <=> null",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}},
			other:    &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "array <=> object",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}},
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{}},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "array <=> string array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s2")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}}},
		},
		{
			name: "array <=> number array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Number: vnum("1")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Number: vnum("2")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Number: vnum("1")},
			}}},
		},
		{
			name: "array <=> true array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{True: vtrue()},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{True: vtrue()},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{True: vtrue()},
			}}},
		},
		{
			name: "array <=> false array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{False: vfalse()},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{False: vfalse()},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{False: vfalse()},
			}}},
		},
		{
			name: "array <=> null array 1",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
			}}},
		},
		{
			name: "array <=> null array 2",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
			}}},
		},
		{
			name: "array <=> null array 3",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
			}}},
		},
		{
			name: "array <=> empty array 1",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}}},
		},
		{
			name:  "array <=> empty array 2",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}}},
		},
		{
			name:     "array <=> empty array 3",
			value:    &jsonast.JsonArray{Elements: []*jsonast.JsonValue{}},
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{}}},
		},
		{
			name:  "array <=> nil 1",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: vstr("s")}}},
			other: nil,
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}}},
		},
		{
			name:  "array <=> nil 2",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: vstr("s")}, {Number: vnum("1")}}},
			other: nil,
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
			}}},
		},
		{
			name:     "array <=> nil 3",
			value:    &jsonast.JsonArray{Elements: []*jsonast.JsonValue{}},
			other:    nil,
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{}}},
		},
		{
			name: "array <=> composite array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
				{Number: vnum("1")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s2")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
			}}},
		},
		{
			name: "array <=> nested array 1",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{String: vstr("s")},
				}}},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{String: vstr("s2")},
				}}},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{String: vstr("s")},
				}}},
			}}},
		},
		{
			name: "array <=> nested array 2",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{String: vstr("s")},
				}}},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{Number: vnum("1")},
				}}},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{Null: vnull()},
				}}},
			}}},
		},
		{
			name: "array <=> object array 1",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
				}}},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
				}}},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{
					Members: []*jsonast.JsonObjectMember{
						{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
						{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
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
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
			}},
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "object <=> false",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
			}},
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "object <=> string",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
			}},
			other:    &jsonast.JsonValue{String: vstr("s")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "object <=> number",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
			}},
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "object <=> null",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
			}},
			other:    &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "object <=> array",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
			}},
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name: "object <=> object 1",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s")}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
					{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s")}},
				},
				OmittableKeys: map[string]struct{}{"str": {}, "str2": {}},
			}},
		},
		{
			name: "object <=> object 2",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s2")}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
				},
				OmittableKeys: map[string]struct{}{},
			}},
		},
		{
			name: "object <=> object 3",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
				{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2'")}},
				{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
					{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
					{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
				},
				OmittableKeys: map[string]struct{}{"str": {}, "str3": {}},
			}},
		},
		{
			name: "object <=> object 4",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
				},
				OmittableKeys: map[string]struct{}{"str": {}},
			}},
		},
		{
			name:  "object <=> object 5",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s")}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s")}},
				},
				OmittableKeys: map[string]struct{}{"str2": {}},
			}},
		},
		{
			name: "object <=> array object 1",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: vstr("s")}}}}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: vstr("s2")}}}}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: vstr("s")}}}}},
				},
				OmittableKeys: map[string]struct{}{},
			}},
		},
		{
			name: "object <=> array object 2",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: vstr("s")}}}}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{Number: vnum("1")}}}}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "ary", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{Null: vnull()}}}}},
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
