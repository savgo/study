package main

func noop() {

}

func builtinType() {
  println("内置基本数据类型---builtinType---")
  // 1 变量声明
  // var 变量名 变量类型

  // 2 内置基本数据类型
  
  // 2.1 布尔类型, true and false.
  println("布尔类型")
  var v1 bool
  println("bool", v1)

  // 2.2 整形 IntegerType
  println("整形")
  // Range: 0 through 255.
  var v2 uint8
  println("uint8", v2)
  // Range: 0 through 65535.
  var v3 uint16
  println("uint16", v3)
  // Range: 0 through 4294967295.
  var v4 uint32
  println("uint32", v4)
  // Range: 0 through 18446744073709551615.
  var v5 uint64
  println("uint64", v5)
  // Range: -128 through 127.
  var v6 int8
  println("int8", v6)
  // Range: -32768 through 32767.
  var v7 int16
  println("int16", v7)
  // Range: -2147483648 through 2147483647.
  var v8 int32
  println("int32", v8)
  // Range: -9223372036854775808 through 9223372036854775807.
  var v9 int64
  println("int64", v9)
  // 平台相关有符号整形, 至少32位, 在64位系统上为64位
  var v10 int
  println("int", v10)
  // 平台相关无符号整形, 至少32位, 在64位系统上为64位
  var v11 uint
  println("uint", v11)
  // 字节是uint8类型的别名, 可以用来代表单个字节的值
  var v12 byte
  println("byte", v12)
  // int32类型的别名, 可以用来代表4个字节的值
  var v19 rune
  println("rune", v19)

  // 2.3 浮点类型 FloatType
  println("浮点类型")
  // 32位浮点数
  var v13 float32
  println("float32", v13)
  // 64位浮点数
  var v14 float64
  println("float64", v14)
  
  // 2.4 复数类型 ComplexType (参考C语言中的Complex复数类型和imaginary虚数类型)
  println("复数类型")
  // 32位复数
  var v15 complex64
  println("complex64", v15)
  // 64位复数
  var v16 complex128
  println("complex128", v16)

  // 2.5 字符串类型 string
  // 字符串是8字节队列, 可以为空, 不能为nil, 队列元素不可显式更改(没有下标)
  println("字符串类型")
  var v17 string
  println("string", v17)
  // 2.6 指针类型
  println("指针类型")
  var v18 uintptr
  println("uintptr", v18)

  // 2.7 interface类型
  println("interface类型")
  var v20 interface{}
  println("interface", v20)

  // 2.8 error类型 是interface类型的一个特殊定义
  println("error类型")
  var v21 error
  println("error", v21)

  // 2.9 slice数组类型
  println("slice数组类型")
  var v22 []int
  println("slice", v22)

  // 2.10 map图类型
  println("map图类型")
  var v23 map[string]int
  println("map", v23)

  // 2.11 channel管道类型
  println("channel管道类型")
  var v24 chan <- int
  println("chan", v24)
}

func builtinValue () {
  println("内置常量和变量---builtinValue---")
  // bool类型就是 true 和 false 两个常量的集合
  println("布尔常量", "true", true)
  println("布尔常量", "false", false)

  // iota常量, iota只能在常量的表达式中使用
  // println(iota) // undefined: iota
  // var x = iota // undefined: iota
  const initV = iota
  println("初始常量", "iota", initV)
  // @USEAGE golang 使用 iota
  // @SEE https://studygolang.com/articles/2192
  
  // 空值nil是作为变量值存在的, 除非特别指定，否则无法使用 nil 对变量赋值
  // 空值nil, 可以作为以下数据类型的值 pointer, channel, func, interface, map, slice
  // var x string = nil // @ERROR 字符串不能为nil  cannot use nil as type string in assignment
  // var x = nil // @ERROR 没有特别指定类型, 无法识别 use of untyped nil
  // var x uintptr = nil // @TODO useage, cannot use nil as type uintptr in assignment
  // var x func = nil // @TODO useage, 语法错误
  var null interface{} = nil
  println("空值变量", "interface", null)

  // error类型
  var err error = nil
  println("空值变量", "error", err)

  var arr []int = nil
  println("空值变量", "slice", arr)

  var m1 map[string]int = nil
  println("空值变量", "map", m1)

  var c1 chan <- string = nil
  println("空值变量", "chan", c1)
}

