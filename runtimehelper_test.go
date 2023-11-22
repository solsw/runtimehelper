package runtimehelper

import (
	"runtime"
	"testing"
)

func TestCallerName(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		pc, _, _, _ := runtime.Caller(0)
		want := runtime.FuncForPC(pc).Name()
		got := CallerName()
		if got != want {
			t.Errorf("CallerName() got = %v, want %v", got, want)
		}
	})
}
