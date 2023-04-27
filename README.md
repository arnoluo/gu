# gu
Go util

## Install
`go get github.com/arnoluo/gu`

## Usage

### common func
```go
import (
    "github.com/arnoluo/gu"
)

// common func
// getenv
isDebug := gu.Env("DEBUG")
```

### IntType(int64) func
```go
// convert string value to int
s123 := gu.It.Str(123) // "123"

// If
i331 := gu.It.If(1 == 2, 332, 331) // 331

// convert to other int types
var existedVal int64 1234
var myInt int8
if err := it.ConvertTo(testVal, &myInt); err != nil {
    myInt = 0
}
```

### UintType(uint64) func
```go
s123 := gu.Ut.Str(123) // "123"
```

### FloatType(float64) func
```go
s123 := gu.Ft.Str(123) // "123"
```

### BoolType(bool) func
```go
val := gu.Bt.If(1 == 2, 1, 2) // 2
```

### StrType(string) func
```go
val := gu.St.Int("123", 0) // 123
```


### AnyType(interface{}) func
```go
// convert string/intN/uintN value to int
var i int
from := "123456"
if err := gu.At.Int(from, &i); err != nil {
    fmt.Println(err)
}
fmt.Println(i)
```