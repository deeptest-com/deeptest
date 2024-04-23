# [LeYanAPI](https://leyanapi.nancalcloud.com/)

## LeYanAPI, Software API management and automation testing platform - Frontend
 written in Go with Iris, Gorm, MySQL and Vue3, Protobuf, WebSocket. 

## 乐研API, 软件API管理和自动化测试平台

### 快速开始

```bash
# 拉取后端的配置项目
git clone git@gitlab.nancalcloud.com:leyanapi/leyanapi-backend.git

# 拉取后端的配置项目
cd leyanapi-backend
git clone git@gitlab.nancalcloud.com:leyanapi/backend-config.git

# 启动server服务
go rum cmd/server/main.go

# 启动agent服务
go rum cmd/agent/main.go
```


### 打包客户端
sh ./init.project.sh
```bash
# 打包客户端 mac
sudo make ly-mac

# 打包客户端 windows
sudo make ly-win64
```
