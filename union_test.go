package jsonast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/jsonast"
)

// TODO: other tests

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			union := tt.value.UnionType(tt.other)
			assert.Equal(t, tt.expected, union)
		})
	}
}
