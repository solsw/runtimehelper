package runtimehelper

import (
	"math"
	"runtime"
)

// NthCallerName returns the name of the n-th caller function of NthCallerName.
// In case of any failure an empty string is returned.
// For n's meaning see [runtime.Caller] documentation.
func NthCallerName(n int) string {
	// runtime.Caller adds one to the skip count internally.
	if n < 0 || n == math.MaxInt {
		return ""
	}
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

// CallerName returns the name of the function that called CallerName.
// In case of any failure an empty string is returned.
//
// The //go:noinline pragma keeps CallerName as its own stack frame so that
// the fixed skip count passed to NthCallerName stays correct.
//
//go:noinline
func CallerName() string {
	return NthCallerName(2)
}

// CallerCallerName returns the name of the function that called the function that called CallerCallerName.
// In case of any failure an empty string is returned.
//
// The //go:noinline pragma keeps CallerCallerName as its own stack frame so that
// the fixed skip count passed to NthCallerName stays correct.
//
//go:noinline
func CallerCallerName() string {
	return NthCallerName(3)
}
