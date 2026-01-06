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
				Number: vnum("1"),
			},
		},
		{
			name: "float",
			json: "1.1",
			expected: &jsonast.JsonValue{
				Number: vnum("1.1"),
			},
		},
		{
			name: "false",
			json: "false",
			expected: &jsonast.JsonValue{
				False: vfalse(),
			},
		},
		{
			name: "null",
			json: "null",
			expected: &jsonast.JsonValue{
				Null: vnull(),
			},
		},
		{
			name: "true",
			json: "true",
			expected: &jsonast.JsonValue{
				True: vtrue(),
			},
		},
		{
			name: "string",
			json: `"hello"`,
			expected: &jsonast.JsonValue{
				String: vstr("hello"),
			},
		},
		{
			name: "true-string",
			json: `"true"`,
			expected: &jsonast.JsonValue{
				String: vstr("true"),
			},
		},
		{
			name: "false-string",
			json: `"false"`,
			expected: &jsonast.JsonValue{
				String: vstr("false"),
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
								String: vstr("s"),
							},
						},
						{
							Key: "num",
							Value: &jsonast.JsonValue{
								Number: vnum("1"),
							},
						},
						{
							Key: "t",
							Value: &jsonast.JsonValue{
								True: vtrue(),
							},
						},
						{
							Key: "f",
							Value: &jsonast.JsonValue{
								False: vfalse(),
							},
						},
						{
							Key: "null",
							Value: &jsonast.JsonValue{
								Null: vnull(),
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
												String: vstr("s"),
											},
										},
										{
											Key: "num",
											Value: &jsonast.JsonValue{
												Number: vnum("1"),
											},
										},
										{
											Key: "t",
											Value: &jsonast.JsonValue{
												True: vtrue(),
											},
										},
										{
											Key: "f",
											Value: &jsonast.JsonValue{
												False: vfalse(),
											},
										},
										{
											Key: "null",
											Value: &jsonast.JsonValue{
												Null: vnull(),
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
											String: vstr("s"),
										},
										{
											Number: vnum("1"),
										},
										{
											True: vtrue(),
										},
										{
											False: vfalse(),
										},
										{
											Null: vnull(),
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
								String: vstr("s"),
							},
						},
						{
							Key: "num",
							Value: &jsonast.JsonValue{
								Number: vnum("1"),
							},
						},
						{
							Key: "t",
							Value: &jsonast.JsonValue{
								True: vtrue(),
							},
						},
						{
							Key: "f",
							Value: &jsonast.JsonValue{
								False: vfalse(),
							},
						},
						{
							Key: "null",
							Value: &jsonast.JsonValue{
								Null: vnull(),
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
															String: vstr("s"),
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
															Number: vnum("1"),
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
															True: vtrue(),
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
															False: vfalse(),
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
															Null: vnull(),
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
																		String: vstr("s"),
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
																		Number: vnum("1"),
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
																		True: vtrue(),
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
																		False: vfalse(),
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
																		Null: vnull(),
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
											String: vstr("s"),
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
											Number: vnum("1"),
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
											True: vtrue(),
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
											False: vfalse(),
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
											Null: vnull(),
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
														String: vstr("s"),
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
														Number: vnum("1"),
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
														True: vtrue(),
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
														False: vfalse(),
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
														Null: vnull(),
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

func TestSameTypeAs(t *testing.T) {
	tests := []string{
		"False",
		"Null",
		"True",
		"Object",
		"Array",
		"Number",
		"String",
	}

	for _, name := range tests {
		others := []string{
			"False",
			"Null",
			"True",
			"Object",
			"Array",
			"Number",
			"String",
		}

		for _, other := range others {

			t.Run(name+"<=>"+other, func(t *testing.T) {
				valueOf := func(s string) *jsonast.JsonValue {
					v := &jsonast.JsonValue{}
					switch s {
					case "False":
						v.False = vfalse()
					case "Null":
						v.Null = vnull()
					case "True":
						v.True = vtrue()
					case "Object":
						v.Object = &jsonast.JsonObject{}
					case "Array":
						v.Array = &jsonast.JsonArray{}
					case "Number":
						v.Number = vnum("")
					case "String":
						v.String = vstr("")
					}
					return v
				}

				assert.Equal(t, valueOf(name).SameTypeAs(valueOf(other)), name == other)
			})
		}
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
				v.False = vfalse()
			}
			if tt.Null {
				v.Null = vnull()
			}
			if tt.True {
				v.True = vtrue()
			}
			if tt.Object {
				v.Object = &jsonast.JsonObject{}
			}
			if tt.Array {
				v.Array = &jsonast.JsonArray{}
			}
			if tt.Number {
				v.Number = vnum("")
			}
			if tt.String {
				v.String = vstr("")
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

func TestValue(t *testing.T) {
	tests := []struct {
		name     string
		False    bool
		Null     bool
		True     bool
		Object   bool
		Array    bool
		Number   bool
		String   bool
		expected any
	}{
		{name: "ValueOfFalse", False: true, expected: vfalse()},
		{name: "ValueOfNull", Null: true, expected: vnull()},
		{name: "ValueOfTrue", True: true, expected: vtrue()},
		{name: "ValueOfObject", Object: true, expected: &jsonast.JsonObject{}},
		{name: "ValueOfArray", Array: true, expected: &jsonast.JsonArray{}},
		{name: "ValueOfNumber", Number: true, expected: vnum("")},
		{name: "ValueOfString", String: true, expected: vstr("")},
		{name: "ValueOfNil", expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &jsonast.JsonValue{}
			if tt.False {
				v.False = vfalse()
			}
			if tt.Null {
				v.Null = vnull()
			}
			if tt.True {
				v.True = vtrue()
			}
			if tt.Object {
				v.Object = &jsonast.JsonObject{}
			}
			if tt.Array {
				v.Array = &jsonast.JsonArray{}
			}
			if tt.Number {
				v.Number = vnum("")
			}
			if tt.String {
				v.String = vstr("")
			}

			assert.Equal(t, tt.expected, v.Value())
		})
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
			elems = append(elems, &jsonast.JsonValue{String: vstr(s)})
		}
		v := &jsonast.JsonArray{Elements: elems}
		assert.Equal(t, tt.expected, v.Len())
	}
}
