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
  // @useage golang 使用 iota
  // @see https://studygolang.com/articles/2192
  
  // 空值nil是作为变量值存在的, 除非特别指定，否则无法使用 nil 对变量赋值
  // 空值nil, 可以作为以下数据类型的值 pointer, channel, func, interface, map, slice
  // var x string = nil // 字符串不能为nil  cannot use nil as type string in assignment
  // var x = nil // 没有特别指定类型, 无法识别 use of untyped nil
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

func main() {
  noop() // 函数的最简定义
  builtinType() // 内置基本数据类型
  builtinValue() // 内置常量和变量
}
