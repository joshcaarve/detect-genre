#!/usr/bin/env python3

import requests
import sys
import os
import json

def print_json(dict_json):
    print(json.dumps(dict_json, indent=4))

port = "8080"
base = "localhost"

#URL = "http://{}:{}/{}".format(sys.argv[1], sys.argv[2], sys.argv[3])
URL = "http://{}:{}/{}".format(base, port, sys.argv[1])
headers = {
    "Content-Type": "application/json"}
data = {
    "genre": "soul",
    "parentGenre": "Deep"
}
data_to_send = json.dumps(data).encode("utf-8")
response = requests.post(URL, headers=headers, data=data_to_send)
print_json(response.json())
