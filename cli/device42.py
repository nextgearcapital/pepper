### In project packages ###
from cli.pepper import pepper


@pepper.group()
def device42():
    """Interact with the device42 API."""
