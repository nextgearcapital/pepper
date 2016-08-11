# pepper

A CLI for the datacenter!

This is super alpha software. The goal with pepper is to be a one-stop shop for accessing anything we deal with
from a provisioning perspective.

### How to install locally

Requirements:
* python 3.4 or above
* pip
* virtualenv

Make sure you have Python 3.4 or above installed. Get it from Homebrew:

```sh
$ brew install python3
```

You can install pip like this:

```sh
$ curl -L https://bootstrap.pypa.io/get-pip.py | sudo -H python3
```

Then you need virtualenv. You can install virtualenv like this:

```sh
$ pip3 install virtualenv
```

Clone this repo:

```sh
$ git clone https://github.com/nextgearcapital/pepper
```

cd to the pepper repo and get a virtualenv started:

```sh
$ virtualenv venv
```

Activate your virtualenv:

```sh
$ . venv/bin/activate
```

Install pepper locally:

```sh
$ pip install --editable .
```

If you're done with the environment, you can do this in the root of the repo:

```sh
$ deactivate
$ make clean
```
