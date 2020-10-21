# AWS/阿里云 Brook自动安装改端口脚本
如果你的梯子经常被封，可以看下这个
Tested **AWS** **Aliyun**

## Feature
+ http接口查询/更改端口
+ 自动安装部署brook脚本

## TL;DR
+ 端口被封: `curl -k -d 'port=1234' https://x.x.x.x/port` 设置brook监听端口为1234(更改'port=xxxx'可随意更改监听端口)
+ IP被封: `curl -L https://github.com/kaiwen/bai/releases/download/script/install.sh | sh`

## 安装

1. 配置一个AWS或者阿里云ECS, 系统选择ubuntu 20.04或者18.04
2. `curl -L https://github.com/kaiwen/bai/releases/download/script/install.sh | sudo sh`
3. 或者 `git clone https://github.com/kaiwen/bai.git && cd bai && sudo ./brook_aliyun_ubuntu18.04.sh`

安装后
+ brook默认监听端口**3000**
+ 密码**superk** (请手动改密码)

## 查看状态

+ `curl -k https://x.x.x.x/port` 获取当前brook监听端口
+ `curl -k -d 'port=1234' https://x.x.x.x/port` 设置brook监听端口为1234(更改'port=xxxx'可随意更改监听端口)

## TODO

1. 后续其他系统脚本增加(目前主要根据自己需要)
