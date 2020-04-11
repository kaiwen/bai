#!/usr/bin/python3

from flask import Flask, json, request
import subprocess
import time

app = Flask(__name__)
listen_port = 3000

@app.route('/port', methods=['GET', 'POST'])
def port():
	global listen_port
	if request.method == 'GET':
		return json.jsonify(port=listen_port)
	elif request.method == 'POST':
		listen_port = int(request.form['port'])
		subprocess.call(['killall', 'brook'])
		time.sleep(1)
		subprocess.call(f'brook server -l 0.0.0.0:{listen_port} -p superk &', shell=True)
		return json.jsonify(success=True, port=listen_port)

if __name__ == '__main__':
	app.run(debug=True)
