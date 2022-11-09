# 阿里云Brook自动安装改端口脚本

## Feature
+ 自动安装部署brook脚本
+ http更改brook端口

## 安装
### 自动安装 (默认密码)
1. 配置一个AWS或者阿里云ECS, 系统选择ubuntu 22.04
2. `curl -L https://github.com/kaiwen/bai/releases/download/v2.0/install_ubuntu2204.sh | sudo bash -s -- 3000 pass` 其中`pass`为你的密码，可以自己改

## 查看/更改端口
+ `curl -k https://x.x.x.x/port` 获取当前brook监听端口
+ `curl -d 'port=1234' http://x.x.x.x/port` 设置brook监听端口为1234
