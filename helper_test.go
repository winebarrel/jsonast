package jsonast_test

import (
	"github.com/winebarrel/jsonast"
)

func vstr(v string) *jsonast.JsonString {
	s := &jsonast.JsonString{}
	s.UnmarshalText([]byte(v)) //nolint:errcheck
	return s
}

func pstr(v string) *jsonast.JsonString {
	s := vstr(v)
	jsonast.MakeStrNullable(s)
	return s
}

func vnum(v string) *jsonast.JsonNumber {
	n := &jsonast.JsonNumber{}
	n.UnmarshalText([]byte(v)) //nolint:errcheck
	return n
}

func pnum(v string) *jsonast.JsonNumber {
	n := vnum(v)
	jsonast.MakeNumNullable(n)
	return n
}

func vtrue() *jsonast.JsonTrue {
	return &jsonast.JsonTrue{}
}

func ptrue() *jsonast.JsonTrue {
	t := vtrue()
	jsonast.MakeTrueNullable(t)
	return t
}

func vfalse() *jsonast.JsonFalse {
	return &jsonast.JsonFalse{}
}

func pfalse() *jsonast.JsonFalse {
	f := vfalse()
	jsonast.MakeFalseNullable(f)
	return f
}

func vnull() *jsonast.JsonNull {
	return &jsonast.JsonNull{}
}

func anynull() *jsonast.JsonNull {
	null := &jsonast.JsonNull{}
	jsonast.MakeNullAny(null)
	return null
}
