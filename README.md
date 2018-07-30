# 介绍
golang plugin热更新测试，源代码包含两部分，一个是engine,主要负责数据存储和so加载，logic则是相应的逻辑代码，最终我们将他编译为so，供engine加载。

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