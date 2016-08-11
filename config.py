import logging
import os

### In project packages ###
from cli.pepper import pepper

log = logging.getLogger(__name__)

@pepper.command()
def init():

    if not os.path.exists("/etc/pepper"):
        log.warning(
            "/etc/pepper doesn't exist so your config most likely doesn't exist either.\n")
        log.info("Creating directory...")
        try:
            os.makedirs("/etc/pepper/config.d")
            os.makedirs("/etc/pepper/provider.d")
        except OSError:
            raise SystemExit("Couldn't create the /etc/pepper, /etc/pepper/config.d, or /etc/pepper/provider.d directories.")

        log.info("Created directory.")
        log.info("Generating example /etc/pepper/config.d/template.toml for you...\n")

        configd_toml = open("/etc/pepper/config.d/template.toml", "w")

        configd_tmpl = """provider = "vcenter01"
        dhcp = true
        network = "Development"
        gateway = 192.168.1.1
        subnet = 255.255.255.0
        domain = "google.com"
        dns_servers = [
            8.8.8.8,
            8.8.4.4,
        ]
        cluster = "Development"
        folder = "Development"
        datastore = "test"
        """

        try:
            configd_toml.write(configd_tmpl)
        except OSError:
            log.critical(
                "Couldn't write /etc/pepper/config.d/template.toml, Jim! Probably a permissions thing?")
            raise SystemExit

        configd_toml.close()

        log.info("Success! /etc/pepper/config.d/template.toml has been generated for you. Please go and change the values.")

        log.info("Generating example /etc/pepper/provider.d/template.toml for you...\n")

        providerd_toml = open("/etc/pepper/provider.d/template.toml", "w")

        providerd_tmpl = """base_url = "https://d42.company.com"
        username = "admin"
        password = "superl33t"
        ip_range = 172.18.16.0/24
        service_level = "Development"
        """

        try:
            providerd_toml.write(providerd_tmpl)
        except OSError:
            log.critical(
                "Couldn't write /etc/pepper/provider.d/template.toml, Jim! Probably a permissions thing?")
            raise SystemExit

        providerd_toml.close()

        log.info("Success! /etc/pepper/provider.d/template.toml has been generated for you. Please go and change the values.")
