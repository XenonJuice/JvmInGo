# JvmInGo

这个分支添加了一些用于模拟 Java 8 classpath 识别和 class 文件读取的测试文件。

## 测试内容

当前测试数据覆盖以下 classpath 场景：

- 目录形式的 classpath
- jar 文件形式的 classpath
- 通配符形式的 classpath
- 多个 classpath 拼接的复合 classpath
- 模拟 Java 8 `jre/lib/*` 的 boot classpath

## 测试命令

在项目根目录下运行：

```bash
go run . -cp testdata/classes com.example.Hello
go run . -cp testdata/lib/hello-test.jar com.example.Hello
go run . -cp 'testdata/lib/*' com.example.Hello
go run . -cp 'testdata/classes:testdata/lib/hello-test.jar' com.example.Hello
```

也可以测试模拟的 boot classpath：

```bash
go run . java.lang.String
```

## 预期结果

成功时会输出当前解析到的 classpath、目标 class 名称，以及读取到的 class 文件字节数据。

示例输出：

```text
classPath: <project-root>/testdata/classes class: com.example.Hello args: []
class data : [202 254 186 190 ...]
```

其中：

```text
202 254 186 190
```

对应十六进制：

```text
CA FE BA BE
```

这是 Java `.class` 文件的 magic number，表示已经成功读取到 class 文件内容。

不同测试命令的 `classPath` 会不同，例如：

```text
<project-root>/testdata/classes
<project-root>/testdata/lib/hello-test.jar
<project-root>/testdata/lib/commons-lang3-3.14.0.jar:<project-root>/testdata/lib/hello-test.jar
<project-root>/testdata/classes:<project-root>/testdata/lib/hello-test.jar
```

只要输出中包含：

```text
class data : [202 254 186 190 ...
```

就说明 classpath 查找和 class 文件读取流程已经执行成功。

## 测试文件说明

```text
jre/lib/rt-test.jar
```

用于模拟 Java 8 的 boot classpath，其中包含测试用的 `java/lang/String.class`。

```text
testdata/classes/com/example/Hello.class
testdata/lib/hello-test.jar
testdata/lib/commons-lang3-3.14.0.jar
```

用于测试用户 classpath 的目录、jar、通配符和复合路径场景。
