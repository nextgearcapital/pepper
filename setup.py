#!/usr/bin/env python3

from setuptools import setup

with open("requirements.txt") as f:
    requirements = f.read().splitlines()

setup(
    name="pepper",
    version="0.1.0",
    description="Pepper is a wrapper around salt-cloud that can also generate salt profiles.",
    packages=["cli", "resources"],
    include_package_data=True,
    install_requires=requirements,
    entry_points={
        "console_scripts": [
            "pepper=cli.main:main"
        ]
    }
)
