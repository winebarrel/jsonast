package jsonast_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/jsonast"
)

func TestParse_EmptyString(t *testing.T) {
	_, err := jsonast.ParseBytes("<filename>", []byte(""))
	assert.ErrorContains(t, err, `<filename>:1:1: unexpected token "<EOF>"`)
	_, err = jsonast.Parse("<filename>", strings.NewReader(""))
	assert.ErrorContains(t, err, `<filename>:1:1: unexpected token "<EOF>"`)
}

func TestParse_ParseErr(t *testing.T) {
	_, err := jsonast.ParseBytes("<filename>", []byte(`{`))
	assert.ErrorContains(t, err, `<filename>:1:2: unexpected token "<EOF>" (expected "}")`)
	_, err = jsonast.Parse("<filename>", strings.NewReader(`{`))
	assert.ErrorContains(t, err, `<filename>:1:2: unexpected token "<EOF>" (expected "}")`)
}

func TestParse_LexErr(t *testing.T) {
	_, err := jsonast.ParseBytes("<filename>", []byte(`{"foo:"bar"}`))
	assert.ErrorContains(t, err, `<filename>:1:8: invalid character 'b' after object key`)
	_, err = jsonast.Parse("<filename>", strings.NewReader(`{"foo:"bar"}`))
	assert.ErrorContains(t, err, `<filename>:1:8: invalid character 'b' after object key`)
}

