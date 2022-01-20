# about-me - Source code for trevor-edris.io

[![Go Report Card](https://goreportcard.com/badge/github.com/TrevorEdris/about-me)](https://goreportcard.com/report/github.com/TrevorEdris/about-me)
[![Test](https://github.com/TrevorEdris/about-me/actions/workflows/test.yml/badge.svg)](https://github.com/TrevorEdris/about-me/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoT](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev)

**A large majority of the base code for this project has been adapted from [https://github.com/mikestefanello/pagoda](https://github.com/mikestefanello/pagoda).** Many thanks to mikestefanello for a great template repository.

## Developing

### Requirements

* Docker
* Docker Compose

`make dev`

Use `make dev` to bring up the application with "Live reloading" enabled (via [cosmtrek/air](https://github.com/cosmtrek/air)). This will detect changes to files and automatically restart the application.

Included in the `docker-compose.dev.yml` is an SMTP server (via [maildev](https://github.com/maildev/maildev)) that allows for the SMTP functionality to be tested locally. To view the mail server, visit `http://localhost1080` in a browser.


```shell
❯ make dev
docker-compose -f docker-compose.dev.yml up -d
Creating network "about-me-network" with driver "bridge"
Creating about-me_mail_1 ... done
Creating api             ... done
make -s dev-logs
Attaching to api
api     |
api     |   __    _   ___
api     |  / /\  | | | |_)
api     | /_/--\ |_| |_| \_ , built with Go
api     |
api     | watching .
api     | watching cmd
api     | watching cmd/api
. . .
api     | building...
api     | running...
api     |
api     |    ____    __
api     |   / __/___/ /  ___
api     |  / _// __/ _ \/ _ \
api     | /___/\__/_//_/\___/ v4.6.1
api     | High performance, minimalist Go web framework
api     | https://echo.labstack.com
api     | ____________________________________O/_______
api     |                                     O\
api     | ⇨ http server started on [::]:8000
```

Once the server is running, visit `http://localhost:8000` in your browser.

Any change made to a file being "watched" by the `cosmtrek/air` container will trigger a rebuild and
rerun of the application, allowing for extremely fast development.

```shell
api     | embedded/link.go has changed
api     | building...
api     | running...
api     |
api     |    ____    __
api     |   / __/___/ /  ___
api     |  / _// __/ _ \/ _ \
api     | /___/\__/_//_/\___/ v4.6.1
api     | High performance, minimalist Go web framework
api     | https://echo.labstack.com
api     | ____________________________________O/_______
api     |                                     O\
api     | ⇨ http server started on [::]:8000
```

## Testing

### Requirements

* Docker
* Docker Compose

`make test`

This command will run the full test suite, running both unit tests and integration tests. The integration test
portion will also output the code coverage.

Files with the suffix `_test.go` will have build annotations at the top of the file, indicating whether they are
part of the unit test or integration test suite. The tests executed by `go test` are controlled via an additional
argument `-tags`, such as `-tags=unit` or `-tags=integration`.

```go
//go:build integration
// +build integration

. . .

//go:build unit
// +build unit
```

## Deploying

Deploying changes to the actual website is currently a manual process, as the website is hosted with AWS App Runner, however there are very few manual steps required.

1. `make publish`
2. Update the AWS App Runner project to point to the image in ECR that was just published

The `make publish` command will run through the whole pipeline, building the docker image, testing the project
(both unit and integration tests), creating a finalized docker image with just the runnable binary (to minimize
image size), and finally pushing the finalized image to the ECR registry.

## Tools

### `tools/version`

The `tools/version` script calculates the semantic version of the project based on multiple factors:

* The value in the `VERSION` file
* The number of commits since the `VERSION` file was last included in a commit
* The current branch
* The "cleanliness" of a current branch (uncommitted changes or not)

Running `make version` will output the current calculated version of the project. For example,
at the time of writing this, the project is on the `task-update-docs` branch with uncommitted changes
in the `README.md` and `Makefile`.

```shell
❯ make version
v0.0.14-task-update-docs-7621f2f-dirty
```

This version begins with the semantic version `v0.0.14` but is annotated with the branch name
`task-update-docs` as well as the most recent commit hash `7621f2f` and due to the uncommitted changes,
`dirty` is appended.

By stashing the current changes, the `dirty` annotation will be removed.

```shell
❯ git status
On branch task-update-docs
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   Makefile
        modified:   README.md

no changes added to commit (use "git add" and/or "git commit -a")
❯ git stash
Saved working directory and index state WIP on task-update-docs: 7621f2f Add Projects page (#23)
❯ make version
v0.0.14-task-update-docs-7621f2f
```

The `tools/version` script will only update the patch number of the semantic version. The major and minor
values are completely controlled via manual updates to the `VERSION` file.

To create `v0.1.0`, simply modify the contents of `VERSION` to `0.1.0` and push to the main branch. From
then on, the calculated semantic version will be `v0.1.1`, `v0.1.2`, ... and so on.
