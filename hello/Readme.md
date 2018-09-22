# go编程实践：自定义包和测试程序

---

参考文档：[官方文档——如何使用go编程](https://go-zh.org/doc/code.html)

---

## 1. centos配置go环境
可以参考我的另一篇博客
- [csdn]()
- [github]()

## 2. 基本概念补充
### 2.1 工作空间
go的工作空间实际上并没有强制规定，但为了提高工作效率，就不得不统一不同人使用工作空间的文件结构的习惯，否则每个人都是一套自己的文件结构和文件命名格式，在团队合作的时候可能很容易造成混乱。因此官方提出了一种工作空间文件结构的标准，希望大家能大致按这种文件结构和文件格式来组织工作空间，就像其他IDE自动生成的工作空间结构一样。

引入官方推荐结构：[https://golang.org/doc/code.html](https://golang.org/doc/code.html)
```
bin/
    hello              # 可执行的命令
    outyet             # 可执行的命令
pkg/
    linux_amd64/
        github/golang/example/
            stringutil.a          # pakcage objects
src/
    github.com/golang/example/
        .git/                      # Git repository metadata
	hello/
	    hello.go               # command source
	outyet/
	    main.go                # command source
	    main_test.go           # test source
	stringutil/
	    reverse.go             # package source
	    reverse_test.go        # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
	bmp/
	    reader.go              # package source
	    writer.go              # package source
    ... (many more repositories and packages omitted) ...

```
基于官方推荐的目录结构，我个人的目录结构组织成以下形式：
```
bin/
    hello              # executable command
    ...
    others             # executable command
pkg/
    linux_amd64/
        github.com/gitgier/ServiceComputing/
            stringutil.a          # pakcage objects
        github.com/other-users    # other package objects' dir
            ...
        golang.org/               # other package objects' dir
            ...                   
src/
    github.com/gitgiter/ServiceComputing/
        .git/                     # git repository metadata
        hello/
            hello.go              # command source
            Readme.md             # instruction or report
        stringuitl/
            reverse.go            # package source
            reverse_test.go       # test source
        ...
        others                    # other package's dirs
    github.com/other-users        # source from other users
        ...
    golang.org/                   # source from other website
        ...
```
这里简单说明和补充一下工作空间的文件目录：
- bin：用来存放可执行文件的目录。一般go install成功后的可执行文件都会出现在这个目录下面
- pkg：用来存放编译好并归档的包。一般go install成功后可执行文件所用到的包都会放在这个目录下
- src：用来存放go的源代码，各种包
    - github.com：该目录下的代码都是和github有关联的代码（clone下来的或准备push的）
        - user：github.com下有以不同用户名命名的文件夹，这些用户名一般都以github账户用户名为准，比如我的github用户名是gitgiter
            - example：可以认为是一个仓库名，因为每个用户下面都有许多仓库，所以用户文件夹下直属的就可以理解为仓库文件夹，每个仓库文件夹下可以有许多包。比如我的一个github仓库叫ServiceComputing，文件夹内的文件都是直接与仓库关联的
    - golang.org：类比github.com，可以理解成关联golang.org这个网站的代码存放的地方。有个特别的地方就是golang.org这个网站目前国内上不了，所以无法直接关联，然而我们需要的一些go工具在golang.org上，陷入尴尬的境地。现在最常用的一种解决办法就是去github.com上找相关的镜像，先下载到github.com目录下再迁移到golang.org下（假装是真的从golang.org上下的工具）

### 2.2 一些常用的go命令
- go get：下载和安装包和依赖
- go build：编译包和依赖
- go install：编译并安装包和依赖
- go run：编译并运行go程序
- go test：结合 *_test.go 文件测试某个包
- go env：查看go的环境配置信息
- go version：查看go的版本

**注意build、install、run的区别**

- go run是编译并直接运行程序，不会产生输出，虽然中间会产生临时文件
- go build只是编译，原则上也不会产生输出（直接编译main包例外），目的是看有没有编译错误
- go install可看成由两步组成，一步是编译导入的包文件，就像接下来将要提到的hello.go中导入的stringutil包，第二步是生成可执行文件放在bin目录下，编译后的包放到pkg目录下。

## 3. helloworld编写、调用和测试自定义包
### 3.1 第一个helloworld程序
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}
```

```sh
## global run
go run $GOPATH/src/github.com/gitgiter/ServiceComputing/hello/hello.go

## if in current dir, you can just run locally
go run hello.go

## you can install, too
go install $GOPATH/src/github.com/gitgiter/ServiceComputing/hello

## if in parent dir
go install hello
```

直接run是可以直接输入helloworld的结果的；如果是安装之后，则可以直接在命令行输入hello执行hello，前提是bin目录已经加入环境变量中。

### 3.2 编写一个自定义包
这里编写自定义包的目的是为了熟悉一下自定义包的流程，因此包的内容就没有过多的要求，所以我这里就直接采用教程文档里提供的例子。

对于我的工作空间文件结构，在ServiceComputing下创建一个包文件叫stringutil，将以下代码放入stringutil文件夹下的reverse.go文件中，这段代码的功能是将字符串倒序反转。
```go
// stringutil 包含有用于处理字符串的工具函数。
package stringutil

// Reverse 将其实参字符串以符文为单位左右反转。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

接下来可以用go build命令测试是否有编译错误等

## 3.3 更改hello.go，使其调用自定义包
导入自定义包的时候要指定路径
```go
package main

import "fmt"
import "github.com/gitgiter/ServiceComputing/stringutil"

/*

也可以写成

import （
    “fmt"
    "github.com/gitgiter/ServiceComputing/stringutil"
)

*/

func main() {
	fmt.Printf(stringutil.Reverse("\n!oG ,olleH"))
}

```

### 3.4 编写包测试程序
在reverse.go所在目录下，新建reverse_test.go文件
```go
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

### 3.5 go test查看测试结果
```sh
## 同理，可以全局也可以局部
go test $GOPATH/src/github.com/gitgiter/ServiceComputing/stringutil
go test ./stringutil
```

系列测试结果：

![test](test.jpg)