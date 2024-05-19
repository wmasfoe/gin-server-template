## 开发环境启动

项目使用 [air](https://github.com/cosmtrek/air) 做了开发环境热更新：

```sh
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
air -v
```
> .air.toml 是 air 的配置文件，配置项见：[air_example.toml](https://github.com/cosmtrek/air/blob/master/air_example.toml)

设置 GO 环境变量：

```sh
GOPROXY=https://goproxy.cn,direct
GO111MODULE=on
```

启动：

```sh
go install github.com/swaggo/swag/cmd/swag@latest
go mod tidy
swag init
# 推荐用 air 启动，支持热更新
air
# 也可以选择用 go 启动，不支持热更新
go run .
```
