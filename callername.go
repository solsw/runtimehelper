package runtimehelper

import (
	"path"
	"runtime"
	"strings"
)

// NthCallerName returns name of the n-th caller function of NthCallerName.
// In case of any failure empty string is returned.
// For n's meaning see [runtime.Caller] documentation.
func NthCallerName(n int) string {
	pc, _, _, ok := runtime.Caller(n)
	if !ok {
		return ""
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return ""
	}
	return f.Name()
}

// CallerName returns name of the function that called CallerName.
// In case of any failure empty string is returned.
//
//go:noinline
func CallerName() string {
	return NthCallerName(2)
}

// CallerCallerName returns name of the function that called the function that called CallerCallerName.
// In case of any failure empty string is returned.
//
//go:noinline
func CallerCallerName() string {
	return NthCallerName(3)
}

// JustPackageFunctionName returns just function name preceded with package name
// of the full function name returned by [Func.Name].
// Generic type parameters (e.g. "[int]") are stripped from the result.
//
// [Func.Name]: https://pkg.go.dev/runtime#Func.Name
func JustPackageFunctionName(nm string) string {
	r := path.Base(nm)
	r, _, _ = strings.Cut(r, "[")
	return r
}

// JustFunctionName returns just function name of the full function name returned by [Func.Name].
//
// [Func.Name]: https://pkg.go.dev/runtime#Func.Name
func JustFunctionName(nm string) string {
	r := JustPackageFunctionName(nm)
	_, r, _ = strings.Cut(r, ".")
	return r
}
