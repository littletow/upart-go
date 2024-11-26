# 命令行上传文章小工具

## 安装方法

从项目的 release 下载对应操作系统的预编译的二进制包，解压缩到本地。

如果没有对应的操作系统安装包，请自行编译项目。方法如下：

```
git clone git@gitee.com:littletow/upart-go.git
cd upart-go
go build -ldflags='-s -w'
```

然后在命令行中运行`gart.exe init`，完成账号绑定。

## 使用方法

可以使用 `gart.exe --help` 查看有效命令。

例如：上传文章

```
gart.exe upload title keyword filename ispub islock
```

## 使用工具上传文章步骤

1. 制作 markdown 文件
2. 将文件放置到 files 文件夹中
3. 调用工具上传文章

## 豆子碎片小程序码

可扫下方小程序码查看效果：
![alt 豆子碎片](https://imgs.91demo.top/visit.webp)
