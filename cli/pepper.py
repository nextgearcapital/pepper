import click

@click.group()
@click.version_option("0.1.0")
@click.pass_context
def pepper(ctx):
    pass
