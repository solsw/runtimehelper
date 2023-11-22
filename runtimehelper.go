package runtimehelper

import (
	"runtime"
)

// nthCallerName returns name of the n-th caller function of nthCallerName.
func nthCallerName(n int) string {
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
	return nthCallerName(2)
}

// CallerCallerName returns name of the function that called the function that called CallerCallerName.
// In case of any failure empty string is returned.
func CallerCallerName() string {
	return nthCallerName(3)
}
