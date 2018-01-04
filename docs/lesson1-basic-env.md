# GO基础环境

## helloworld

下载安装go将程序路径加入环境变量PATH中

添加GOPATH环境变量,值为某个目录

GOPATH目录下建立 `bin pkg src`目三个录

命令行`cd $GOPATH/src/`或`cd %GOPATH%\src\`

新建并进入到helloword目录`mkdir helloword && cd helloword`

在src下新建`helloword.go`

编辑helloword.go

```go
package main

func main() {
  print("Hello World")
}

```

运行`go run helloword.go`

## 思考

如果不设置GOPATH会怎么样?

helloword目录为什么要放到GOPATH路径下?

GOPATH下面的三个目录都有什么作用?

`go run helloword.go` 是怎么运作的?
