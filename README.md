# filterStr

install

```shell script
go get github.com/charliecen/filterStr
```

doc

```go
// 剪切中文或英文字符
res := filterStr.Cut.CutString("有钱没钱回家过年", 3)
fmt.Println(res)    //有钱没

// 过滤敏感词
err := filterStr.Filter.FilStr("考试作弊")
fmt.Println(err)    // 存在敏感词

```
> 注意: 存放敏感词的文件和目录拷贝到当前项目的根目录下，敏感词可以在`config`目录下的文件中添加。