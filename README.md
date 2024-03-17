# Brook自动安装改端口脚本

## Feature
+ 一键安装brook脚本
+ http更改brook端口

## Install
```shell
wget https://github.com/kaiwen/bai/releases/download/v1.0.0/bagent
chmod +x bagent
SERVER_PORT=80 BROOK_PORT=3030 BROOK_PASS=sec3et ./bagent
```

## API
| method | endpoint | paramaters | body | note |
| -- | -- | -- | -- | -- |
| GET | /port | pass={password} | | get brook port |
| POST | /port | | port={port}&pass={password} | set brook port |
| POST | /pass | | old_pass={oldPassword}&new_pass={newPassword} | set brook password |
| POST | /port_pass | | old_pass={oldPassword}&new_pass={newPassword}&port={port} | set brook password and port |