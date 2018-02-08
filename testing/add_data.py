#! /usr/bin/env python3

import urllib.request
import json
# import psycopg2
# conn = psycopg2.connect(host='127.0.0.1', port=5432, user='postgres')
# cur = conn.cursor()

def chunks(l, n):
    for i in range(0, len(l), n):
        yield l[i:i+n]


addurl = "http://localhost:3003/method/players.add"
updateurl = "http://localhost:3003/method/players.update"

# # one by one
# with open('nogit/names.csv', 'r') as f:
#     for line in f:
#         name = line[:-1]
#         data = {'v': 2.1, 'players': [{'name': name}]}
#
#         req = urllib.request.Request(addurl)
#         req.add_header('Content-Type', 'application/json; charset=utf-8')
#         jsondata = json.dumps(data)
#         jsondataasbytes = jsondata.encode('utf-8')   # needs to be bytes
#         req.add_header('Content-Length', len(jsondataasbytes))
#         # print (jsondataasbytes)
#         response = urllib.request.urlopen(req, jsondataasbytes)
#         # print (response)
#         # print (response.read())

# testing with data
# with open('nogit/names.csv', 'r') as f:
#     l = [{"name": line[:-1], "rating": 0.1} for line in f]
#     for chunk in chunks(l, 1000):
#         data = {'v': 2.1, 'players': chunk}
#         req = urllib.request.Request(addurl)
#         req.add_header('Content-Type', 'application/json; charset=utf-8')
#         jsondata = json.dumps(data)
#         jsondataasbytes = jsondata.encode('utf-8')   # needs to be bytes
#         req.add_header('Content-Length', len(jsondataasbytes))
#         # print (jsondataasbytes)
#         response = urllib.request.urlopen(req, jsondataasbytes)
#         # print (response.read())
#         # break
#
#         # sql_names = ["INSERT INTO players(name) VALUES ('{}')".format(x['name']) for x in chunk]
#         # for sql in sql_names:
#         #     cur.execute(sql)
#         # conn.commit()
#         # break\
#
url = updateurl
body = {"token":"abc", "v": 2.1, "players": [{"id": 1, "name": "Sasha"}, {"id": 2, "name": "Sasha2"}]}
req = urllib.request.Request(url)
req.add_header('Content-Type', 'application/json; charset=utf-8')
jsondata = json.dumps(body)
jsondataasbytes = jsondata.encode('utf-8')   # needs to be bytes
req.add_header('Content-Length', len(jsondataasbytes))
# print (jsondataasbytes)
response = urllib.request.urlopen(req, jsondataasbytes)
# print (response)
print (response.read())
