# -*- coding: UTF-8 -*-
import urllib2
import json
import requests
import md5
import base64

headers = {'content-type': 'application/json'}

#register&auth
'''
origin_passwd = "test"
m = md5.new(origin_passwd + "mu77")
passwd = m.hexdigest()

post_data = {
    "channel" : "mu77",
    "user": "cccccc",
    "passwd":passwd,
    "device_id":"asdasdasda"
}

res = requests.post("http://127.0.0.1:8080/signup", post_data , timeout=600)
print(res.text)

ret = json.loads(res.text)
access_token = ret["access_token"]

post_data = {
    "platform" : "mu77",
    "channel" : "mu77",
    "user": "cccccc",
    "device_id":"asdasdasda",
    "access_token" : access_token
}

res = requests.post("http://127.0.0.1:8080/auth", post_data , timeout=600)
print(res.text)

#login&auth
origin_passwd = "test"
m1 = md5.new(origin_passwd + "mu77")
passwd = m1.hexdigest()
device_id = "asdasdasd"
m = md5.new(passwd+device_id)
passwd = m.hexdigest()
post_data = {
    "user": "aaaaaa",
    "passwd":passwd,
    "device_id":device_id,
    "is_guest" : False
}
res = requests.post("http://127.0.0.1:8080/signin", post_data , timeout=600)
print(res.text)

ret = json.loads(res.text)
access_token = ret["access_token"]
print(len(access_token))

post_data = {
    "platform" : "mu77",
    "channel" : "mu77",
    "user": "aaaaaa",
    "device_id":"asdasdasd",
    "access_token" : access_token
}

res = requests.post("http://127.0.0.1:8080/auth", post_data , timeout=600)
print(res.text)
'''

#guest
m1 = md5.new("test")
passwd = m1.hexdigest()
device_id = "asdasdasdqweq"
m = md5.new(passwd+device_id)
passwd = m.hexdigest()
platform = "guest"
post_data = {
    "platform":platform,
    "device_id":device_id,
#    "is_guest" : True
}
#res = requests.post("http://127.0.0.1:8080/signin", post_data , timeout=600)
res = requests.post("http://127.0.0.1:8080/login", json.dumps(post_data), timeout=600, headers=headers)
print(res.text)

ret = json.loads(res.text)
access_token = ret["access_token"]

post_data = {
    "uid" : ret["uid"],
    "access_token" : access_token,
}

res = requests.post("http://127.0.0.1:8080/auth", json.dumps(post_data), timeout=600, headers=headers)
print(res.text)
