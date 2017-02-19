# archivematica-workflow [![Travis CI](https://travis-ci.org/artefactual-labs/archivematica-workflow.svg?branch=master)](https://travis-ci.org/artefactual-labs/archivematica-workflow) [![Coverage Status](https://coveralls.io/repos/artefactual-labs/archivematica-workflow/badge.svg?branch=master&service=github)](https://coveralls.io/github/artefactual-labs/archivematica-workflow?branch=master) [![Go Report Card](https://goreportcard.com/badge/artefactual-labs/archivematica-workflow)](https://goreportcard.com/report/artefactual-labs/archivematica-workflow)

**archivematica-workflow** is a small service used to distribute Archivematica workflows within other Archivematica services. The current implementation is limited to a single workflow extracted from the the work-in-progress Archivematica 1.7.

This is the first step toward finding better ways to describe digital preservation workflows and decoupling workflow management from our application stack.

Ongoing discussion: https://goo.gl/8c4X7G.

## Usage

`archivematica-workflow` is distributed as a self-contained, statically-linked binary. You can download it from the [Releases](https://github.com/artefactual-labs/archivematica-workflow/releases) page or deploy it with the [Ansible role](https://github.com/artefactual-labs/ansible-archivematica-workflow).

The only currently available workflow can be found in [dist/workflows/default.json](dist/workflows/default.json). The HTTP API and gRPC API let you access to it. Visit the [proto](proto/workflow.json) file for more details.

## Translations

Visit our [Transifex project](https://www.transifex.com/artefactual/archivematica/dashboard/) if you want to contribute in the translation of the Archivematica workflow data or in other components related to the Archivematica project.

## Development

### Building from sources

Install [Go 1.7 or newer](https://golang.org/dl/) and run:

    $ go get -u github.com/artefactual-labs/archivematica-workflow/...

Binaries `archivematica-workflow` and `archivematica-workflow-cli` should now be available under `$GOPATH/bin` directory.

### Development tools

These are some Go tools that you will need at some point if you are planning to contribute to this project as a developer. You can install them by running:

    $ go get -u github.com/kardianos/govendor
    $ go get -u github.com/jteeuwen/go-bindata/go-bindata
    $ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

You will also need:

- [protoc 3.1 or newer](https://github.com/google/protobuf/releases), the protobuf compiler.
- [transifex-client](https://pypi.python.org/pypi/transifex-client), the Transifex CLI used to push our messages and pull the tranlations.

### Vendoring dependencies

This repo includes a copy of its dependencies under the `vendor/` directory. We use [govendor](https://github.com/kardianos/govendor) to work with these dependencies. We will likely adopt [dep](https://github.com/golang/dep) once it matures - at the time of writing, dep is still considered pre-alpha.

govendor is very intuitive though (see [cheatsheet](https://github.com/kardianos/govendor/wiki/Govendor-CheatSheet)) - these are some of the things that you can do with it:

- Compile the sources: `govendor install +local`.
- Run tests: `govendor test +local`.
- List vendored dependencies: `govendor list +vendor`

### Embedding workflows

Workflows are currently embedded into the final binary to simplify the distribution process. Further development may incorporate more sophisticated workflow management options.

The only currently available workflow is [`default.json`](dist/default.json). It has been generated with [`json-links`](https://github.com/artefactual/archivematica-devtools/commit/147141fec86a14dd0585129dfb99f48134142548) which is part of the `archivematica-devtools` repository.

**The schema used to encode the workflow data has not been specified yet. We're investigating better ways to do this, like linked data.**

Running the following command will generate new Go code with the contents of the [`dist`](dist/) directory embedded in binary format:

    $ go generate ./pkg/service/dist

The new code is checked in the repository so the application is "go gettable". In the future this step could be done as part of the build process.

### Extracting workflow data from Archivematica

There is an ongoing effort in Archivematica (pull request [#506](https://github.com/artefactual/archivematica/pull/506)) to read the workflow data using archivematica-workflow's API.

The `dist/default/default.json` file contains the JSON-encoded workflow data available in Archivematica (`qa/1.x` branch). It is generated with [`json-links`](https://github.com/artefactual/archivematica-devtools/commit/147141fec86a14dd0585129dfb99f48134142548) which is part of [archivematica-devtools](https://github.com/artefactual/archivematica-devtools). For example, the following command will encode the workflow database existing in your database:

    $ am json-links --locale-dir=dist/workflows/locale/default dist/workflows/default.json

When `--locale-dir` is used, predicatably, `json-links` does the following things:

- Populates the latest version of `en.json` so it can be pushed to Transifex (with `tx push -s).
- Populates the translation objects in `dist/workflows/default.json` using the translations available under the given directory. Make sure that the translations are up todate in advance with `tx pull -a`.

Transifex is our localization platform. Workflow messages can be found at https://www.transifex.com/artefactual/archivematica/mcp-workflow/.
