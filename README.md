# gu
Go util

## Install
`go get github.com/arnoluo/gu`

## Usage
### gu 类型方法说明
`gu` 提供一些通用的工具函数

#### 调用方式:
`gu.Func()`

#### Func List:
- `Env(name string) string`: 获取环境变量
- `EnvInt(name string, defaultValue int) int`: 获取环境变量，返回int, 发生错误时返回指定默认值


#### Const List:
- `Ymd`: "2006-01-02"
- `YmdHis`: "2006-01-02 15:04:05"


### gu.At 类型方法说明(AnyType)

`AnyType` 是一个`any`类型的工具类。提供了一些实用的接口相关的方法。

#### 调用方式:
`gu.At.Func()`

#### Func List:
- `Float(fromVal any, toVal *float64) (err error)`: Float 将 any 类型的值转换为 float64 类型的值。 支持以下类型：uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, string（可以是浮点数字符串）。
- `If(isTrue bool, trueValue, falseValue any) any`: If 根据条件判断返回不同的值。
- `InArray(item, stack any) bool`: Return true if stack has the element item, return false otherwise 由于使用了反射，可能在性能上有一定的负担，并且由于没有类型检查，不能有效地避免类型不匹配的情况 不推荐使用这种方式，请转换为具体类型后执行类型下的InArray
- `Int(fromVal any, toVal *int) error`: Int 将 any 类型的值转换为 int 类型的值。 支持以下类型：uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, string（可以是 10 进制数字字符串）。
- `Int64(fromVal any, toValue *int64) (err error)`: Int64 将 any 类型的值转换为 int64 类型的值。 支持以下类型：uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, string（可以是 10 进制数字字符串）。 注：float 类型转换时将损失精度，所以请传入 1111.0 这样不会损失精度的值
- `Int64Array(arr []any, dstArr *[]int64) error`: any 数组转为 int64 数组
- `StructTo(src, dst any) error`: 结构转换, 使用src内所有kv关系，对dst进行赋值 src: struct dst: struct pointer
- `Uint(fromVal any, toVal *uint) error`: Uint 将 any 类型的值转换为 uint 类型的值。 支持以下类型：uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, string（可以是 10 进制数字字符串）。
- `Uint64(fromVal any, toValue *uint64) (err error)`: Uint64 将 any 类型的值转换为 uint64 类型的值。 支持以下类型：uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, string（可以是 10进制数字字符串）。
- `Uint64Array(arr []any, dstArr *[]uint64) error`: any 数组转为 uint64 数组



### gu.Bt 类型方法说明(BoolType)

`BoolType` 是一个`bool`类型的工具类。提供了一些实用的布尔值相关的方法。

#### 调用方式:
`gu.Bt.Func()`

#### Func List:
- `And(values ...bool) bool`: And 对 bool 类型进行逻辑与操作。
- `If(isTrue, trueValue, falseValue bool) bool`: If 根据条件判断返回不同的值。
- `Not(value bool) bool`: Not 对 bool 类型进行逻辑非操作。
- `Or(values ...bool) bool`: Or 对 bool 类型进行逻辑或操作。



### gu.Ft 类型方法说明(FloatType)

`FloatType` 是一个`float64`类型的工具类。提供了一些实用的浮点数相关的方法。

#### 调用方式:
`gu.Ft.Func()`

