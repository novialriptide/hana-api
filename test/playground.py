"""
This is just a temp test file that'll get replaced by a go
unittest collection later.
"""
import requests

a = requests.post("http://localhost:25565/albums", params={
    "album_name": "lmao",
    "song_ids": ""
}).json()

print(a)