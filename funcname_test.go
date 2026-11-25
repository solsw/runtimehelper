package runtimehelper

import (
	"testing"
)

type Class struct{}

func (*Class) Method1() string {
	return NthCallerName(1)
}

func (c Class) Method2() string {
	return NthCallerName(1)
}

type box[T any] struct{}

func (*box[T]) method() string { return CallerName() }

func TestGenericReceiverName(t *testing.T) {
	name := new(box[int]).method()
	if got := JustPackageFunctionName(name); got != "runtimehelper.(*box).method" {
		t.Errorf("JustPackageFunctionName(%q) = %q, want runtimehelper.(*box).method", name, got)
	}
	if got := JustFunctionName(name); got != "(*box).method" {
		t.Errorf("JustFunctionName(%q) = %q, want (*box).method", name, got)
	}
}

func TestJustPackageFunctionName(t *testing.T) {
	tests := []struct {
		name string
		nm   string
		want string
	}{
		{name: "full path",
			nm:   "github.com/solsw/pkg.Foo",
			want: "pkg.Foo",
		},
		{name: "generic",
			nm:   "github.com/solsw/pkg.Foo[int,string]",
			want: "pkg.Foo",
		},
		{name: "nested receiver argument with import path",
			nm:   "example.org/pkg.Box[other.org/types.List[int]].Get",
			want: "pkg.Box.Get",
		},
		{name: "no path separator",
			nm:   "pkg.Foo",
			want: "pkg.Foo",
		},
		{name: "no dot",
			nm:   "Foo",
			want: "Foo",
		},
		{name: "empty",
			nm:   "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JustPackageFunctionName(tt.nm)
			if got != tt.want {
				t.Errorf("JustPackageFunctionName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJustFunctionName(t *testing.T) {
	tests := []struct {
		name string
		nm   string
		want string
	}{
		{name: "0",
			nm:   NthCallerName(0),
			want: "NthCallerName",
		},
		{name: "self",
			nm:   NthCallerName(1),
			want: "TestJustFunctionName",
		},
		{name: "Method1",
			nm:   (&Class{}).Method1(),
			want: "(*Class).Method1",
		},
		{name: "Method2",
			nm:   Class{}.Method2(),
			want: "Class.Method2",
		},
		{name: "generic",
			nm:   "github.com/solsw/pkg.Foo[int,string]",
			want: "Foo",
		},
		{name: "no dot",
			nm:   "Foo",
			want: "Foo",
		},
		{name: "no dot generic",
			nm:   "Foo[int]",
			want: "Foo",
		},
		{name: "empty",
			nm:   "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JustFunctionName(tt.nm)
			if got != tt.want {
				t.Errorf("JustFunctionName() = %v, want %v", got, tt.want)
			}
		})
	}
}
