from app import app
from flask import Flask, Response, json, request

@app.route('/')
def home():
	return "hello world!"

@app.route('/api/<text>', methods=['POST'])
def index(text):
	json_string = json.dumps(text)
	json_string = json.loads(json_string)
	json_string = json.dumps(text)
	json_string = json.loads(json_string)
	json_string = json.dumps(text)

	return Response(json_string, mimetype="aplication/json; charset=utf-8")
