#! /usr/bin/env python3

import urllib.request
import json
from tester import chunks, request_json, baseurl

if __name__ == '__main__':
    # body = {"token":"abc", "v": 2.1, "players": [{"name": "Added1", "description": "Desk1"}, {"name": "Added2", "description": "Desk2"}]}
    # resp = request_json(body, baseurl.format('players', 'add'))
    body = {"token":"abc", "v": 1.0, "offset": 0, "limit": 200}
    resp = request_json(body, baseurl.format('players', 'get'))
    print(resp)
    print(len(resp['data']['items']))

    # resp = request_json({'key': 'mysecretkey'}, 'http://localhost:8080')
    # print(resp)
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
#         # break
