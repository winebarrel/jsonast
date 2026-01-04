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
			elems = append(elems, &jsonast.JsonValue{String: ptr(jsonast.JsonString(s))})
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{False: ptr(jsonast.JsonFalse(""))}, {Array: &jsonast.JsonArray{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{False: ptr(jsonast.JsonFalse(""))}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{False: ptr(jsonast.JsonFalse(""))}, {False: ptr(jsonast.JsonFalse(""))}},
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{False: ptr(jsonast.JsonFalse(""))}, {Array: &jsonast.JsonArray{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{False: ptr(jsonast.JsonFalse(""))}},
			expected: []*jsonast.JsonFalse{ptr(jsonast.JsonFalse(""))},
		},
		{
			values:   []*jsonast.JsonValue{{False: ptr(jsonast.JsonFalse(""))}, {False: ptr(jsonast.JsonFalse(""))}},
			expected: []*jsonast.JsonFalse{ptr(jsonast.JsonFalse("")), ptr(jsonast.JsonFalse(""))},
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{Null: ptr(jsonast.JsonNull(""))}, {Array: &jsonast.JsonArray{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{Null: ptr(jsonast.JsonNull(""))}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{Null: ptr(jsonast.JsonNull(""))}, {Null: ptr(jsonast.JsonNull(""))}},
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{Null: ptr(jsonast.JsonNull(""))}, {Array: &jsonast.JsonArray{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{Null: ptr(jsonast.JsonNull(""))}},
			expected: []*jsonast.JsonNull{ptr(jsonast.JsonNull(""))},
		},
		{
			values:   []*jsonast.JsonValue{{Null: ptr(jsonast.JsonNull(""))}, {Null: ptr(jsonast.JsonNull(""))}},
			expected: []*jsonast.JsonNull{ptr(jsonast.JsonNull("")), ptr(jsonast.JsonNull(""))},
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{True: ptr(jsonast.JsonTrue(""))}, {Array: &jsonast.JsonArray{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{True: ptr(jsonast.JsonTrue(""))}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{True: ptr(jsonast.JsonTrue(""))}, {True: ptr(jsonast.JsonTrue(""))}},
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{True: ptr(jsonast.JsonTrue(""))}, {Array: &jsonast.JsonArray{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{True: ptr(jsonast.JsonTrue(""))}},
			expected: []*jsonast.JsonTrue{ptr(jsonast.JsonTrue(""))},
		},
		{
			values:   []*jsonast.JsonValue{{True: ptr(jsonast.JsonTrue(""))}, {True: ptr(jsonast.JsonTrue(""))}},
			expected: []*jsonast.JsonTrue{ptr(jsonast.JsonTrue("")), ptr(jsonast.JsonTrue(""))},
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {Object: &jsonast.JsonObject{}}},
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {Object: &jsonast.JsonObject{}}},
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {Array: &jsonast.JsonArray{}}},
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {Array: &jsonast.JsonArray{}}},
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{Number: ptr(jsonast.JsonNumber(""))}, {Array: &jsonast.JsonArray{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{Number: ptr(jsonast.JsonNumber(""))}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{Number: ptr(jsonast.JsonNumber(""))}, {Number: ptr(jsonast.JsonNumber(""))}},
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
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{Number: ptr(jsonast.JsonNumber(""))}, {Array: &jsonast.JsonArray{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{Number: ptr(jsonast.JsonNumber(""))}},
			expected: []*jsonast.JsonNumber{ptr(jsonast.JsonNumber(""))},
		},
		{
			values:   []*jsonast.JsonValue{{Number: ptr(jsonast.JsonNumber(""))}, {Number: ptr(jsonast.JsonNumber(""))}},
			expected: []*jsonast.JsonNumber{ptr(jsonast.JsonNumber("")), ptr(jsonast.JsonNumber(""))},
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
			values:   []*jsonast.JsonValue{{Number: ptr(jsonast.JsonNumber(""))}, {Number: ptr(jsonast.JsonNumber(""))}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {Array: &jsonast.JsonArray{}}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}},
			expected: true,
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
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
			values:   []*jsonast.JsonValue{{Number: ptr(jsonast.JsonNumber(""))}, {Number: ptr(jsonast.JsonNumber(""))}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {Array: &jsonast.JsonArray{}}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}},
			expected: []*jsonast.JsonString{ptr(jsonast.JsonString(""))},
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr(jsonast.JsonString(""))}, {String: ptr(jsonast.JsonString(""))}},
			expected: []*jsonast.JsonString{ptr(jsonast.JsonString("")), ptr(jsonast.JsonString(""))},
		},
	}

	for _, tt := range tests {
		v := &jsonast.JsonArray{Elements: tt.values}
		assert.Equal(t, tt.expected, v.StringArray())
	}
}
