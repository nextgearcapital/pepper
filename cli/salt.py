### In project packages ###
from cli.pepper import pepper


@pepper.group()
def salt():
    """Interact with salt/salt-cloud."""