func TestParse_OK(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		expected *jsonast.JsonValue
	}{
		{
			name: "int",
			json: "1",
			expected: &jsonast.JsonValue{
				Number: ptr("1"),
			},
		},
		{
			name: "float",
			json: "1.1",
			expected: &jsonast.JsonValue{
				Number: ptr("1.1"),
			},
		},
		{
			name: "false",
			json: "false",
			expected: &jsonast.JsonValue{
				False: ptr("false"),
			},
		},
		{
			name: "null",
			json: "null",
			expected: &jsonast.JsonValue{
				Null: ptr("null"),
			},
		},
		{
			name: "true",
			json: "true",
			expected: &jsonast.JsonValue{
				True: ptr("true"),
			},
		},
		{
			name: "string",
			json: `"hello"`,
			expected: &jsonast.JsonValue{
				String: ptr("hello"),
			},
		},
		{
			name: "true-string",
			json: `"true"`,
			expected: &jsonast.JsonValue{
				String: ptr("true"),
			},
		},
		{
			name: "false-string",
			json: `"false"`,
			expected: &jsonast.JsonValue{
				String: ptr("false"),
			},
		},
		{
			name: "empty object",
			json: "{}",
			expected: &jsonast.JsonValue{
				Object: &jsonast.JsonObject{},
			},
		},
		{
			name: "object",
			json: `{"str":"s","num":1,"t":true,"f":false,"null":null,"obj":{"str":"s","num":1,"t":true,"f":false,"null":null},"ary":["s",1,true,false,null]}`,
			expected: &jsonast.JsonValue{
				Object: &jsonast.JsonObject{
					Members: []*jsonast.JsonObjectMember{
						{
							Key: "str",
							Value: &jsonast.JsonValue{
								String: ptr("s"),
							},
						},
						{
							Key: "num",
							Value: &jsonast.JsonValue{
								Number: ptr("1"),
							},
						},
						{
							Key: "t",
							Value: &jsonast.JsonValue{
								True: ptr("true"),
							},
						},
						{
							Key: "f",
							Value: &jsonast.JsonValue{
								False: ptr("false"),
							},
						},
						{
							Key: "null",
							Value: &jsonast.JsonValue{
								Null: ptr("null"),
							},
						},
						{
							Key: "obj",
							Value: &jsonast.JsonValue{
								Object: &jsonast.JsonObject{
									Members: []*jsonast.JsonObjectMember{
										{
											Key: "str",
											Value: &jsonast.JsonValue{
												String: ptr("s"),
											},
										},
										{
											Key: "num",
											Value: &jsonast.JsonValue{
												Number: ptr("1"),
											},
										},
										{
											Key: "t",
											Value: &jsonast.JsonValue{
												True: ptr("true"),
											},
										},
										{
											Key: "f",
											Value: &jsonast.JsonValue{
												False: ptr("false"),
											},
										},
										{
											Key: "null",
											Value: &jsonast.JsonValue{
												Null: ptr("null"),
											},
										},
									},
								},
							},
						},
						{
							Key: "ary",
							Value: &jsonast.JsonValue{
								Array: &jsonast.JsonArray{
									Elements: []*jsonast.JsonValue{
										{
											String: ptr("s"),
										},
										{
											Number: ptr("1"),
										},
										{
											True: ptr("true"),
										},
										{
											False: ptr("false"),
										},
										{
											Null: ptr("null"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "array in object",
			json: `{"str":"s","num":1,"t":true,"f":false,"null":null,"ary":[{"str":"s"},{"num":1},{"t":true},{"f":false},{"null":null},[{"str":"s"},{"num":1},{"t":true},{"f":false},{"null":null}]]}`,
			expected: &jsonast.JsonValue{
				Object: &jsonast.JsonObject{
					Members: []*jsonast.JsonObjectMember{
						{
							Key: "str",
							Value: &jsonast.JsonValue{
								String: ptr("s"),
							},
						},
						{
							Key: "num",
							Value: &jsonast.JsonValue{
								Number: ptr("1"),
							},
						},
						{
							Key: "t",
							Value: &jsonast.JsonValue{
								True: ptr("true"),
							},
						},
						{
							Key: "f",
							Value: &jsonast.JsonValue{
								False: ptr("false"),
							},
						},
						{
							Key: "null",
							Value: &jsonast.JsonValue{
								Null: ptr("null"),
							},
						},
						{
							Key: "ary",
							Value: &jsonast.JsonValue{
								Array: &jsonast.JsonArray{
									Elements: []*jsonast.JsonValue{
										{
											Object: &jsonast.JsonObject{
												Members: []*jsonast.JsonObjectMember{
													{
														Key: "str",
														Value: &jsonast.JsonValue{
															String: ptr("s"),
														},
													},
												},
											},
										},
										{
											Object: &jsonast.JsonObject{
												Members: []*jsonast.JsonObjectMember{
													{
														Key: "num",
														Value: &jsonast.JsonValue{
															Number: ptr("1"),
														},
													},
												},
											},
										},
										{
											Object: &jsonast.JsonObject{
												Members: []*jsonast.JsonObjectMember{
													{
														Key: "t",
														Value: &jsonast.JsonValue{
															True: ptr("true"),
														},
													},
												},
											},
										},
										{
											Object: &jsonast.JsonObject{
												Members: []*jsonast.JsonObjectMember{
													{
														Key: "f",
														Value: &jsonast.JsonValue{
															False: ptr("false"),
														},
													},
												},
											},
										},
										{
											Object: &jsonast.JsonObject{
												Members: []*jsonast.JsonObjectMember{
													{
														Key: "null",
														Value: &jsonast.JsonValue{
															Null: ptr("null"),
														},
													},
												},
											},
										},
										{
											Array: &jsonast.JsonArray{
												Elements: []*jsonast.JsonValue{
													{
														Object: &jsonast.JsonObject{
															Members: []*jsonast.JsonObjectMember{
																{
																	Key: "str",
																	Value: &jsonast.JsonValue{
																		String: ptr("s"),
																	},
																},
															},
														},
													},
													{
														Object: &jsonast.JsonObject{
															Members: []*jsonast.JsonObjectMember{
																{
																	Key: "num",
																	Value: &jsonast.JsonValue{
																		Number: ptr("1"),
																	},
																},
															},
														},
													},
													{
														Object: &jsonast.JsonObject{
															Members: []*jsonast.JsonObjectMember{
																{
																	Key: "t",
																	Value: &jsonast.JsonValue{
																		True: ptr("true"),
																	},
																},
															},
														},
													},
													{
														Object: &jsonast.JsonObject{
															Members: []*jsonast.JsonObjectMember{
																{
																	Key: "f",
																	Value: &jsonast.JsonValue{
																		False: ptr("false"),
																	},
																},
															},
														},
													},
													{
														Object: &jsonast.JsonObject{
															Members: []*jsonast.JsonObjectMember{
																{
																	Key: "null",
																	Value: &jsonast.JsonValue{
																		Null: ptr("null"),
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "empty array",
			json: "[]",
			expected: &jsonast.JsonValue{
				Array: &jsonast.JsonArray{},
			},
		},
		{
			name: "object in array",
			json: `[{"str":"s"},{"num":1},{"t":true},{"f":false},{"null":null},[{"str":"s"},{"num":1},{"t":true},{"f":false},{"null":null}]]`,
			expected: &jsonast.JsonValue{
				Array: &jsonast.JsonArray{
					Elements: []*jsonast.JsonValue{
						{
							Object: &jsonast.JsonObject{
								Members: []*jsonast.JsonObjectMember{
									{
										Key: "str",
										Value: &jsonast.JsonValue{
											String: ptr("s"),
										},
									},
								},
							},
						},
						{
							Object: &jsonast.JsonObject{
								Members: []*jsonast.JsonObjectMember{
									{
										Key: "num",
										Value: &jsonast.JsonValue{
											Number: ptr("1"),
										},
									},
								},
							},
						},
						{
							Object: &jsonast.JsonObject{
								Members: []*jsonast.JsonObjectMember{
									{
										Key: "t",
										Value: &jsonast.JsonValue{
											True: ptr("true"),
										},
									},
								},
							},
						},
						{
							Object: &jsonast.JsonObject{
								Members: []*jsonast.JsonObjectMember{
									{
										Key: "f",
										Value: &jsonast.JsonValue{
											False: ptr("false"),
										},
									},
								},
							},
						},
						{
							Object: &jsonast.JsonObject{
								Members: []*jsonast.JsonObjectMember{
									{
										Key: "null",
										Value: &jsonast.JsonValue{
											Null: ptr("null"),
										},
									},
								},
							},
						},
						{
							Array: &jsonast.JsonArray{
								Elements: []*jsonast.JsonValue{
									{
										Object: &jsonast.JsonObject{
											Members: []*jsonast.JsonObjectMember{
												{
													Key: "str",
													Value: &jsonast.JsonValue{
														String: ptr("s"),
													},
												},
											},
										},
									},
									{
										Object: &jsonast.JsonObject{
											Members: []*jsonast.JsonObjectMember{
												{
													Key: "num",
													Value: &jsonast.JsonValue{
														Number: ptr("1"),
													},
												},
											},
										},
									},
									{
										Object: &jsonast.JsonObject{
											Members: []*jsonast.JsonObjectMember{
												{
													Key: "t",
													Value: &jsonast.JsonValue{
														True: ptr("true"),
													},
												},
											},
										},
									},
									{
										Object: &jsonast.JsonObject{
											Members: []*jsonast.JsonObjectMember{
												{
													Key: "f",
													Value: &jsonast.JsonValue{
														False: ptr("false"),
													},
												},
											},
										},
									},
									{
										Object: &jsonast.JsonObject{
											Members: []*jsonast.JsonObjectMember{
												{
													Key: "null",
													Value: &jsonast.JsonValue{
														Null: ptr("null"),
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := jsonast.ParseBytes("", []byte(tt.json))
			require.NoError(t, err)
			assert.Equal(t, tt.expected, v)
			v, err = jsonast.Parse("", strings.NewReader(tt.json))
			require.NoError(t, err)
			assert.Equal(t, tt.expected, v)
		})
	}
}

func TestIsXXX(t *testing.T) {
	tests := []struct {
		name   string
		False  bool
		Null   bool
		True   bool
		Object bool
		Array  bool
		Number bool
		String bool
	}{
		{name: "IsFalse", False: true},
		{name: "IsNull", Null: true},
		{name: "IsTrue", True: true},
		{name: "IsObject", Object: true},
		{name: "IsArray", Array: true},
		{name: "IsNumber", Number: true},
		{name: "IsString", String: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &jsonast.JsonValue{}
			if tt.False {
				v.False = ptr("")
			}
			if tt.Null {
				v.Null = ptr("")
			}
			if tt.True {
				v.True = ptr("")
			}
			if tt.Object {
				v.Object = &jsonast.JsonObject{}
			}
			if tt.Array {
				v.Array = &jsonast.JsonArray{}
			}
			if tt.Number {
				v.Number = ptr("")
			}
			if tt.String {
				v.String = ptr("")
			}

			assert.Equal(t, tt.False, v.IsFalse())
			assert.Equal(t, tt.Null, v.IsNull())
			assert.Equal(t, tt.True, v.IsTrue())
			assert.Equal(t, tt.Object, v.IsObject())
			assert.Equal(t, tt.Array, v.IsArray())
			assert.Equal(t, tt.Number, v.IsNumber())
			assert.Equal(t, tt.String, v.IsString())
		})
	}
}

func TestIsObjectArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected bool
	}{
		{
			values:   []*jsonast.JsonValue{{String: ptr("")}, {String: ptr("")}},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: false,
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr("")}, {Object: &jsonast.JsonObject{}}},
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
		v := &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: tt.values}}
		assert.Equal(t, tt.expected, v.IsObjectArray())
	}
}

func TestObjectArray(t *testing.T) {
	tests := []struct {
		values   []*jsonast.JsonValue
		expected []*jsonast.JsonObject
	}{
		{
			values:   []*jsonast.JsonValue{{String: ptr("")}, {String: ptr("")}},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{},
			expected: nil,
		},
		{
			values:   []*jsonast.JsonValue{{String: ptr("")}, {Object: &jsonast.JsonObject{}}},
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
		v := &jsonast.JsonValue{Array: &jsonast.JsonArray{Elements: tt.values}}
		assert.Equal(t, tt.expected, v.ObjectArray())
	}
}

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
			elems = append(elems, &jsonast.JsonValue{String: &s})
		}
		v := &jsonast.JsonArray{Elements: elems}
		assert.Equal(t, v.Len(), tt.expected)
	}
}