func declareVar() {
  println("变量的声明")
  // 1. 声明变量
  var v1 int // 1.1 声明一个变量, 使用默认值
  println("声明变量", "int", v1)
  var v2 int = 2 // 1.2 声明并初始化赋值
  println("声明变量", "int", v2)
  var v3 = "s" // 1.3 声明一个变量, 自动推导数据类型
  // var x // @ERROR 不允许这样写的, 因为无法推导数据类型
  println("声明变量", "auto", v3)
  var v4 = 5 // @QUESTING v4 如何检测是什么数据类型?
  println("声明变量", "auto", v4)
  // 2. 声明多个变量
  var m1, m2 int // 2.1 声明多个同数据类型的变量
  println("声明变量", "m1", m1, "m2", m2)
  // var x1 int, x2 int // @ERROR 不允许显式声明一组不同数据类型的变量
  var m3, m4 int = 2, 8 // 2.2 声明多个同数据类型的变量, 并初始化赋值
  println("声明变量", "m3", m3, "m4", m4)
  var m5, m6 = "s", 8 // 2.3 声明多个自动推导类型的变量, 必须初始化赋值
  println("声明变量", "m5", m5, "m6", m6)
  // 3. 因式分解法声明变量
  var (n7 int) // 3.1 声明单个变量
  println("声明变量", "n7", n7)
  var (n8 int = 1) // 3.2 声明单个变量并赋值
  println("声明变量", "n8", n8)
  var (n9 = 1) // 3.3 声明单个推导变量
  println("声明变量", "n9", n9)
  var (n10 int; n11 float32) // 3.4 声明多个变量 , 需要使用分号或换行分割
  println("声明变量", "n10", n10, "n11", n11)
  var (n12 int = 3; n13 float32 = .8) // 3.5 声明多个变量, 并初始化赋值
  println("声明变量", "n12", n12, "n13", n13)
  var (n14 = 3; n15 = .8) // 3.6 声明多个推导变量
  println("声明变量", "n14", n14, "n15", n15)
  // 4. 推导声明赋值(或者叫auto声明) @NOTE 只能在函数体中出现
  p1 := 5 // 4.1 单个推导声明赋值, 不能指定类型
  println("声明变量", "p1", p1)
  p2, p3 := 3, 8 // 4.2 多个推导声明赋值
  println("声明变量", "p2", p2, "p3", p3)
}

func declareConst() {
  println("常量的声明")
  // @NOTE 常量的声明必须初始化赋值
  // @NOTE 常量不能推导声明赋值
  // 1. 声明常量
  const v2 int = 2 // 1.1 声明并初始化赋值
  println("声明常量", "int", v2)
  const v3 = "s" // 1.2 声明一个常量, 自动推导数据类型
  // const x // @ERROR 不允许这样写的, 因为无法推导数据类型
  println("声明常量", "auto", v3)
  const v4 = 5 // @QUESTING v4 如何检测是什么数据类型?
  println("声明常量", "auto", v4)
  // 2. 声明多个常量
  const m3, m4 int = 2, 8 // 2.1 声明多个同数据类型的常量, 并初始化赋值
  println("声明常量", "m3", m3, "m4", m4)
  const m5, m6 = "s", 8 // 2.2 声明多个自动推导类型的常量, 必须初始化赋值
  println("声明常量", "m5", m5, "m6", m6)
  // 3. 因式分解法声明常量
  const (n8 int = 1) // 3.1 声明单个常量并赋值
  println("声明常量", "n8", n8)
  const (n9 = 1) // 3.2 声明单个推导常量
  println("声明常量", "n9", n9)
  const (n12 int = 3; n13 float32 = .8) // 3.3 声明多个常量, 并初始化赋值
  println("声明常量", "n12", n12, "n13", n13)
  const (n14 = 3; n15 = .8) // 3.4 声明多个推导常量
  println("声明常量", "n14", n14, "n15", n15)
}

func main() {
  noop() // 函数的最简定义
  builtinType() // 内置基本数据类型
  builtinValue() // 内置常量和变量
  declareVar() // 变量的声明及赋值
  declareConst() // 常量的声明及赋值
  // 运算
  // 
}
