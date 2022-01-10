# about-me - Source code for trevor-edris.io

[![Go Report Card](https://goreportcard.com/badge/github.com/TrevorEdris/about-me)](https://goreportcard.com/report/github.com/TrevorEdris/about-me)
[![Test](https://github.com/TrevorEdris/about-me/actions/workflows/test.yml/badge.svg)](https://github.com/TrevorEdris/about-me/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoT](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev)

**A large majority of the base code for this project has been adapted from [https://github.com/mikestefanello/pagoda](https://github.com/mikestefanello/pagoda).** Many thanks to mikestefanello for a great template repository.

## Developing

Planned future work

## Deploying

### Local

#### Requirements

* Go 1.17
* Bash (tested with 5.0.17, though 3.x+ should work)
* Blackbox

`make deploy_local`

This command will deploy the project to the local environment, using `go run cmd/api/main.go` under the hood.
The first step will be to decrypt the `secrets/config.env.gpg` file, which can only be done if your GPG
key is listed in `.blackbox/blackbox-admins.txt` and has been used to encrypt the `.gpg` file itself.

The intention behind this is to allow for the secure storage of sensitive information necessary for
the functionality of the project in various deployment environments.

```shell
❯ make deploy_local
========== Importing keychain: START
gpg: WARNING: nothing exported
gpg: no valid OpenPGP data found.
gpg: Total number processed: 0
========== Importing keychain: DONE
========== Decrypting new/changed files: START
========== Decrypting new/changed files: DONE

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.6.1
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8000
```

Once the server is running, visit `http://localhost:8000` in your browser.


### AWS Lambda

Planned future work

### Kubernetes

Planned future work
