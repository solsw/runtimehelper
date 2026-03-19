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
		{name: "2",
			args: args{n: 2},
			want: "testing.tRunner",
		},
		{name: "3",
			args: args{n: 3},
			want: "runtime.goexit",
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

func TestJustName(t *testing.T) {
	tests := []struct {
		name string
		nm   string
		want string
	}{
		{name: "0",
			nm:   NthCallerName(0),
			want: "NthCallerName",
		},
		{name: "2",
			nm:   NthCallerName(2),
			want: "tRunner",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JustName(tt.nm)
			if got != tt.want {
				t.Errorf("JustName() = %v, want %v", got, tt.want)
			}
		})
	}
}
