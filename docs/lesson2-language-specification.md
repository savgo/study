# GO语言基础

## 内置基本数据类型

### 布尔类型

> 布尔值只有`true`和`false`

```go
var v1 bool
```

### 整形 IntegerType

```go
  var v2 uint8   // Range: 0 through 255.
  var v3 uint16  // Range: 0 through 65535.
  var v4 uint32  // Range: 0 through 4294967295.
  var v5 uint64  // Range: 0 through 18446744073709551615.
  var v6 int8    // Range: -128 through 127.
  var v7 int16   // Range: -32768 through 32767.
  var v8 int32   // Range: -2147483648 through 2147483647.
  var v9 int64   // Range: -9223372036854775808 through 9223372036854775807.

  var v10 int    // 平台相关有符号整形, 至少32位, 在64位系统上为64位
  var v11 uint   // 平台相关无符号整形, 至少32位, 在64位系统上为64位

  var v12 byte   // 字节是uint8类型的别名, 可以用来代表单个字节的值

  var v19 rune   // int32类型的别名, 可以用来代表4个字节的值
```

> 问题

  - [ ] int和uint 用在什么场合
  - [ ] byte 用途
  - [ ] rune 用途

### 浮点类型 FloatType

```go
  var v13 float32  // 32位浮点数
  var v14 float64  // 64位浮点数
```

### 复数类型 ComplexType

> 备注

  1. 参考C语言中的Complex复数类型和imaginary虚数类型

```go
  var v15 complex64  // 32位浮点数+虚数组成
  var v16 complex128 // 64位浮点数+虚数组成
```

### 字符串类型 string

> 备注

  1. 字符串是8字节队列, 可以为空, 不能为nil, 队列元素不可显式更改(没有下标)

```go
  var v17 string
```

### 指针类型

```go
  var v18 uintptr
```

### interface类型

```go
  var v20 interface{}
```

### error类型

> 备注

  1. error类型 是interface类型的一个特殊定义

```go
  var v21 error
```

### 数组类型 slice

```go
  var v22 []int
```

### 图类型 map

```go
  var v23 map[string]int
```

### channel管道类型

```go
  var v24 chan <- int
```

## 内置常量和变量

### 布尔常量 true 和 false

> 备注

  1. bool类型就是 true 和 false 两个常量值的集合

  2. 官方定义是常量, 不是变量, 也不是字面量

> 问题

  1. [ ] 什么是字面量?

### iota常量

> 备注

  1. 相当于初始字面量 0
  2. iota只能在常量的表达式中使用

```go
  const initV = iota
```

