# runtimehelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/runtimehelper.svg)](https://pkg.go.dev/github.com/solsw/runtimehelper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/runtimehelper)

Helpers for Go's [`runtime`](https://pkg.go.dev/runtime) package.

## Install

```sh
go get github.com/solsw/runtimehelper
```

```go
import "github.com/solsw/runtimehelper"
```

## Overview

The package provides two groups of helpers:

- **Caller names** — retrieve the name of a function somewhere up the call stack.
- **Name trimming** — reduce a full function name (as returned by [`runtime.Func.Name`](https://pkg.go.dev/runtime#Func.Name)) to just the package-qualified name or the bare function name.

## API

### Caller names

#### `NthCallerName(n int) string`

Returns the full name of the function at depth `n`. If `main.main` calls `main.work`, the following calls inside `work` identify these functions.

| Call inside `work` | Function identified |
| --- | --- |
| `NthCallerName(0)` | `runtimehelper.NthCallerName` (including its import path) |
| `NthCallerName(1)` or `CallerName()` | `main.work` |
| `NthCallerName(2)` or `CallerCallerName()` | `main.main` |

A negative depth or a depth beyond the available stack frames returns `""`. This includes `-1` and `math.MaxInt`. All three caller helpers also return `""` if the runtime cannot resolve the function name.

#### `CallerName() string`

Returns the name of the function that called `CallerName`.

```go
func work() string {
	return runtimehelper.CallerName() // "main.work" in package main
}
```

#### `CallerCallerName() string`

Returns the name of the function that called the function that called `CallerCallerName` (one frame further up than `CallerName`).

> `CallerName` and `CallerCallerName` are marked `//go:noinline` so they always occupy their own stack frame, keeping the fixed skip count passed to `NthCallerName` correct.

### Name trimming

Both helpers accept a full function name from [`runtime.Func.Name`](https://pkg.go.dev/runtime#Func.Name). For a generic receiver such as `*Box[int]`, the runtime can report the method as `pkg.(*Box[...]).Get`. The literal `[...]` means the type arguments were omitted from the name. Both helpers remove that bracketed portion and keep the receiver and method. They also handle explicit type arguments, including nested types and import paths.

#### `JustPackageFunctionName(funcName string) string`

Returns just the function name preceded by the package name. Generic type arguments (e.g. `"[int]"`) are stripped.

| Input | Output |
| --- | --- |
| `github.com/solsw/pkg.Foo` | `pkg.Foo` |
| `github.com/solsw/pkg.Foo[int,string]` | `pkg.Foo` |
| `github.com/solsw/pkg.(*Box[...]).Get` | `pkg.(*Box).Get` |
| `pkg.Foo` | `pkg.Foo` |
| `Foo` | `Foo` |
| `""` | `""` |

#### `JustFunctionName(funcName string) string`

Returns just the function name. If `funcName` contains no package separator, it is returned (with generic type arguments stripped) unchanged.

| Input | Output |
| --- | --- |
| `github.com/solsw/pkg.Foo` | `Foo` |
| `github.com/solsw/pkg.Foo[int,string]` | `Foo` |
| `github.com/solsw/pkg.(*Box[...]).Get` | `(*Box).Get` |
| `Foo` | `Foo` |
| `Foo[int]` | `Foo` |
| `""` | `""` |

## Example

```go
package main

import (
	"fmt"

	"github.com/solsw/runtimehelper"
)

type Box[T any] struct{}

func (*Box[T]) Get() {
	full := runtimehelper.CallerName()
	fmt.Println(full)
	fmt.Println(runtimehelper.JustPackageFunctionName(full))
	fmt.Println(runtimehelper.JustFunctionName(full))
	fmt.Println(runtimehelper.CallerCallerName())
}

func main() {
	new(Box[int]).Get()
}
```

Output

```text
main.(*Box[...]).Get
main.(*Box).Get
(*Box).Get
main.main
```
