# 介绍
由于golang是编译型语言，所有执行的程序都已经编译成汇编，所以不能像脚本语言那样动态更新代码。如果是无状态服务，比如web(通常数据都是依赖第三方程序提供，比如:mysql、redis)，可以配合fd继承来做到优雅重启。但是如果是类似游戏的有状态服务，一旦重启数据就丢了(当然也可以在重启前保存所有数据，然后启动后重新加载，只是这样就不够优雅了，毕竟有时候为了一个小的改动，大费周章重启服务，有点繁琐)。如果我们将数据和算法分离，数据由引擎部分保存，算法由另一个模块提供，修改代码后，编译成so，重新加载，这样不就可以完成了，正好golang提供了plugin这样的功能，可以轻松实现。源代码包含两部分，一个是engine,主要负责数据存储和so加载，logic则是相应的逻辑代码，最终我们将他编译为so。

# 测试步骤
1、编译engine.
```shell
sh build.sh
```

2、编译第1个版本so(注意后面有个参数）
```shell
sh build_so.sh 1
```

3、将src/logic/main.go里面的modelVersion和modelName分别改成1002和game2
```go
var (
	modelVersion = 1002
	modelName = "game2"
)
```

4、编译第2个版本so(注册后面的参数)
```shell
sh build_so.sh 2
```

5、运行engine
```shell
./engine
```

6、浏览器输入127.0.0.1:12345/hello，会看到如下显示（这是使用的第一个版本so)
```shell
hello test, this is golang plugin test!, version=1001, name=game1, oldversion=0, oldName=
```

7、浏览器输入127.0.0.1:12345/load?name=plugin2.so（这里输出done,就说明加载so成功了)

8、再次输入127.0.0.1:12345/hello，会看到如下显示。
```shell
hello test, this is golang plugin test!, version=1002, name=game2, oldversion=1001, oldName=game1
```