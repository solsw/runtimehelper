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
func CallerName() string {
	return NthCallerName(2)
}

// CallerCallerName returns name of the function that called the function that called CallerCallerName.
// In case of any failure empty string is returned.
func CallerCallerName() string {
	return NthCallerName(3)
}

// JustNameWithPackage returns just function name preceeded with package name
// of the full function name returned by [Func.Name].
//
// [Func.Name]: https://pkg.go.dev/runtime#Func.Name
func JustNameWithPackage(nm string) string {
	r := path.Base(nm)
	r, _, _ = strings.Cut(r, "[")
	return r
}

// JustName returns just function name of the full function name returned by [Func.Name].
//
// [Func.Name]: https://pkg.go.dev/runtime#Func.Name
func JustName(nm string) string {
	r := JustNameWithPackage(nm)
	_, r, _ = strings.Cut(r, ".")
	return r
}
