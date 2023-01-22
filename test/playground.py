"""
This is just a temp test file that'll get replaced by a go
unittest collection later.
"""
import requests

a = requests.post(
    "http://localhost:25565/songs/fgfgfg/file",
    files={
        "file": open("Mittsies - Vitality (t+pazolite Remix).mp3", "rb")
    }
).json()

print(a)