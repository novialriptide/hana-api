"""
This is just a temp test file that'll get replaced by a go
unittest collection later.
"""
import requests

a = requests.post(
    "http://localhost:25565/songs",
    params={
        "album_id": "",
        "song_genre_id": "",
        "song_source": "",
        "song_name": "",
        "artist_ids": "",
    }
).json()["message"]
print(a)

b = requests.post(
    f"http://localhost:25565/songs/{a}/file",
    files={
        "file": open("Mittsies - Vitality (t+pazolite Remix).mp3", "rb")
    }
).json()
print(b)