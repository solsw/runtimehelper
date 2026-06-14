package runtimehelper

import (
	"path"
	"strings"
)

// JustPackageFunctionName returns just the function name preceded with
// the package name of the full function name (returned by [Func.Name]).
// Generic type parameters (e.g. "[int]") are stripped from the result.
//
// [Func.Name]: https://pkg.go.dev/runtime#Func.Name
func JustPackageFunctionName(funcName string) string {
	if funcName == "" {
		return ""
	}
	r := path.Base(funcName)
	r, _, _ = strings.Cut(r, "[")
	return r
}

// JustFunctionName returns just the function name of the full function
// name (returned by [Func.Name]).
// If 'funcName' contains no package separator, it is returned
// (with generic type parameters stripped) unchanged.
//
// [Func.Name]: https://pkg.go.dev/runtime#Func.Name
func JustFunctionName(funcName string) string {
	r := JustPackageFunctionName(funcName)
	if _, after, found := strings.Cut(r, "."); found {
		r = after
	}
	return r
}