#### Func List:
- `Abs(value float64) float64`: 返回浮点数的绝对值。
- `Ceil(value float64) float64`: 返回不小于 value 的最小整数，也就是向上取整。
- `Find(value float64, arr []float64) int`: 查找数组。如果数组长度超过规定值，使用二分查找，否则使用遍历查找。成功时返回查找到的数组下标，失败返回 -1。注意，若传入未排序的数组，结果可能并不符合预期。
- `FindSorted(value float64, arr []float64, isAsc bool) int`: 查找已排序数组。如果数组长度超过规定值，使用二分查找，否则使用遍历查找。成功时返回查找到的数组下标，失败返回 -1。
- `Floor(value float64) float64`: 返回不大于 value 的最大整数，也就是向下取整。
- `If(isTrue bool, trueValue, falseValue float64) float64`: 根据条件判断返回不同的值。
- `InArray(value float64, arr []float64) bool`: 在切片中查找浮点数，可传入 threshold 参数，表示最大可容忍的误差范围。
- `InSortedArray(value float64, arr []float64, isAsc bool) bool`: 在已排序的数组中查找浮点数，可传入 threshold 参数，表示最大可容忍的误差范围。
- `LoopFind(value float64, arr []float64) int`: 遍历查找数组。如果数组长度较长或对同一数组做多次 `LoopFind`，建议先 `ArrayAsc` 后使用 `BinFind`。成功时返回查找到的数组下标，失败返回 -1。
- `Max(values ...float64) float64`: 获取浮点数数组中的最大值。
- `Min(values ...float64) float64`: 获取浮点数数组中的最小值。
- `Round(value float64, places int) float64`: 对浮点数进行四舍五入操作，places 参数表示小数点后保留位数。
- `SortAndBinSearch(value float64, arr []float64) int`: 对数组进行排序并进行二分查找法，成功返回查找到的数组下标，失败返回 -1。
- `Str(value float64) string`: 将浮点类型的值转换为字符串类型的值。
- `Sum(values ...float64) float64`: 计算浮点数数组中所有值的总和。



### gu.It 类型方法说明(IntType)

`IntType` 是一个`int64`类型的工具类。提供了一些实用的整型相关的方法。

#### 调用方式:
`gu.It.Func()`

#### Func List:
- `Abs(value int64) int64`: Abs：获取 int64 类型的值的绝对值。
- `ArrayAsc(arr []int64)`: 数组排序 asc
- `ArrayDesc(arr []int64)`: 数组排序 desc
- `Avg(values ...int64) float64`: Avg：计算 int64 类型的值数组中所有值的平均值。
- `BinFind(value int64, arr []int64, isAsc bool) int`: 二分查找数组 此函数适用于数组值已排序或将对同一数组进行多次查找，传入已排序的数组以提高效率。 注：若传入未排序的数组，结果可能并不符合预期。成功时返回查找到的数组下标，失败返回 -1
- `BinSearch(value int64, arr []int64, isAsc bool) int`: 二分查找(此函数不会进行排序，请输入已排序的数组，否则会产生非预期结果) 成功返回查找到的数组下标，失败返回 -1
- `ConvertTo(from int64, toVal any) error`:  将 int64 类型转换为 其他int 类型， 仅限int类型内转换 toVal 必须为 int相关类型的指针
- `Find(value int64, arr []int64) int`: 查找数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找 二分查找时，将会对数组进行升序排序，查找成功返回的也会是升序后的下标 成功时返回查找到的数组下标，失败返回 -1
- `FindSorted(value int64, arr []int64, isAsc bool) int`: 查找已排序数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找 成功时返回查找到的数组下标，失败返回 -1
- `If(isTrue bool, trueValue, falseValue int64) int64`: If 根据条件判断返回不同的值。
- `InArray(value int64, arr []int64) bool`: 整数切片查找，threshold 参数生效
- `InInt16Range(val int64) bool`: 判断给定的整数是否在 int16 范围内
- `InInt32Range(val int64) bool`: 判断给定的整数是否在 int32 范围内
- `InInt8Range(val int64) bool`: 判断给定的整数是否在 int8 范围内
- `InIntRange(val int64) bool`: 判断给定的整数是否在 int 范围内
- `InRange(value, min, max int64) bool`: 实现判断指定的 int64 类型值是否在指定的范围内的功能，返回布尔值
- `InSortedArray(value int64, arr []int64, isAsc bool) bool`: 已排序数组查找，threshold 参数生效
- `InUint16Range(val int64) bool`: 判断给定的整数是否在 uint16 范围内
- `InUint32Range(val int64) bool`: 判断给定的整数是否在 uint32 范围内
- `InUint64Range(val int64) bool`: 判断给定的整数是否在 uint64 范围内
- `InUint8Range(val int64) bool`: 判断给定的整数是否在 uint8 范围内
- `InUintRange(val int64) bool`: 判断给定的整数是否在 uint 范围内
- `IsEven(value int64) bool`: 实现判断指定的 int64 类型值是否为偶数的功能，返回布尔值
- `IsOdd(value int64) bool`: 实现判断指定的 int64 类型值是否为奇数的功能，返回布尔值
- `LoopFind(value int64, arr []int64) int`: 遍历查找数组 如果数组长度较长或对同一数组做多次 LoopFind，建议先 ArrayAsc 后使用 BinFind 成功时返回查找到的数组下标，失败返回 -1
- `Max(values ...int64) int64`: Max：获取 int64 类型的值数组中的最大值。
- `Min(values ...int64) int64`: Min：获取 int64 类型的值数组中的最小值。
- `SortAndBinSearch(value int64, arr []int64) int`: 对数组排序并进行二分查找法，成功返回查找到的数组下标，失败返回 -1
- `Str(value int64) string`: Str：将 int64 类型的值转换为字符串类型的值。
- `Sum(values ...int64) int64`: Sum：计算 int64 类型的值数组中所有值的总和。




