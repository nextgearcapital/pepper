Pepper
===
[![GoDoc](https://godoc.org/github.com/nextgearcapital/pepper?status.svg)](https://godoc.org/github.com/nextgearcapital/pepper)
[![Build Status](https://travis-ci.org/nextgearcapital/pepper.svg?branch=master)](https://travis-ci.org/nextgearcapital/pepper)

Most of the code here is not very good. There was a desperate need for this at the time
so we needed to work very quickly. I plan on refactoring all of this at some point and
adding tests.

VMware vSphere:

    case "nano":
		CPU = 1
		Memory = 512
		DiskSize = 20
    case "micro":
		CPU = 1
		Memory = 1024
		DiskSize = 20
    case "small":
		CPU = 1
		Memory = 2048
		DiskSize = 40
	case "medium":
		CPU = 2
		Memory = 4096
		DiskSize = 60
	case "large":
		CPU = 2
		Memory = 8192
		DiskSize = 80
	case "xlarge":
		CPU = 4
		Memory = 16384
		DiskSize = 100
	case "ultra":
		CPU = 8
		Memory = 32768
		DiskSize = 160
	case "mega":
		CPU = 16
		Memory = 65536
		DiskSize = 200
