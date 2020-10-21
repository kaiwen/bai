#!/bin/bash

set -e

PORT=${1:-3000}
PASS=$2

if [[ -z "$2" ]]
then
	PASS=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
fi

# update
apt update
apt upgrade -y
apt install nginx uwsgi-plugin-python3 python3-pip -y
apt autoremove -y

pip3 install flask

# change nginx config
sed -i -e 's/listen 80 .*/listen 443 ssl http2 default_server;/' \
	-e '/^\s\{1,\}listen \[/d; /^\s\{1,\}listen 443/a\\tssl_certificate certs/cert.pem;\n\tssl_certificate_key certs/key.pem;' \
	-e '/^\s\{1,\}location /i\\tlocation /port {\n\t\tinclude uwsgi_params;\n\t\tuwsgi_pass unix:/tmp/fb.sock;\n\t}\n' \
	/etc/nginx/sites-available/default

mkdir /etc/nginx/certs
yes '' | openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout /etc/nginx/certs/key.pem -out /etc/nginx/certs/cert.pem
systemctl reload nginx

nohup uwsgi config.ini >fb.log &

# download brook
URL=$(curl -i https://github.com/txthinking/brook/releases/latest | awk '/^location/{print $2}' | tr -d '\r\n')
echo "latest brook download page: $URL"

linux_URL="${URL/tag/download}/brook_linux_amd64"
windows_URL="${URL/tag/download}/Brook.exe"

echo "download brook(linux): $linux_URL"
echo "download brook(windows): $windows_URL"

wget $linux_URL
wget $windows_URL
cp -f brook_linux_amd64 /usr/local/bin/brook
chmod +x /usr/local/bin/brook
nohup brook server -l 0.0.0.0:$PORT -p $PASS >brook.log &

sleep 1

echo
echo "PORT=$PORT PASS=$PASS"
echo "Client to use brook_linux_amd64 Brook.exe"
echo
echo "Install Successfuly!"
