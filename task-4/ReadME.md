# 个人博客系统后端

## 项目结构

```text
task-4/
├── ReadME.md
├── 个人博客系统后端接口文档.md
├── blog-server/
    ├── server.exe
    ├── main.go
    ├── config.yaml
    ├── db.sql
    ├── go.mod
    ├── go.sum
    ├── config/
    │   └── config.go
    ├── models/
    │   ├── user.go
    │   └── post.go
    ├── api/
    │   ├── authApi.go
    │   └── postApi.go
    ├── middleware/
    │   ├── jwt.go
    │   └── errorHandler.go
    ├── database/
    │   └── connection.go
    ├── responses/
    │   └── response.go
    └── logger/
        └── logger.go
```

## 项目运行环境

- 项目主函数：`main.go`
- 项目启动方式：
  - 方式一：双击启动 `server.exe`执行文件启动服务
  - 方式二：使用命令行进入`/task-4/blog-server/`目录，执行`go run main.go`命令启动服务

- 运行环境：Windows10 - 11

## 测试用例及测试结果

请查阅参考`个人博客系统后端接口文档.md`
