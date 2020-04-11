# 阿里云Brook自动安装脚本

## 安装
1. 配置一个阿里云ECS
2. 网页VPN登陆，防止被防火墙发现
3. 运行 `./brook_aliyun_ubuntu18.04.sh`

## 查看状态
+ `curl -k https://x.x.x.x/port` 获取当前brook监听端口
+ `curl -k -d 'port=1234' https://x.x.x.x/port` 设置brook监听端口为1234

## TODO
后续各个系统脚本增加中，目前主要根据我经常用来翻墙购买的ECS配置 ;-)
