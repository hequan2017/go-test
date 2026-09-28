[简体中文](README.md) | [English](README.en.md)

# go-test

> Hequan's personal Go practice repository — sample code written while learning the Go language. Not a usable project.

## 📖 Introduction

This is a hands-on repository from my Go learning journey, following Li Wenzhou's Go tutorials. Code is organized into directories by date and topic; each `main.go` is a small standalone, runnable exercise.

## 📚 Learning Materials

- Bilibili video: https://www.bilibili.com/video/av76169197
- Instructor's blog: https://www.liwenzhou.com/posts/Go/go_menu/

## 📁 Contents

```
day01/            basics: variables, constants, string operations (byte/rune), for loops, 9x9 multiplication table
day02/            placeholder directory (empty for now)
练习/
├── main.go           interface exercises (Animal / Cat / Dog)
├── web/http/         net/http standard-library web server example (listens on :8000)
├── 基础/类型/         strings, reflect, json and logrus practice
├── 时间/             time-handling exercise (empty for now)
└── 权限/对象权限/     casbin ABAC object-permission exercise (with abac_model.conf)
```

## 🚀 Run

Any exercise can be run directly, for example:

```bash
go run day01/main.go
go run 练习/web/http/main.go
```

## 🛠 Tech Stack

- Go 1.12+
- Libraries used in exercises: casbin, logrus, etc.

## 📝 Progress

- 2019/12/24: variables, constants, iota, numeric types, common string operations (strings package), byte/rune, if/for/switch, goto; watched up to episode 20

## 📄 License

[MIT](LICENSE) © 2019 何全 (Hequan)
