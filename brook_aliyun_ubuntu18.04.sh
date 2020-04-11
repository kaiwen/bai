#!/bin/sh

set -e

# change sshd port first
sed -i 's/#Port 22/Port 2000/g' /etc/ssh/sshd_config
systemctl reload sshd

# update
apt update
apt upgrade -y
apt autoremove -y
apt install nginx uwsgi-plugin-python3 -y

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
wget -O /usr/local/bin/brook https://github.com/txthinking/brook/releases/download/v20200201/brook
chmod +x /usr/local/bin/brook
nohup brook server -l 0.0.0.0:3000 -p superk >brook.log &

echo
echo
echo
echo "Install Successfuly!"
