import base64

### In project packages ###
import config


def _b64encode(string):
    b = bytes(string, "utf-8")
    return base64.b64encode(b).decode("ascii")


def _b64decode(string):
    return base64.b64decode(string).decode("ascii")
