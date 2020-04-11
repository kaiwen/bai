# 阿里云Brook自动安装改端口脚本
如果你的梯子经常被封，可以看下这个

## Feature
+ http接口查询/更改端口
+ 自动安装部署brook脚本

## TL;DR
+ 端口被封: `curl -k -d 'port=1234' https://x.x.x.x/port` 设置brook监听端口为1234(更改'port=xxxx'可随意更改监听端口)
+ IP被封: `curl -L https://github.com/kaiwen/bai/releases/download/script/install.sh | sh`

## 安装

1. 配置一个阿里云ECS
2. 网页VNC登陆，防止被防火墙发现(部分网络运营商会阻止22端口)
3. `curl -L https://github.com/kaiwen/bai/releases/download/script/install.sh | sh`
4. 或者 `git clone https://github.com/kaiwen/bai.git && cd bai && ./brook_aliyun_ubuntu18.04.sh`

注意： 安装后sshd端口已改为**2000**，可自行修改

## 查看状态

+ `curl -k https://x.x.x.x/port` 获取当前brook监听端口
+ `curl -k -d 'port=1234' https://x.x.x.x/port` 设置brook监听端口为1234(更改'port=xxxx'可随意更改监听端口)

## TODO

后续各个系统脚本增加中，目前主要根据我经常用来翻墙购买的ECS配置 ;-)
