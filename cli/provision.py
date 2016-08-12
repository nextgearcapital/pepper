import click

from cli.pepper import pepper


@pepper.group()
def provision():
    """Provisisons new servers."""
    raise SystemExit("Not implemented.")

@provision.command()
@click.option("--name", required=True,
              help="The server's hostname.")
@click.option("--template", required=True,
              help="The template (OS) you want to use.")
@click.pass_obj
def virtual(config, name, location):
    """Provisions a virtual server."""



