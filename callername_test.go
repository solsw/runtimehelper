package runtimehelper

import (
	"path"
	"runtime"
	"testing"
)

func TestNthCallerName(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0",
			args: args{n: 0},
			want: "runtimehelper.NthCallerName",
		},
		{name: "1",
			args: args{n: 1},
			want: "runtimehelper.TestNthCallerName.func1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NthCallerName(tt.args.n)
			got = path.Base(got)
			if got != tt.want {
				t.Errorf("NthCallerName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNthCallerNameOutOfRange(t *testing.T) {
	if got := NthCallerName(999); got != "" {
		t.Errorf("NthCallerName(999) = %q, want %q", got, "")
	}
}

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

//go:noinline
func helperCallerCallerName() string {
	return CallerCallerName()
}

func TestCallerCallerName(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	want := runtime.FuncForPC(pc).Name()
	got := helperCallerCallerName()
	if got != want {
		t.Errorf("CallerCallerName() got = %v, want %v", got, want)
	}
}
