# filterStr

### 编译
```
go build -o main main.go
```

### 执行命令

```
// 帮助命令
❯ .\main.exe -h
Usage of C:\go\src\filterStr\main.exe:
  -l int
        长度 (default 6)
  -n string
        字符串 (default "string")
  -t string
        类型,c:截取字符串, f:过滤字符串是否有效 (default "c")

```

### 测试
```
// 截取字符串
❯ .\main.exe -t=c -n=过年了可以回家了哦 -l=3
过年了

// 过滤敏感词
❯ .\main.exe -t=f -n=考试作弊
存在敏感词

```

> 敏感词可以在`config`目录下的文件中添加。当然还可以先截取字符串，然后再过滤敏感词。