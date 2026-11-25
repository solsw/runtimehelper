package runtimehelper

import (
	"path"
	"strings"
)

// JustPackageFunctionName returns just the function name preceded with
// the package name of the full function name (returned by [Func.Name]).
// Generic type arguments (e.g. "[int]") are stripped from the result.
//
// [Func.Name]: https://pkg.go.dev/runtime#Func.Name
func JustPackageFunctionName(funcName string) string {
	if strings.Contains(funcName, "[") {
		// Strip type arguments before path.Base since they can contain import paths.
		depth := 0
		funcName = strings.Map(func(r rune) rune {
			switch {
			case r == '[':
				depth++
			case r == ']' && depth > 0:
				depth--
			case depth == 0:
				return r
			}
			return -1
		}, funcName)
	}
	if funcName == "" {
		return ""
	}
	return path.Base(funcName)
}

// JustFunctionName returns just the function name of the full function
// name (returned by [Func.Name]).
// If 'funcName' contains no package separator, it is returned
// (with generic type arguments stripped) unchanged.
//
// [Func.Name]: https://pkg.go.dev/runtime#Func.Name
func JustFunctionName(funcName string) string {
	r := JustPackageFunctionName(funcName)
	if _, after, found := strings.Cut(r, "."); found {
		r = after
	}
	return r
}