> 参考
 - [golang 使用 iota](https://studygolang.com/articles/2192)

### 空值nil

> 备注

  1. 空值nil在go中是作为变量存在的, 官方定义为`var nil Type`

  2. 空值nil是作为变量值存在的, 除非特别指定，否则无法使用 nil 对变量赋值
  
  3. 空值nil, 可以作为以下数据类型的值 pointer, channel, func, interface, map, slice

```go
  // 声明时 必须指定类型
  var i *int = nil
  var err error = nil
  var arr []int = nil
  var m1 map[string]int = nil
  var c1 chan <- string = nil
  // 不支持的情况
  // var x = nil // @ERROR 没有特别指定类型, 无法识别 use of untyped nil
  // var x string = nil // @ERROR 字符串不能为nil  cannot use nil as type string in assignment
  // var x uintptr = nil // @TODO useage, cannot use nil as type uintptr in assignment
  // var x func = nil // @TODO useage, 语法错误
```

## 变量的声明

### 单个变量的声明

```go
  var v1 int     // 1.1 声明一个变量, 使用默认值
  var v2 int = 2 // 1.2 声明并初始化赋值
  var v3 = "s"   // 1.3 声明一个变量, 自动推导数据类型
```

### 声明多个变量

```go
  var m1, m2 int        // 2.1 声明多个同数据类型的变量
  var m3, m4 int = 2, 8 // 2.2 声明多个同数据类型的变量, 并初始化赋值
  var m5, m6 = "s", 8   // 2.3 声明多个自动推导类型的变量, 必须初始化赋值

  // var x1 int, x2 int // @ERROR 不允许显式声明一组不同数据类型的变量
```

### 因式分解法声明变量

> 备注

  1. 声明多个变量时, 需要使用分号或换行分割

```go
  var (n7 int)          // 3.1 声明单个变量
  var (n8 int = 1)      // 3.2 声明单个变量并赋值
  var (n9 = 1)          // 3.3 声明单个推导变量
  var (n10 int; n11 float32)    // 3.4 声明多个变量 , 需要使用分号或换行分割
  var (n12 int = 3; n13 float32 = .8)   // 3.5 声明多个变量, 并初始化赋值
  var (n14 = 3; n15 = .8)       // 3.6 声明多个推导变量
```

### 推导声明

> 备注

  1. 只能在函数体中出现
  2. Go语言从不做隐式转换

```go
  p1 := 5           // 4.1 单个推导声明赋值, 不能指定
  p2, p3 := 3, 8    // 4.2 多个推导声明
```

> 问题

  1. [ ] 推导类型是什么数据类型(整形是int, 其他未测试)

## 常量的声明

> 备注

  1. 常量的声明必须初始化赋值
  2. 常量不能推导声明

### 单个常量的声明

```go
  const v2 int = 2 // 1.1 声明并初始化赋值
  const v3 = "s"   // 1.2 声明一个常量, 自动推导数据类型

  // const x // @ERROR 不允许这样写的, 因为无法推导数据类型
```

### 声明多个常量

```go
  const m3, m4 int = 2, 8   // 2.1 声明多个同数据类型的常量, 并初始化赋值
  const m5, m6 = "s", 8     // 2.2 声明多个自动推导类型的常量, 必须初始化赋值
```

### 因式分解法声明常量

```go
  const (n8 int = 1)    // 3.1 声明单个常量并赋值
  const (n9 = 1)        // 3.2 声明单个推导常量
  const (n12 int = 3; n13 float32 = .8)     // 3.3 声明多个常量, 并初始化赋值
  const (n14 = 3; n15 = .8)     // 3.4 声明多个推导常量
```

## 控制结构

### if 语句

> 语法

`if [init]; condation {`

`[} else [if [init]; condation] {]`


> 备注

  1. if语句 只接收bool类型的常量, 变量或表达式
  2. 关键字和花括号必须在一行, 除非在初始语句换行(不推荐)

```go
  var x int = 10
  var y = false
  if false { // 常量
    // skip
  } else if y { // 变量
    // skip
  } else if x < int(y) { // 表达式
    // skip
  } else if n := x; n < 0 { // 初始语句可以是 推导声明
    // skip
  } else if m, n := 0, 1; // 初始语句换行
  m > n {                 // 多个推导声明
    // skip
  } else if false && true { // 表达式
    // skip
  } else { // 也可以没有 else
    println("if")
  }
```

### switch 语句

> 语法

`switch [condation] {`

`case condation : [body][fallthrough]`

`[default : {]`

> 备注

  1. switch语句设计的比较灵活
  2. switch参数可以是常量, 变量, 表达式, 或省略(参数为bool类型的true值)
  3. case参数可以是常量, 变量, 表达式, 参数类型需要和switch参数值类型一致
  4. switch参数可以是 推导声明语句, 此时case参数类型为bool
  5. case可以带多个参数, 用逗号隔开
  6. switch语句命中则不继续匹配, 除非用`fallthrough`关键字

```go
  var m int = 0
  var n int = -1
  switch m { // switch参数为 int类型变量, case 参数类型相同
    case -1 : // 常量 空的 case 体
    case n : return // 变量 可以用 return 返回函数值
    case n - 1 : println("skip") // 表达式
    case 'a', n - 2 : println("skip") // 可以匹配多项, 用逗号隔开
    case n + 1 : println("m = -1 + 1") // 匹配执行
      fallthrough // 不跳出, 继续匹配
    case 0 : println("m = 1")
    default : // 可省略
  }

  switch 1 { // 常量
    case 0: println("skip")
  }

  switch { // 省略 相当于 switch true
    case false: println("skip")
  }

  switch m + 1 { // 表达式
    case -1: println("skip")
  }

  switch x := m; { // 推导声明语句, case参数为bool
    case x > 0: println("skip")
    case x < 0: println("skip")
  }

  switch x, y := m, n; { // 推导声明语句, case参数为bool
    case x < y : println("skip")
    case x == y : println("skip")
  }
```
