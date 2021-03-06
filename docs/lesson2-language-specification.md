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

* [ ] int和uint 用在什么场合
* [ ] byte 用途
* [ ] rune 用途

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

1. 字符串是8字节队列, 可以为空, 不能为nil, 队列元素不可显式更改\(没有下标\)

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

### slice类型

> 备注

1. 切片, 可以理解为动态数组

```go
  var v22 []int
```

### map类型

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
>
> * [golang 使用 iota](https://studygolang.com/articles/2192)

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

1. [ ] 推导类型是什么数据类型\(整形是int, 其他未测试\)

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
2. 关键字和花括号必须在一行, 除非在初始语句换行\(不推荐\)

```go
  var x int = 10
  var y = false
  if false { // 常量
    // skip
  } else if y { // 变量
    // skip
  } else if x < 0 { // 表达式
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
2. switch参数可以是常量, 变量, 表达式, 或省略\(参数为bool类型的true值\)
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

### for 语句

> 语法

`for [init]; condation; [post] {`

`}`

`for [condation] {`

`}`

> 备注

1. go语言没有while语句
2. for语句可以通过标签方便跳出嵌套循环
3. 请尽可能忘掉标签和goto语句

```go
  // 常规的for循环
  for i:= 0; i < 10; i++ {
    if i < 5 {
      continue
    }
    println("for i", i)
    if i > 4 {
      break
    }
  }
  var i = 0
  for ; i < 2; i++ { // 省略 init
  }
  for j := 0; j < 2; { // 省略 post
    j++
  }
  for ; i < 3; { // 省略 init 和 post
    i++
  }
  for i < 4 { // 只有表达式, 类似 while () {}
    i++
  }
  for { // 无限循环 类似 for(;;) {} 或 while (true){}
    break
  }

  var v = 0;
  // 再for关键字前定义 循环标签 loopFor
  loopFor: for i :=0; i < 5; i++ {
    v++
    for j := 0; j < 5; j++ {
      v++
      for k := 0; k < 5; k++ {
        v++
        if k > j {
          break loopFor // 跳出 父级循环标签, 不要妄想跳到其他地方去
        }
      }
    }
  }
  println("break loopFor:", v)
```

### for range 语句

> 语法

`for condation range condation {`

`}`

> 备注

1. range迭代器可以用在 array, slice, map, string, channel等类型

```go
  // array迭代
  for id, val := range [2]string{"a", "b"} {
    if id < 0 {
      println("skip", id, val)
    }
  }
  // slice迭代
  for id, val := range []string{"a", "b"} {
    if id < 0 {
      println("skip", id, val)
    }
  }
  // map迭代
  for key, val := range map[string]int{"a": 1, "b": 2} {
    if val < 0 {
      println("skip", key, val)
    }
  }
  // string迭代
  for id, char := range "abcd" {
    if id < 0 {
      println("skip", id, char)
    }
  }
  // channel迭代
  c := make(chan int)
  go func (c chan int) { // 简单的协程和闭包
    c <- 1
    c <- 2
    close(c)
  }(c) // 虽然func定义可以用()包裹, 看起来会更像闭包, 何如
  for val := range c {
    println("get:", val)
  }
```

## 运算

### 运算符分类

> 备注

1. `++` 和`--` 在go语言中是语句, 因此不能和运算符混用
2. 字符串相拼接的 `+` 号不是运算符, 而是连接符
3. 不支持三目运算符(只是为了防止深入嵌套?)
4. 不支持运算符重载, 泛型和模板(如此想不优雅都难)

| 分类名称              | 运算符                         | 使用场景                      |
| :---------------- | :-------------------------- | :------------------------ |
|                   |                             |                           |
| 赋值运算符             | `=`                         | 赋值运算                      |
| 算术运算符\(可与赋值运算复合\) | `+` `-` `*` `%`             | 数字类型算术运算                  |
| 位运算符\(可与赋值运算复合\)  | `<<` `>>`  `|` `^` `&`      | 数字类型位运算                   |
| 逻辑运算符             | `!` `&&` `||`               | 布尔类型逻辑运算                  |
| 关系运算符             | `<` `<=` `>` `>=` `==` `!=` | 比较运算, 结果为布尔类型             |
| 指针运算符             | `*` `&`                     |                           |
| 下标运算符             | `[]`                        | string, array, slice, map |
| 管道运算符             | `<-`                        | 只能用在chan中                 |

### 运算符优先级

> 备注

1. 可以使用`()`进行分组提高优先级
2. 运算符对于操作数的数据类型要求严格, 可以减少分组的使用

| 级别   | 运算符                            |      |
| ---- | ------------------------------ | ---- |
| 7    | `^` `!`                        |      |
| 6    | `*` `/` `%` `<<` `>>` `&` `&^` |      |
| 5    | `+` `-` `|` `^`                |      |
| 4    | `==` `!=` `<` `<=` `<=` `>`    |      |
| 3    | `<-`                           |      |
| 2    | `&&`                           |      |
| 1    | `||`                           |      |

### 运算符的使用

```go
  // bool 逻辑清晰的可以不用分组
  println("bool", !true && 3 > 4 || 4 > 3)
  // int 程序员的梦(魇)
  println("int", 1 + 2 * 3 ^ 5 >> 1 & 4 | 7)
  // string 字符串拼接而非运算
  println("string", "str" + "ing")
  // 下标运算
  println("string[]", "str"[0]) // 取出来的值是byte
  println("array[]", [...]int{2, 3}[0]) // 数组可以使用...让编译器判断长度
  println("slice[]", []int{1,2}[0])
  println("map[]", map[int]string{1: "s"}[1])
  // 指针运算
  var n1 int
  var p1 *int = &n1
  var pp1 **int = &p1
  *p1++   // n1++
  **pp1++ // n1++
  println("n1", n1)
  // 管道运算
  ch := make(chan int)
  var send chan <- int = ch // 定义一个发送者
  var recv <- chan int = ch // 定义一个接受者
  go func () {
    send <- 1 // 或使用 ch <- 1
    close(ch)
  }()
  println("cha", <- recv) // 取出, 或使用 <- ch
```

## 保留字

保留字包括关键字(已用)和预留关键字(暂无)

### 关键字

- 包和导入: `package` `import`
- 声明定义: `var` `const`
- 类型定义: `func` `map` `chan` `interface` `struct` `type`
- 控制结构: `if` `else` `for` `range`  `switch` `case` `default` `fallthrough` 
- 控制跳转: `break`  `continue` `goto` `return`
- 协程异步: `go` `defer` `select`

### 关键字用例

```go
/*
 * 关键字使用例子
 */
/**
 * 包定义, 可执行程序包名必须是main
 * @keyword type
 */
package main
/**
 * 引入依赖包
 * @keyword import
 */
import (
  "fmt" // 打印
  "math/rand" // 生成随机数
  "time" // 用于生成不同的随机数
)
/**
 * 定义性别类型(仿枚举)
 * go 没有枚举类型, 可以使用type定义sex数据类型, 真实类型为uint8
 * @keyword type
 */
type sex uint8
/**
 * 性别定义
 * @keyword const
 * @type {sex}
 */
const (
  _ = iota // 空白符(blank identifier), 占位作用, 不能作为常量或变量使用
  male
  female
)
/**
 * 性别本地化, 用一个包内全局map变量来取数据
 * @keyword var map
 * @type {map[sex]string}
 */
var sexNames = map[sex]string{
  male: "男",
  female: "女",
}
/**
 * 定义用户结构
 * @keyword struct
 */
type user struct {
  name string
  sex sex // 使用我们定义的sex枚举类型
  age uint32
}
/**
 * 定义人类接口
 * @keyword interface
 */
type Person interface {
  Speek ()
  profile () user
}
/**
 * 给user加上说话功能
 * @keyword func if else
 */
func (user user) Speek () {
  str := fmt.Sprintf("大家好, 我叫 %s %s", user.name, sexNames[user.sex])
  if user.age > 0 {
    str += fmt.Sprintf(", 今年 %d 岁了.\n", user.age)
  } else {
    str += "\n"
  }
  fmt.Printf(str)
}
/**
 * 获取用户信息
 * @keyword return
 */
func (user user) profile () user{
  return user
}
// 定义一个课堂房间
type ClassRoom struct {
  persons []Person
}

/**
 * 开始上课
 * @keyword select case range switch fallthrough default break continue chan
 */
func (room * ClassRoom) Start (ready <- chan bool) {
  // <- ready //取出丢弃, 可简写
  select {
    case <- ready:
    // case ok := <- ready: // 不丢弃可继续处理
  }
  fmt.Println("开始上课", "同学们请发言")
  var users []user = make([]user, len(room.persons))
  for index, person := range room.persons {
    person.Speek()
    users[index] = person.profile()
  }
  var maleCount, femaleCount, unknownSex int
  var ageNum, ageTotal uint32
  for i := 0; i < len(users); i++ {
    user := users[i]
    switch (user.sex) {
      case male: maleCount++
      case female: femaleCount++
      case 0: fallthrough
      default: unknownSex++
        break
    }
    if user.age < 10 || user.age > 90 {
      continue
    }
    ageTotal += user.age
    ageNum ++
  }
  fmt.Printf("性别统计: 男同学%d, 女同学%d, 未知性别:%d\n",
    maleCount, femaleCount, unknownSex)
  if ageNum > 0 {
    fmt.Printf("平均年龄(10-90): %d\n", ageTotal / uint32(ageNum))
  }
}
func (room *ClassRoom) addPerson (person Person) {
  room.persons = append(room.persons, person)
}
func (room *ClassRoom) CreatePersion (name string, sex sex, age uint32) {
  room.addPerson(Person(user{name, sex, age}))
}
func (room *ClassRoom) CreateSome (names []string){
  for _, name := range names {
    room.CreatePersion(name, 0, 0)
  }
}
func (room * ClassRoom) Close() {
  room.persons = room.persons[0: 0]
  fmt.Printf("课堂关闭")
}
func NewClassRoom () (room *ClassRoom) {
  room = &ClassRoom{}
  return
}
/**
 * 一个main包相当于一个应用, 只能有一个main函数
 * @keyword defer go
 */
func main () {
  var room *ClassRoom
  { // 证明 defer的作用域是函数
    room = NewClassRoom()
    defer room.Close()
  }
  ready := make(chan bool)
  go func () {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    room.CreatePersion("张三", 0, uint32(r.Intn(100)))
    room.CreatePersion("李磊", male, uint32(r.Intn(100)))
    room.CreatePersion("韩梅梅", female, uint32(r.Intn(100)))
    room.CreatePersion("lucy", female, uint32(r.Intn(100)))
    room.CreateSome([]string{"蝙蝠侠", "本山", "大婶"})
    ready <- true
    close(ready)
  }()
  room.Start(ready)
}

```







