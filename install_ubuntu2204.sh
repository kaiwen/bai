#!/bin/bash

set -e

PORT=${1:-3000}
PASS=$2

if [[ -z "$2" ]]
then
	PASS=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
fi

# download brook
URL=$(curl -i https://github.com/txthinking/brook/releases/latest | awk '/^location/{print $2}' | tr -d '\r\n')
echo "latest brook download page: $URL"

linux_URL="${URL/tag/download}/brook_linux_amd64"
apk_URL="${URL/tag/download}/Brook.apk"

echo "download brook(linux): $linux_URL"
echo "download brook(apk): $apk_URL"

wget $linux_URL
wget $apk_URL
cp -f brook_linux_amd64 /usr/local/bin/brook
chmod +x /usr/local/bin/brook

# download agent
wget https://github.com/kaiwen/bai/releases/download/v2.0/bagent -O /usr/local/bin/bagent
chmod +x /usr/local/bin/bagent
nohup /usr/local/bin/bagent $PORT $PASS &

echo
echo "PORT=$PORT PASS=$PASS"
echo "Install Successfuly!"
