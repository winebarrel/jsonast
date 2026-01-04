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
				Number: pnum("1"),
			},
		},
		{
			name: "float",
			json: "1.1",
			expected: &jsonast.JsonValue{
				Number: pnum("1.1"),
			},
		},
		{
			name: "false",
			json: "false",
			expected: &jsonast.JsonValue{
				False: pfalse("false"),
			},
		},
		{
			name: "null",
			json: "null",
			expected: &jsonast.JsonValue{
				Null: pnull("null"),
			},
		},
		{
			name: "true",
			json: "true",
			expected: &jsonast.JsonValue{
				True: ptrue("true"),
			},
		},
		{
			name: "string",
			json: `"hello"`,
			expected: &jsonast.JsonValue{
				String: pstr("hello"),
			},
		},
		{
			name: "true-string",
			json: `"true"`,
			expected: &jsonast.JsonValue{
				String: pstr("true"),
			},
		},
		{
			name: "false-string",
			json: `"false"`,
			expected: &jsonast.JsonValue{
				String: pstr("false"),
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
								String: pstr("s"),
							},
						},
						{
							Key: "num",
							Value: &jsonast.JsonValue{
								Number: pnum("1"),
							},
						},
						{
							Key: "t",
							Value: &jsonast.JsonValue{
								True: ptrue("true"),
							},
						},
						{
							Key: "f",
							Value: &jsonast.JsonValue{
								False: pfalse("false"),
							},
						},
						{
							Key: "null",
							Value: &jsonast.JsonValue{
								Null: pnull("null"),
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
												String: pstr("s"),
											},
										},
										{
											Key: "num",
											Value: &jsonast.JsonValue{
												Number: pnum("1"),
											},
										},
										{
											Key: "t",
											Value: &jsonast.JsonValue{
												True: ptrue("true"),
											},
										},
										{
											Key: "f",
											Value: &jsonast.JsonValue{
												False: pfalse("false"),
											},
										},
										{
											Key: "null",
											Value: &jsonast.JsonValue{
												Null: pnull("null"),
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
											String: pstr("s"),
										},
										{
											Number: pnum("1"),
										},
										{
											True: ptrue("true"),
										},
										{
											False: pfalse("false"),
										},
										{
											Null: pnull("null"),
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
								String: pstr("s"),
							},
						},
						{
							Key: "num",
							Value: &jsonast.JsonValue{
								Number: pnum("1"),
							},
						},
						{
							Key: "t",
							Value: &jsonast.JsonValue{
								True: ptrue("true"),
							},
						},
						{
							Key: "f",
							Value: &jsonast.JsonValue{
								False: pfalse("false"),
							},
						},
						{
							Key: "null",
							Value: &jsonast.JsonValue{
								Null: pnull("null"),
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
															String: pstr("s"),
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
															Number: pnum("1"),
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
															True: ptrue("true"),
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
															False: pfalse("false"),
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
															Null: pnull("null"),
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
																		String: pstr("s"),
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
																		Number: pnum("1"),
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
																		True: ptrue("true"),
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
																		False: pfalse("false"),
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
																		Null: pnull("null"),
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
											String: pstr("s"),
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
											Number: pnum("1"),
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
											True: ptrue("true"),
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
											False: pfalse("false"),
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
											Null: pnull("null"),
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
														String: pstr("s"),
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
														Number: pnum("1"),
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
														True: ptrue("true"),
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
														False: pfalse("false"),
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
														Null: pnull("null"),
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
		name      string
		False     bool
		Null      bool
		True      bool
		Object    bool
		Array     bool
		Number    bool
		String    bool
		primitive bool
	}{
		{name: "IsFalse", False: true, primitive: true},
		{name: "IsNull", Null: true, primitive: true},
		{name: "IsTrue", True: true, primitive: true},
		{name: "IsObject", Object: true},
		{name: "IsArray", Array: true},
		{name: "IsNumber", Number: true, primitive: true},
		{name: "IsString", String: true, primitive: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &jsonast.JsonValue{}
			if tt.False {
				v.False = pfalse("")
			}
			if tt.Null {
				v.Null = pnull("")
			}
			if tt.True {
				v.True = ptrue("")
			}
			if tt.Object {
				v.Object = &jsonast.JsonObject{}
			}
			if tt.Array {
				v.Array = &jsonast.JsonArray{}
			}
			if tt.Number {
				v.Number = pnum("")
			}
			if tt.String {
				v.String = pstr("")
			}

			assert.Equal(t, tt.False, v.IsFalse())
			assert.Equal(t, tt.Null, v.IsNull())
			assert.Equal(t, tt.True, v.IsTrue())
			assert.Equal(t, tt.Object, v.IsObject())
			assert.Equal(t, tt.Array, v.IsArray())
			assert.Equal(t, tt.Number, v.IsNumber())
			assert.Equal(t, tt.String, v.IsString())
			assert.Equal(t, tt.primitive, v.IsPrimitive())
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
		{name: "ValueOfFalse", False: true, expected: pfalse("")},
		{name: "ValueOfNull", Null: true, expected: pnull("")},
		{name: "ValueOfTrue", True: true, expected: ptrue("")},
		{name: "ValueOfObject", Object: true, expected: &jsonast.JsonObject{}},
		{name: "ValueOfArray", Array: true, expected: &jsonast.JsonArray{}},
		{name: "ValueOfNumber", Number: true, expected: pnum("")},
		{name: "ValueOfString", String: true, expected: pstr("")},
		{name: "ValueOfNil", expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &jsonast.JsonValue{}
			if tt.False {
				v.False = pfalse("")
			}
			if tt.Null {
				v.Null = pnull("")
			}
			if tt.True {
				v.True = ptrue("")
			}
			if tt.Object {
				v.Object = &jsonast.JsonObject{}
			}
			if tt.Array {
				v.Array = &jsonast.JsonArray{}
			}
			if tt.Number {
				v.Number = pnum("")
			}
			if tt.String {
				v.String = pstr("")
			}

			assert.Equal(t, tt.expected, v.Value())
		})
	}
}
