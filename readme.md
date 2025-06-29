# ddd

domain-application-ports-adapter

基础设施-领域-应用-接口


## 依赖的工具

- task
- goreleaser

```golang
go install github.com/goreleaser/goreleaser/v2@latest
go install github.com/go-task/task/v3/cmd/task@latest
```
- upx

## 目标

- 连接查询的结果与一般实体
- gorm gen 工具添加
- 多种命令支持
- env添加
- 配置文件支持
- 依赖注入支持
- 测试编写

## 参考

架构整洁之道
https://github.com/leewaiho/Clean-Architecture-zh?tab=readme-ov-file

https://gorm.io/gen/index.html