[简体中文](README.md) | [English](README.en.md)

# go-test

> 何全的 Go 个人练习仓库，记录学习 Go 语言过程中的示例代码，不构成可用项目。

## 📖 项目介绍

这是个人学习 Go 时的练手仓库，主要跟随李文周老师的 Go 教程边学边写。代码按学习日期和主题分目录存放，每个 `main.go` 都是一个独立可运行的小练习。

## 📚 学习资料

- B 站视频：https://www.bilibili.com/video/av76169197
- 讲师博客：https://www.liwenzhou.com/posts/Go/go_menu/

## 📁 目录内容

```
day01/            基础语法：变量、常量、字符串操作（byte/rune）、for 循环、九九乘法表
day02/            学习占位目录（暂无内容）
练习/
├── main.go           interface 接口练习（Animal / Cat / Dog）
├── web/http/         net/http 标准库 Web 服务示例（监听 :8000）
├── 基础/类型/         strings、reflect、json、logrus 用法练习
├── 时间/             时间处理练习（暂无内容）
└── 权限/对象权限/     casbin ABAC 对象权限练习（含 abac_model.conf）
```

## 🚀 运行

任意练习均可直接运行，例如：

```bash
go run day01/main.go
go run 练习/web/http/main.go
```

## 🛠 技术栈

- Go 1.12+
- 练习中用到：casbin、logrus 等库

## 📝 学习进度

- 2019/12/24：变量、常量、iota、数字类型、字符串常用操作（strings 包）、byte/rune、if/for/switch、goto；视频看到第 20 集

## 📄 License

[MIT](LICENSE) © 2019 何全
