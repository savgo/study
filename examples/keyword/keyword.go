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
