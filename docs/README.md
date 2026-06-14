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

Returns the name of the n-th caller function of `NthCallerName`. On any failure an empty string is returned. The meaning of `n` matches [`runtime.Caller`](https://pkg.go.dev/runtime#Caller): `0` identifies `NthCallerName` itself, `1` its caller, and so on.

```go
func work() string {
	return runtimehelper.NthCallerName(1) // "...work" — work itself (the direct caller)
}
```

`NthCallerName(2)` from inside `work` would instead return the name of whoever called `work`.

#### `CallerName() string`

Returns the name of the function that called `CallerName`.

```go
func work() {
	name := runtimehelper.CallerName() // name of work's caller
	_ = name
}
```

#### `CallerCallerName() string`

Returns the name of the function that called the function that called `CallerCallerName` (one frame further up than `CallerName`).

> `CallerName` and `CallerCallerName` are marked `//go:noinline` so they always occupy their own stack frame, keeping the fixed skip count passed to `NthCallerName` correct.

### Name trimming

These operate on the full function name string returned by [`runtime.Func.Name`](https://pkg.go.dev/runtime#Func.Name), e.g. `"github.com/solsw/pkg.Foo"`.

#### `JustPackageFunctionName(funcName string) string`

Returns just the function name preceded by the package name. Generic type parameters (e.g. `"[int]"`) are stripped.

| Input | Output |
| --- | --- |
| `github.com/solsw/pkg.Foo` | `pkg.Foo` |
| `github.com/solsw/pkg.Foo[int,string]` | `pkg.Foo` |
| `pkg.Foo` | `pkg.Foo` |
| `Foo` | `Foo` |
| `` (empty) | `` (empty) |

#### `JustFunctionName(funcName string) string`

Returns just the function name. If `funcName` contains no package separator, it is returned (with generic type parameters stripped) unchanged.

| Input | Output |
| --- | --- |
| `github.com/solsw/pkg.Foo` | `Foo` |
| `github.com/solsw/pkg.Foo[int,string]` | `Foo` |
| `Foo` | `Foo` |
| `Foo[int]` | `Foo` |
| `` (empty) | `` (empty) |

## Example

```go
package main

import (
	"fmt"

	"github.com/solsw/runtimehelper"
)

func main() {
	full := runtimehelper.NthCallerName(1) // "main.main"
	fmt.Println(runtimehelper.JustPackageFunctionName(full)) // "main.main"
	fmt.Println(runtimehelper.JustFunctionName(full))        // "main"
}
```
