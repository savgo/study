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
	// 32位浮点数+虚数组成
	var v15 complex64
	println("complex64", v15)
	// 64位浮点数+虚数组成
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

	// 2.9 slice切片类型
	println("slice切片类型")
	var v22 []int
	println("slice", v22)

	// 2.10 map图类型
	println("map图类型")
	var v23 map[string]int
	println("map", v23)

	// 2.11 channel管道类型
	println("channel管道类型")
	var v24 chan<- int
	println("chan", v24)
}

func builtinValue() {
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

	var i *int = nil
	println("空值变量", "pointer", i)

	// error类型
	var err error = nil
	println("空值变量", "error", err)

	var arr []int = nil
	println("空值变量", "slice", arr)

	var m1 map[string]int = nil
	println("空值变量", "map", m1)

	var c1 chan<- string = nil
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
	var (
		n7 int
	) // 3.1 声明单个变量
	println("声明变量", "n7", n7)
	var (
		n8 int = 1
	) // 3.2 声明单个变量并赋值
	println("声明变量", "n8", n8)
	var (
		n9 = 1
	) // 3.3 声明单个推导变量
	println("声明变量", "n9", n9)
	var (
		n10 int
		n11 float32
	) // 3.4 声明多个变量 , 需要使用分号或换行分割
	println("声明变量", "n10", n10, "n11", n11)
	var (
		n12 int     = 3
		n13 float32 = .8
	) // 3.5 声明多个变量, 并初始化赋值
	println("声明变量", "n12", n12, "n13", n13)
	var (
		n14 = 3
		n15 = .8
	) // 3.6 声明多个推导变量
	println("声明变量", "n14", n14, "n15", n15)
	// 4. 推导声明(或者叫auto声明) @NOTE 只能在函数体中出现
	p1 := 5 // 4.1 单个推导声明, 不能指定类型
	println("声明变量", "p1", p1)
	p2, p3 := 3, 8 // 4.2 多个推导声明
	println("声明变量", "p2", p2, "p3", p3)
}

func declareConst() {
	println("常量的声明")
	// @NOTE 常量的声明必须初始化赋值
	// @NOTE 常量不能推导声明
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
	const (
		n8 int = 1
	) // 3.1 声明单个常量并赋值
	println("声明常量", "n8", n8)
	const (
		n9 = 1
	) // 3.2 声明单个推导常量
	println("声明常量", "n9", n9)
	const (
		n12 int     = 3
		n13 float32 = .8
	) // 3.3 声明多个常量, 并初始化赋值
	println("声明常量", "n12", n12, "n13", n13)
	const (
		n14 = 3
		n15 = .8
	) // 3.4 声明多个推导常量
	println("声明常量", "n14", n14, "n15", n15)
}

func controlStructures() {
	println("if 语句")
	var x int = 10
	var y = false
	// if [init]; condation {
	// [} else [if [init]; condation] {]
	// 只接收bool类型的常量, 变量或表达式
	// 关键字和花括号必须在一行, 除非在初始语句换行
	if false { // 常量
		// skip
	} else if y { // 变量
		// skip
	} else if x < 0 { // 表达式
		// skip
	} else if n := x; n < 0 { // 初始语句可以是 推导声明
		// skip
	} else if m, n := 0, 1; // 初始语句换行
	m > n {                 // 多个也没问题
		// skip
	} else if false && true { // 表达式
		// skip
	} else { // 也可以没有 else
		println("if")
	}

	println("switch 语句")
	// switch [condation] {
	// case condation : [body][fallthrough]
	// [default : {]
	// switch语句设计的比较灵活
	// switch参数可以是常量, 变量, 表达式, 或省略(参数为bool类型的true值)
	// case参数可以是常量, 变量, 表达式, 参数类型需要和switch参数值类型一致
	// switch参数可以是 推导声明语句, 此时case参数类型为bool
	// case可以带多个参数, 用逗号隔开
	// switch语句命中则不继续匹配, 除非用fallthrough关键字
	var m int = 0
	var n int = -1
	switch m { // switch参数为 int类型变量, case 参数类型相同
	case -1: // 常量 空的 case 体
	case n:
		return // 变量 可以用 return 返回函数值
	case n - 1:
		println("skip") // 表达式
	case 'a', n - 2:
		println("skip") // 可以匹配多项, 用逗号隔开
	case n + 1:
		println("m = -1 + 1") // 匹配执行
		fallthrough           // 不跳出, 继续匹配
	case 0:
		println("m = 1")
	default: // 可省略
	}

	switch 1 { // 常量
	case 0:
		println("skip")
	}

	switch { // 省略 相当于 switch true
	case false:
		println("skip")
	}

	switch m + 1 { // 表达式
	case -1:
		println("skip")
	}

	switch x := m; { // 推导声明语句, case参数为bool
	case x > 0:
		println("skip")
	case x < 0:
		println("skip")
	}

	switch x, y := m, n; { // 推导声明语句, case参数为bool
	case x < y:
		println("skip")
	case x == y:
		println("skip")
	}

	println("for 语句")
	// for [init]; condation; [post] {
	// }
	// for [condation] {
	// }

	for i := 0; i < 10; i++ {
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
	for i < 3 { // 省略 init 和 post
		i++
	}
	for i < 4 { // 只有表达式, 类似 while () {}
		i++
	}
	for { // 无限循环 类似 for(;;) {} 或 while (true){}
		break
	}

	var v = 0
	// 再for关键字前定义 循环标签 loopFor
loopFor:
	for i := 0; i < 5; i++ {
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

	println("for range 语句")
	// for condation range condation {
	// }
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
	go func(c chan int) { // 简单的协程和闭包
		c <- 1
		c <- 2
		close(c)
	}(c) // 虽然func定义可以用()包裹, 看起来会更像闭包, 何如
	for val := range c {
		println("rang chan:", val)
	}

}

func operationExpression() {
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
	*p1++  	// n1++
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
}

func keyWord() {

}

func main() {
	noop()              // 函数的最简定义
	builtinType()       // 内置基本数据类型
	builtinValue()      // 内置常量和变量
	declareVar()        // 变量的声明及赋值
	declareConst()      // 常量的声明及赋值
	controlStructures() // 控制结构
	operationExpression() // 运算
	keyWord()           // 保留字,关键字
	// 内置函数
	// 运算符和表达式
	// 函数
	// p1 := 5
	// p2 := p1
	// p2 += p1*1223372036854775807
	// var x int64 = int64(p2)
	// m := 2.3
	// println(p1, p2, x, m)
	a := 2 < 3 || 3 > 2 && 1 < 2
	println(a)
}
