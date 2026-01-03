# jsonast

jsonast is a library that parses JSON and converts it to AST.

## Usage

```go
package main

import (
	"fmt"

	"github.com/winebarrel/jsonast"
)

// func strptr(s string) *string { return &s }

func main() {
	json := `{"foo":"bar","zoo":[1,2,3],"baz":{"hoge":true,"fuga":null}}`
	ast, err := jsonast.ParseJSON("<filename>", []byte(json))

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ast)
	// ast:
	// 	&jsonast.JsonValue{
	// 		Object: &jsonast.JsonObject{
	// 			Members: []*jsonast.JsonObjectMember{
	// 				{Key: "foo", Value: &jsonast.JsonValue{String: strptr("bar")}},
	// 				{Key: "zoo", Value: &jsonast.JsonValue{Array: &jsonast.JsonArray{
	// 					Elements: []*jsonast.JsonValue{
	// 						{Number: strptr("1")},
	// 						{Number: strptr("2")},
	// 						{Number: strptr("3")},
	// 					},
	// 				}}},
	// 				{Key: "baz", Value: &jsonast.JsonValue{Object: &jsonast.JsonObject{
	// 					Members: []*jsonast.JsonObjectMember{
	// 						{Key: "hoge", Value: &jsonast.JsonValue{True: strptr("true")}},
	// 						{Key: "fuga", Value: &jsonast.JsonValue{Null: strptr("null")}},
	// 					},
	// 				}}},
	// 			},
	// 		},
	// 	}
}
```