### gu.St 类型方法说明(StrType)

`StrType` 是一个`string`类型的工具类。提供了一些实用的字符串相关的方法。

#### 调用方式:
`gu.St.Func()`

#### Func List:
- `ArrayAsc(arr []string)`: 数组正序
- `ArrayDesc(arr []string)`: 数组倒序
- `BinFind(value string, arr []string, isAsc bool) int`: 二分查找数组 此函数适用于数组值已排序或将对同一数组进行多次查找，传入已排序的数组以提高效率。 注：若传入未排序的数组，结果可能并不符合预期。成功时返回查找到的数组下标，失败返回 -1
- `BinSearch(value string, arr []string, isAsc bool) int`: 二分查找(此函数不会进行排序，请输入已排序的数组，否则会产生非预期结果) 成功返回查找到的数组下标，失败返回 -1
- `Bool(str string, errBool bool) bool`: string 转 bool, 需设置转换错误时的默认值
- `Find(value string, arr []string) int`: 查找数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找 二分查找时，将会对数组进行升序排序，查找成功返回的也会是升序后的下标 成功时返回查找到的数组下标，失败返回 -1
- `FindSorted(value string, arr []string, isAsc bool) int`: 查找已排序数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找 成功时返回查找到的数组下标，失败返回 -1
- `Float(str string, errValue float64) float64`: Covert string to positive int using strconv.Atoi(), return errValue if err != nil or value <= 0
- `Format(format string, a ...any) string`: 格式化字符串，类似 fmt.Sprintf()
- `HasPrefix(str, prefix string) bool`: 判断字符串 str 是否以 prefix 开头
- `HasSuffix(str, suffix string) bool`: 判断字符串 str 是否以 suffix 结尾
- `If(boolValue bool, trueValue, falseValue string) string`: Return string param trueValue if boolValue=true, return string param falseValue otherwise
- `InArray(value string, arr []string) bool`: 字符串切片查找，threshold 参数生效
- `InSortedArray(value string, arr []string, isAsc bool) bool`: 已排序数组查找，threshold 参数生效
- `Index(str, substr string) int`: 查找字符串 substr 在 str 中首次出现的位置，如果找不到返回 -1
- `Int(str string, errValue int) int`: string 转 int, 需设置转换错误时的默认值
- `IsEmpty(value string) bool`:
- `IsInt(str string) bool`:
- `IsNum(str string) bool`:
- `Join(strs []string, sep string) string`: 以 sep 为分隔符拼接字符串数组为一个字符串，同 strings.Join()
- `Len(str string) int`: 返回字符串的长度
- `LoopFind(value string, arr []string) int`: 遍历查找数组 如果数组长度较长或对同一数组做多次 LoopFind，建议先 ArrayAsc 后使用 BinFind 成功时返回查找到的数组下标，失败返回 -1
- `LowerFirst(str string) string`: 将字符串首字母小写
- `Ltrim(str, charsets string) string`: 将字符串左侧指定字符集合 charsets 中的字符去除
- `Pint(str string, errValue uint) uint`: Covert string to positive int using strconv.Atoi(), return errValue if err != nil or value <= 0
- `Rand(length int) string`: Generate random str base on letter&number mixed chars
- `RandChars(chars string, length int) (str string)`: Generate random str base on baseChars
- `RandLetters(length int) string`: Generate random str base on letter chars
- `RandNumbers(length int) string`: Generate random str base on number chars
- `RegReplace(baseStr, regexpPattern, replacement string) string`: Replace baseStr with regexpPattern to replacement
- `Replace(str, old, new string, n int) string`: 替换字符串中的 old 为 new，n 为替换的最大次数（小于 0 表示全部替换）
- `Rtrim(str, charsets string) string`: 将字符串右侧指定字符集合 charsets 中的字符去除
- `SortAndBinSearch(value string, arr []string) int`: 对数组排序并进行二分查找法，成功返回查找到的数组下标，失败返回 -1
- `Split(str, sep string) []string`: 将字符串 str 照sep进行分割，并返回分割后的字符串数组，同 strings.Split()
- `Sub(str string, begin, length int) string`: utf8(6 bytes at most) substring
- `Trim(str, charsets string) string`: 将字符串左右两侧指定字符集合 charsets 中的字符去除
- `TrimSpace(str string) string`: 将字符串首尾的空白字符去除
- `Uint(str string, errValue uint) uint`: Covert string to unsigned int using strconv.Atoi(), return errValue if err != nil or value < 0
- `UpperFirst(str string) string`: 将字符串首字母大写



