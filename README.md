# Brook自动安装改端口脚本

## Feature
+ 自动安装部署brook脚本
+ http更改brook端口

## API
| method | endpoint | paramaters | body | note |
| -- | -- | -- | -- | -- |
| GET | /port | pass={password} | | get brook port |
| POST | /port | | port={port}&pass={password} | set brook port |
| POST | /pass | | old_pass={oldPassword}&new_pass={newPassword} | set brook password |
| POST | /port_pass | | old_pass={oldPassword}&new_pass={newPassword}&port={port} | set brook password and port |