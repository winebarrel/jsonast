package jsonast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/jsonast"
)

func TestArrayLen(t *testing.T) {
	tests := []struct {
		values   []string
		expected int
	}{
		{
			values:   []string{},
			expected: 0,
		},
		{
			values:   []string{"a"},
			expected: 1,
		},
		{
			values:   []string{"a", "b"},
			expected: 2,
		},
	}

	for _, tt := range tests {
		elems := []*jsonast.JsonValue{}
		for _, s := range tt.values {
			elems = append(elems, &jsonast.JsonValue{String: pstr(s)})
		}
		v := &jsonast.JsonArray{Elements: elems}
		assert.Equal(t, tt.expected, v.Len())
	}
}

func TestArrayIsFalseArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected bool
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{False: pfalse("")}, {Array: &jsonast.JsonArray{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{False: pfalse("")}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{False: pfalse("")}, {False: pfalse("")}},
			expected: true,
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.IsFalseArray())
	}
}

func TestArrayFalseArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected []*jsonast.JsonFalse
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{False: pfalse("")}, {Array: &jsonast.JsonArray{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{False: pfalse("")}},
			expected: []*jsonast.JsonFalse{pfalse("")},
		},
		{
			values:   []*jsonast.JsonValue{{False: pfalse("")}, {False: pfalse("")}},
			expected: []*jsonast.JsonFalse{pfalse(""), pfalse("")},
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.FalseArray())
	}
}

func TestArrayIsNullArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected bool
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{Null: pnull("")}, {Array: &jsonast.JsonArray{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{Null: pnull("")}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{Null: pnull("")}, {Null: pnull("")}},
			expected: true,
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.IsNullArray())
	}
}

func TestArrayNullArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected []*jsonast.JsonNull
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{Null: pnull("")}, {Array: &jsonast.JsonArray{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{Null: pnull("")}},
			expected: []*jsonast.JsonNull{pnull("")},
		},
		{
			values:   []*jsonast.JsonValue{{Null: pnull("")}, {Null: pnull("")}},
			expected: []*jsonast.JsonNull{pnull(""), pnull("")},
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.NullArray())
	}
}
func TestArrayIsTrueArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected bool
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{True: ptrue("")}, {Array: &jsonast.JsonArray{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{True: ptrue("")}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{True: ptrue("")}, {True: ptrue("")}},
			expected: true,
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.IsTrueArray())
	}
}

func TestArrayTrueArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected []*jsonast.JsonTrue
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{True: ptrue("")}, {Array: &jsonast.JsonArray{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{True: ptrue("")}},
			expected: []*jsonast.JsonTrue{ptrue("")},
		},
		{
			values:   []*jsonast.JsonValue{{True: ptrue("")}, {True: ptrue("")}},
			expected: []*jsonast.JsonTrue{ptrue(""), ptrue("")},
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.TrueArray())
	}
}

func TestArrayIsObjectArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected bool
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {Object: &jsonast.JsonObject{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{Object: &jsonast.JsonObject{}}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{Object: &jsonast.JsonObject{}}, {Object: &jsonast.JsonObject{}}},
			expected: true,
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.IsObjectArray())
	}
}

func TestArrayObjectArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected []*jsonast.JsonObject
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {Object: &jsonast.JsonObject{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{Object: &jsonast.JsonObject{}}},
			expected: []*jsonast.JsonObject{{}},
		},
		{
			values:   []*jsonast.JsonValue{{Object: &jsonast.JsonObject{}}, {Object: &jsonast.JsonObject{}}},
			expected: []*jsonast.JsonObject{{}, {}},
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.ObjectArray())
	}
}

func TestArrayIsArrayArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected bool
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {Array: &jsonast.JsonArray{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{Array: &jsonast.JsonArray{}}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{Array: &jsonast.JsonArray{}}, {Array: &jsonast.JsonArray{}}},
			expected: true,
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.IsArrayArray())
	}
}

func TestArrayArrayArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected []*jsonast.JsonArray
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {Array: &jsonast.JsonArray{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{Array: &jsonast.JsonArray{}}},
			expected: []*jsonast.JsonArray{{}},
		},
		{
			values:   []*jsonast.JsonValue{{Array: &jsonast.JsonArray{}}, {Array: &jsonast.JsonArray{}}},
			expected: []*jsonast.JsonArray{{}, {}},
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.ArrayArray())
	}
}

func TestArrayIsNumberArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected bool
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{Number: pnum("")}, {Array: &jsonast.JsonArray{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{Number: pnum("")}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{Number: pnum("")}, {Number: pnum("")}},
			expected: true,
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.IsNumberArray())
	}
}

func TestArrayNumberArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected []*jsonast.JsonNumber
	}{
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{Number: pnum("")}, {Array: &jsonast.JsonArray{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{Number: pnum("")}},
			expected: []*jsonast.JsonNumber{pnum("")},
		},
		{
			values:   []*jsonast.JsonValue{{Number: pnum("")}, {Number: pnum("")}},
			expected: []*jsonast.JsonNumber{pnum(""), pnum("")},
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.NumberArray())
	}
}

func TestArrayIsStringArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected bool
	}{
		{
			values:   []*jsonast.JsonValue{{Number: pnum("")}, {Number: pnum("")}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {Array: &jsonast.JsonArray{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: true,
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.IsStringArray())
	}
}

func TestArrayStringArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected []*jsonast.JsonString
	}{
		{
			values:   []*jsonast.JsonValue{{Number: pnum("")}, {Number: pnum("")}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {Array: &jsonast.JsonArray{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}},
			expected: []*jsonast.JsonString{pstr("")},
		},
		{
			values:   []*jsonast.JsonValue{{String: pstr("")}, {String: pstr("")}},
			expected: []*jsonast.JsonString{pstr(""), pstr("")},
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.StringArray())
	}
}
