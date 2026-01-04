package jsonast_test

import "github.com/winebarrel/jsonast"

func pstr(v string) *jsonast.JsonString {
	x := jsonast.JsonString(v)
	return &x
}

func pnum(v string) *jsonast.JsonNumber {
	x := jsonast.JsonNumber(v)
	return &x
}

func ptrue(v string) *jsonast.JsonTrue {
	x := jsonast.JsonTrue(v)
	return &x
}

func pfalse(v string) *jsonast.JsonFalse {
	x := jsonast.JsonFalse(v)
	return &x
}

func pnull(v string) *jsonast.JsonNull {
	x := jsonast.JsonNull(v)
	return &x
}