### gu.Ut 类型方法说明(UintType)

`UintType` 是一个`uint64`类型的工具类。提供了一些实用的无符号整型相关的方法。

#### 调用方式:
`gu.Ut.Func()`

#### Func List:
- `ArrayAsc(arr []uint64)`: 数组排序 asc
- `ArrayDesc(arr []uint64)`: 数组排序 desc
- `Avg(values ...uint64) float64`: Avg：计算 uint64 类型的值数组中所有值的平均值。
- `BinFind(value uint64, arr []uint64, isAsc bool) int`: 二分查找数组 此函数适用于数组值已排序或将对同一数组进行多次查找，传入已排序的数组以提高效率。 注：若传入未排序的数组，结果可能并不符合预期。成功时返回查找到的数组下标，失败返回 -1
- `BinSearch(value uint64, arr []uint64, isAsc bool) int`: 二分查找(此函数不会进行排序，请输入已排序的数组，否则会产生非预期结果) 成功返回查找到的数组下标，失败返回 -1
- `ConvertTo(from uint64, toValue any) error`: 将 uint64 类型转换为 其他int 类型， 仅限int类型内转换 toValue 必须为 int相关类型的指针
- `Find(value uint64, arr []uint64) int`: 查找数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找 二分查找时，将会对数组进行升序排序，查找成功返回的也会是升序后的下标 成功时返回查找到的数组下标，失败返回 -1
- `FindSorted(value uint64, arr []uint64, isAsc bool) int`: 查找已排序数组 如果数组长度超过规定值，使用二分查找，否则使用遍历查找 成功时返回查找到的数组下标，失败返回 -1
- `If(isTrue bool, trueValue, falseValue uint64) uint64`: If 根据条件判断返回不同的值。
- `InArray(value uint64, arr []uint64) bool`: 整数切片查找，threshold 参数生效
- `InRange(value, min, max uint64) bool`: 实现判断指定的 uint64 类型值是否在指定的范围内的功能，返回布尔值
- `InSortedArray(value uint64, arr []uint64, isAsc bool) bool`: 已排序数组查找，threshold 参数生效
- `IsEven(value uint64) bool`: 实现判断指定的 uint64 类型值是否为偶数的功能，返回布尔值
- `IsOdd(value uint64) bool`: 实现判断指定的 uint64 类型值是否为奇数的功能，返回布尔值
- `LoopFind(value uint64, arr []uint64) int`: 遍历查找数组 如果数组长度较长或对同一数组做多次 LoopFind，建议先 ArrayAsc 后使用 BinFind 成功时返回查找到的数组下标，失败返回 -1
- `Max(values ...uint64) uint64`: Max：获取 uint64 类型的值数组中的最大值。
- `Min(values ...uint64) uint64`: Min：获取 uint64 类型的值数组中的最小值。
- `SortAndBinSearch(value uint64, arr []uint64) int`: 对数组排序(不会改变原数组顺序)并进行二分查找法，成功返回查找到的数组下标，失败返回 -1
- `Str(value uint64) string`: 将 uint64 类型的数据转化为字符串类型
- `Sum(values ...uint64) uint64`: Sum：计算 uint64 类型的值数组中所有值的总和。
