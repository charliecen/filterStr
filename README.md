# filterStr

install

```
go get github.com/charliecen/filterStr
```

doc

```
// 剪切中文或英文字符
fliterStr.CutStr("string", 6)

// 过滤敏感词
filterStr.FilStr("string")
```

> 敏感词可以在`config`目录下的文件中添加。当然还可以先截取字符串，然后再过滤敏感词。