# Brook自动安装改端口脚本

## Feature
+ 一键安装brook脚本
+ http更改brook端口

## Install
```shell
URL=$(curl -i https://github.com/txthinking/brook/releases/latest | awk '/^location/{print $2}' | tr -d '\r\n')
LINUX_URL="${URL/tag/download}/brook_linux_amd64"
APK_URL="${URL/tag/download}/Brook.apk"

echo "download brook(linux): $LINUX_URL"
echo "download brook(apk): $APK_URL"

mkdir dist
wget $LINUX_URL -O dist/brook_linux_amd64
wget $APK_URL -O dist/Brook.apk
chmox +x dist/brook_linux_amd64
ln -s $PWD/dist/brook_linux_amd64 /usr/local/bin/brook

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

## TODO
+ brook client download api