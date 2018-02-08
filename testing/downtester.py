#! /usr/bin/env python3

import urllib.request
import json

def chunks(l, n):
    for i in range(0, len(l), n):
        yield l[i:i+n]

def request_json(data, url):
    req = urllib.request.Request(url)
    req.add_header('Content-Type', 'application/json; charset=utf-8')
    jsondata = json.dumps(data)
    jsondataasbytes = jsondata.encode('utf-8')   # needs to be bytes
    req.add_header('Content-Length', len(jsondataasbytes))
    response = urllib.request.urlopen(req, jsondataasbytes)
    return json.loads(response.read())

baseurl = 'http://localhost:3003/method/{}.{}'
geturl = "http://localhost:3003/method/players.get"
addurl = "http://localhost:3003/method/players.add"
updateurl = "http://localhost:3003/method/players.update"
getbyidurl = "http://localhost:3003/method/players.getById"

if __name__ == '__main__':
    for ent in ['player', 'team', 'tournament']:
        ents = ent + 's'
        # ents.add
        method = 'add'
        body = {"token":"abc", "v": 2.1, ents: [{"name": "Added1", "description": "Desk1"}, {"name": "Added2", "description": "Desk2"}]}
        resp = request_json(body, baseurl.format(ents, method))
        print(resp)

        # ents.get
        method = 'get'
        body = {"token":"abc", "v": 1.0, "offset": 0, "limit": 2}
        resp = request_json(body, baseurl.format(ents, method))
        print(resp)

        # ents.update
        method = 'update'
        body = {"token":"abc", "v": 1.0, ents: [{"id": 1, "name": "Updated1"}, {"id": 2, "name": "Updated2"}]}
        resp = request_json(body, baseurl.format(ents, method))
        print(resp)

        # ents.getById
        method = 'getById'
        body = {"token":"abc", "v": 1.0, "id": 3}
        resp = request_json(body, baseurl.format(ents, method))
        print(resp)
