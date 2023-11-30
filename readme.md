# traversing

一个用于遍历生成目录下md文件结构的小工具

## 使用场景

mkdocs的项目需要将文章的信息写入至配置文件中，当一次需要添加非常多文件时就可以使用本工具生成可以使用的文章配置片段


## 构建

使用以下命令构建

```bash
go build -o traversing main.go
```


## 使用

格式如下

```bash
traversing -p 需要遍历的文件夹的路径
```
