# 环境

- go 1.18.3
- postgres 15

# 配置文件

cp config/app.ini.example config/app.ini

修改 里面的配置项

# 启动服务

## 手动重启

go run main.go

## 自动重启

$ go install github.com/cosmtrek/air@latest
$ air
