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
			expected: &jsonast.JsonValue{String: pstr("s")},
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

func TestPtrStringUnionType(t *testing.T) {
	tests := []struct {
		name     string
		value    *jsonast.JsonString
		other    *jsonast.JsonValue
		expected *jsonast.JsonValue
	}{
		{
			name:     "ptr string <=> string",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{String: vstr("s")},
			expected: &jsonast.JsonValue{String: pstr("s")},
		},
		{
			name:     "ptr string <=> number",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr string <=> true",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr string <=> false",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr string <=> null",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{String: pstr("s")},
		},
		{
			name:     "ptr string <=> array",
			value:    pstr("s"),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr string <=> object",
			value:    pstr("s"),
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
			expected: &jsonast.JsonValue{Number: pnum("1")},
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

func TestPtrNumberUnionType(t *testing.T) {
	tests := []struct {
		name     string
		value    *jsonast.JsonNumber
		other    *jsonast.JsonValue
		expected *jsonast.JsonValue
	}{
		{
			name:     "ptr number <=> string",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{String: vstr("s")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr number <=> number",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Number: pnum("1")},
		},
		{
			name:     "ptr number <=> true",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr number <=> false",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr number <=> null",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{Number: pnum("1")},
		},
		{
			name:     "ptr number <=> array",
			value:    pnum("1"),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr number <=> object",
			value:    pnum("1"),
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
			expected: &jsonast.JsonValue{True: ptrue()},
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

func TestPtrTrueUnionType(t *testing.T) {
	tests := []struct {
		name     string
		value    *jsonast.JsonTrue
		other    *jsonast.JsonValue
		expected *jsonast.JsonValue
	}{
		{
			name:     "ptr true <=> string",
			value:    ptrue(),
			other:    &jsonast.JsonValue{String: vstr("s")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr true <=> number",
			value:    ptrue(),
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr true <=> true",
			value:    ptrue(),
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{True: ptrue()},
		},
		{
			name:     "ptr true <=> false",
			value:    ptrue(),
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{True: ptrue()},
		},
		{
			name:     "ptr true <=> null",
			value:    ptrue(),
			other:    &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{True: ptrue()},
		},
		{
			name:     "ptr true <=> array",
			value:    ptrue(),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr true <=> object",
			value:    ptrue(),
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
			expected: &jsonast.JsonValue{False: pfalse()},
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

func TestPtrFalseUnionType(t *testing.T) {
	tests := []struct {
		name     string
		value    *jsonast.JsonFalse
		other    *jsonast.JsonValue
		expected *jsonast.JsonValue
	}{
		{
			name:     "ptr false <=> string",
			value:    pfalse(),
			other:    &jsonast.JsonValue{String: vstr("s")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr false <=> number",
			value:    pfalse(),
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr false <=> true",
			value:    pfalse(),
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{False: pfalse()},
		},
		{
			name:     "ptr false <=> false",
			value:    pfalse(),
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{False: pfalse()},
		},
		{
			name:     "ptr false <=> null",
			value:    pfalse(),
			other:    &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{False: pfalse()},
		},
		{
			name:     "ptr false <=> array",
			value:    pfalse(),
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{}},
			expected: &jsonast.JsonValue{Null: vnull()},
		},
		{
			name:     "ptr false <=> object",
			value:    pfalse(),
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
			expected: &jsonast.JsonValue{String: pstr("s")},
		},
		{
			name:     "null <=> number",
			value:    vnull(),
			other:    &jsonast.JsonValue{Number: vnum("1")},
			expected: &jsonast.JsonValue{Number: pnum("1")},
		},
		{
			name:     "null <=> true",
			value:    vnull(),
			other:    &jsonast.JsonValue{True: vtrue()},
			expected: &jsonast.JsonValue{True: ptrue()},
		},
		{
			name:     "null <=> false",
			value:    vnull(),
			other:    &jsonast.JsonValue{False: vfalse()},
			expected: &jsonast.JsonValue{False: pfalse()},
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
			other:    &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: vstr("s")}}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: vstr("s")}}}},
		},
		{
			name:     "null <=> object",
			value:    vnull(),
			other:    &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{{Key: "n", Value: &jsonast.JsonValue{Number: vnum("1")}}}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{{Key: "n", Value: &jsonast.JsonValue{Number: vnum("1")}}}}},
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
			other: &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
			}}},
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
			name: "array <=> ptr string array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
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
			name: "array <=> ptr number array",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Number: pnum("1")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Number: vnum("2")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Number: pnum("1")},
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
				{String: pstr("s")},
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
				{String: pstr("s")},
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
			name:  "array <=> nil 4",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{{String: vstr("s")}, {Null: vnull()}}},
			other: nil,
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}}},
		},
		{
			name: "array <=> composite array 1",
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
			name: "array <=> composite array 2",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s")},
				{String: pstr("ps")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s2")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("s")},
			}}},
		},
		{
			name: "array <=> composite array 3",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
				{String: pstr("ps")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s2")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("ps")},
			}}},
		},
		{
			name: "array <=> composite array 4",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
				{Null: vnull()},
				{String: pstr("ps")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s2")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: pstr("ps")},
			}}},
		},
		{
			name: "array <=> composite array 5",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
				{Null: vnull()},
				{String: pstr("ps")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{String: vstr("s2")},
				{Number: vnum("1")},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
			}}},
		},
		{
			name: "array <=> composite array 6",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Null: vnull()},
				{Null: vnull()},
				{String: pstr("ps")},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Number: vnum("1")},
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
			name: "array <=> nested array 3",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{String: vstr("s")},
				}}},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{Null: vnull()},
				}}},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
					{String: pstr("s")},
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
		{
			name: "array <=> object array 2",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
					{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
				}}},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
					{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
				}}},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{
					Members: []*jsonast.JsonObjectMember{
						{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
						{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
						{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
					},
					OmittableKeys: map[string]struct{}{"str": {}, "str3": {}},
				}},
			}}},
		},
		{
			name: "array <=> object array 3",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
					{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
				}}},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s2")}},
					{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
				}}},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{
					Members: []*jsonast.JsonObjectMember{
						{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
						{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s2")}},
						{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
					},
					OmittableKeys: map[string]struct{}{"str": {}, "str3": {}},
				}},
			}}},
		},
		{
			name: "array <=> object array 4",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
					{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
				}}},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str2", Value: &jsonast.JsonValue{Number: vnum("1")}},
					{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
				}}},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{
					Members: []*jsonast.JsonObjectMember{
						{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
						{Key: "str2", Value: &jsonast.JsonValue{Null: vnull()}},
						{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
					},
					OmittableKeys: map[string]struct{}{"str": {}, "str3": {}},
				}},
			}}},
		},
		{
			name: "array <=> object array 5",
			value: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
					{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
				}}},
			}},
			other: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
					{Key: "str2", Value: &jsonast.JsonValue{String: pstr("ps2")}},
					{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
				}}},
			}}},
			expected: &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: []*jsonast.JsonValue{
				{Object: &jsonast.JsonObject{
					Members: []*jsonast.JsonObjectMember{
						{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
						{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s2")}},
						{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
					},
					OmittableKeys: map[string]struct{}{"str": {}, "str3": {}},
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
			other: &jsonast.JsonValue{Null: vnull()},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
			}}},
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
			name: "object <=> object 6",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
				{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s2'")}},
				{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
					{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s2")}},
					{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
				},
				OmittableKeys: map[string]struct{}{"str": {}, "str3": {}},
			}},
		},
		{
			name: "object <=> object 7",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
				{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str2", Value: &jsonast.JsonValue{Null: vnull()}},
				{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
					{Key: "str2", Value: &jsonast.JsonValue{String: pstr("s2")}},
					{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
				},
				OmittableKeys: map[string]struct{}{"str": {}, "str3": {}},
			}},
		},
		{
			name: "object <=> object 8",
			value: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
				{Key: "str2", Value: &jsonast.JsonValue{String: vstr("s2")}},
			}},
			other: &jsonast.JsonValue{Object: &jsonast.JsonObject{Members: []*jsonast.JsonObjectMember{
				{Key: "str2", Value: &jsonast.JsonValue{Number: vnum("1")}},
				{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
			}}},
			expected: &jsonast.JsonValue{Object: &jsonast.JsonObject{
				Members: []*jsonast.JsonObjectMember{
					{Key: "str", Value: &jsonast.JsonValue{String: vstr("s")}},
					{Key: "str2", Value: &jsonast.JsonValue{Null: vnull()}},
					{Key: "str3", Value: &jsonast.JsonValue{String: vstr("s3")}},
				},
				OmittableKeys: map[string]struct{}{"str": {}, "str3": {}},
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
