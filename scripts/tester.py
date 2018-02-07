#! /usr/bin/env python3

import urllib.request
import json
# import psycopg2
# conn = psycopg2.connect(host='127.0.0.1', port=5432, user='postgres')
# cur = conn.cursor()

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


geturl = "http://localhost:3003/method/players.get"
addurl = "http://localhost:3003/method/players.add"
updateurl = "http://localhost:3003/method/players.update"
getbyidurl = "http://localhost:3003/method/players.getById"

if __name__ == '__main__':
    passed = True
# players.add

    body = {"token":"abc", "v": 2.1, "players": [{"name": "Added1", "description": "Desk1"}, {"name": "Added2", "description": "Desk2"}]}
    resp = request_json(body, addurl)

    passed &= resp['status'] == 'ok'

    added_ids = resp['data']['items']

    passed &= len(added_ids) == 2

# players.get
    body = {"token":"abc", "v": 1.0, "offset": added_ids[0] - 1, "limit": 2}
    resp = request_json(body, geturl)

    passed &= resp['status'] == 'ok'

    items = resp['data']['items']

    passed &= items[0]['name'] == 'Added1'
    passed &= items[1]['name'] == 'Added2'

    ids_got = [x['id'] for x in items]

    passed &= added_ids == ids_got

# players.update
    body = {"token":"abc", "v": 1.0, "players": [{"id": added_ids[0], "name": "Updated1"}, {"id": added_ids[1], "name": "Updated2"}]}
    resp = request_json(body, updateurl)

    passed &= resp['status'] == 'ok'

    updated_ids = resp['data']['updated_ids']
    passed &= added_ids == updated_ids

# players.getById
    body = {"token":"abc", "v": 1.0, "id": updated_ids[0]}
    resp = request_json(body, getbyidurl)

    passed &= resp['status'] == 'ok'
    passed &= resp['data']['player']['id'] == updated_ids[0]

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
    #         # break\
    #
    #
    # url = updateurl
    # body = {"token":"abc", "v": 2.1, "players": [{"id": 1, "name": "Sasha"}, {"id": 2, "name": "Sasha2"}]}
    # req = urllib.request.Request(url)
    # req.add_header('Content-Type', 'application/json; charset=utf-8')
    # jsondata = json.dumps(body)
    # jsondataasbytes = jsondata.encode('utf-8')   # needs to be bytes
    # req.add_header('Content-Length', len(jsondataasbytes))
    # # print (jsondataasbytes)
    # response = urllib.request.urlopen(req, jsondataasbytes)
    # # print (response)
    # print (response.read())
    if passed:
        print('PASSED')
    else:
        print('NOT PASSED')
