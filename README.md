# Brook自动安装改端口脚本

## Feature
+ 一键安装brook脚本
+ http更改brook端口

## Install
```shell
URL=$(curl -i https://github.com/txthinking/brook/releases/latest | awk '/^location/{print $2}' | tr -d '\r\n')
LINUX_URL="${URL/tag/download}/brook_linux_amd64"
APK_URL="${URL/tag/download}/Brook.apk"
WIN_URL="${URL/tag/download}/brook_windows_amd64.exe"

echo "download brook(linux): $LINUX_URL"
echo "download brook(apk): $APK_URL"
echo "download brook(windows): $WIN_URL"

mkdir dist
wget $LINUX_URL -O dist/brook_linux_amd64
wget $APK_URL -O dist/Brook.apk
wget $WIN_URL -O dist/brook_windows_amd64.exe
chmod +x dist/brook_linux_amd64
ln -s $PWD/dist/brook_linux_amd64 /usr/local/bin/brook

wget https://github.com/kaiwen/bai/releases/download/v1.0.0/bagent
chmod +x bagent
openssl req -newkey rsa:2048 -new -subj "/C=US/ST=FL/L=Miami/O=FO/OU=Personal/CN=Alice" -nodes -x509 -days 3650 -keyout key.pem -out cert.pem
SERVER_PORT=443 BROOK_PORT=3030 BROOK_PASS=sec3et ./bagent
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