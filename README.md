# go-logic

Logic utilities for humans. Zero dependencies.

## Installation

```shell
go get github.com/bennett-jacob/go-logic/logic
```

## Usage

### All / And

`logic.All()` returns `true` if all of the given arguments are `true`.

```golang
logic.All(false, false, false) // false
logic.All(true, true, false)   // false
logic.All(true, true, true)    // true

// equivalent to
logic.And(false, false, false) // false
logic.And(true, true, false)   // false
logic.And(true, true, true)    // true
```

### None / Nor

`logic.None()` returns `true` if none of the given arguments are `true`.

```golang
logic.None(false, false, false) // true
logic.None(true, true, false)   // false
logic.None(true, true, true)    // false

// equivalent to
logic.Nor(false, false, false) // true
logic.Nor(true, true, false)   // false
logic.Nor(true, true, true)    // false
```

### Any / Or

`logic.Any()` returns `true` if any of the given arguments are `true`.

```golang
logic.Any(false, false, false) // false
logic.Any(true, true, false)   // true
logic.Any(true, true, true)    // true

// equivalent to
logic.Or(false, false, false) // false
logic.Or(true, true, false)   // true
logic.Or(true, true, true)    // true
```

### OnlyOne / Xor

`logic.OnlyOne()` returns `true` if exactly one of the given arguments is `true`.

```golang
logic.OnlyOne(false, false, false) // false
logic.OnlyOne(false, true, false)  // true
logic.OnlyOne(true, true, false)   // false
logic.OnlyOne(true, true, true)    // false

// equivalent to
logic.Xor(false, false, false) // false
logic.Xor(false, true, false)  // true
logic.Xor(true, true, false)   // false
logic.Xor(true, true, true)    // false
```
