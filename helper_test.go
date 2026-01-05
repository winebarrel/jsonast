package jsonast_test

import (
	"github.com/winebarrel/jsonast"
)

func vstr(v string) *jsonast.JsonString {
	s := &jsonast.JsonString{}
	s.UnmarshalText([]byte(v))
	return s
}

func vnum(v string) *jsonast.JsonNumber {
	n := &jsonast.JsonNumber{}
	n.UnmarshalText([]byte(v))
	return n
}

func vtrue() *jsonast.JsonTrue {
	return &jsonast.JsonTrue{}
}

func vfalse() *jsonast.JsonFalse {
	return &jsonast.JsonFalse{}
}

func vnull() *jsonast.JsonNull {
	return &jsonast.JsonNull{}
}
