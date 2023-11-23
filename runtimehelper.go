package runtimehelper

import (
	"runtime"
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
