// Package parser provides utilities for parsing and converting data types
// within Directed Acyclic Graphs (DAGs). This package includes functions
// to facilitate the conversion of various data types to a specified type,
// ensuring type safety and ease of use when working with generic data structures.
//
// The primary function in this package, ToAnyType, allows for the conversion
// of an input of any type to a specified type T. This is particularly useful
// when dealing with data that may come in different forms and needs to be
// processed in a type-safe manner.
//
// Example usage:
//
//	package main
//
//	import (
//	    "fmt"
//	    "pkg/parser"
//	)
//
//	func main() {
//	    var input interface{} = 123
//	    result := parser.ToAnyType[int](input)
//	    if result != nil {
//	        fmt.Println(*result) // Output: 123
//	    } else {
//	        fmt.Println("Conversion failed")
//	    }
//	}
package parser

// ToAnyType attempts to convert an input of any type to the specified type T.
// It returns a pointer to the converted value if successful, or nil if the conversion fails.
//
// Parameters:
//   - input: An interface{} representing the value to be converted.
//
// Returns:
//   - A pointer to the converted value of type T, or nil if the conversion fails.
//
// Example:
//
//	var input interface{} = 123
//	result := ToAnyType[int](input)
//	if result != nil {
//	    fmt.Println(*result) // Output: 123
//	} else {
//	    fmt.Println("Conversion failed")
//	}
func ToAnyType[T any](input interface{}) *T {
	// Check if the input is already of type T.
	if v, ok := input.(T); ok {
		return &v
	}

	// Check if the input is a pointer to type T.
	if v, ok := input.(*T); ok {
		return v
	}

	// Handle other potential cases.
	return nil
}
