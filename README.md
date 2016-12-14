# archivematica-workflow [![Go Report Card](https://goreportcard.com/badge/artefactual-labs/archivematica-workflow)](https://goreportcard.com/report/artefactual-labs/archivematica-workflow)

**archivematica-workflow** is a lightweight service part of Archivematica used to distribute Archivematica workflows and make them available over the network.

## Motivation

Ongoing discussion: https://goo.gl/8c4X7G.

Visit our [Transifex project](https://www.transifex.com/artefactual/archivematica/dashboard/) if you want to contribute in the translation of the Archivematica workflow data or in other components related to the Archivematica project.

## Building from sources

Install [Go 1.7 or newer](https://golang.org/dl/), set up your `$GOPATH` and run:

    $ go get -u github.com/artefactual-labs/archivematica-workflow/...

Binaries `archivematica-workflow` and `archivematica-workflow-cli` should now be available under `$GOPATH/bin` directory.

## Development notes

These are some tools that you should install for development purposes:

    $ go get -u github.com/kardianos/govendor
    $ go get -u github.com/jteeuwen/go-bindata/go-bindata
    $ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

You will also need [protoc 3.1 or newer](https://github.com/google/protobuf/releases), the protobuf compiler.

Note that we're using `govendor` to manage our Go dependencies but [`dep`] (https://docs.google.com/document/d/1qnmjwfMmvSCDaY4jxPmLAccaaUI5FfySNE90gB0pTKQ/edit) will be adopted as soon as it is released. `govendor` is very intuitive though (see [cheatsheet](https://github.com/kardianos/govendor/wiki/Govendor-CheatSheet)) - these are some of the things that you can do with it:

- Compile the sources: `govendor install +local`.
- Run tests: `govendor test +local`.
- List vendored dependencies: `govendor list +vendor`

## Workflows

The only workflow currently available is [`default.json`](dist/default.json). It has been generated with [`json-links`](https://github.com/artefactual/archivematica-devtools/commit/147141fec86a14dd0585129dfb99f48134142548) which is part of the `archivematica-devtools` repository.

**The schema used to encode the workflow data has not been specified yet. We're investigating better ways to do this, like linked data.**

Workflows are currently embedded into the final binary to simplify the distribution process. Further development may incorporate more sophisticated workflow management options.

The embedding process of the [`dist`](dist/) directory is done with `go-bindata`, which must be installed previously. [`pkg/service/dist/dist.go`](pkg/service/dist/dist.go) is the file invoking [`go:generate`](https://blog.golang.org/generate). The generated code is located at [`pkg/service/dist/assets.go`](pkg/service/dist/assets.go). The contents of `dist/` can be harvested using the generated API which you can explore with `go doc ./pkg/service/dist`. If you want to embed new data, please run:

    $ go generate ./pkg/service/dist

We check in the generated code in the repository so the application is "go gettable".

## Workflow data extraction and translations

The `dist/default/default.json` file contains the JSON-encoded workflow data available in Archivematica (`qa/1.x` branch). It is generated with [`json-links`](https://github.com/artefactual/archivematica-devtools/commit/147141fec86a14dd0585129dfb99f48134142548) which is part of [archivematica-devtools](https://github.com/artefactual/archivematica-devtools). For example, the following command will encode the workflow database existing in your database:

    $ am json-links --locale-dir=dist/workflows/locale/default dist/workflows/default.json

When `--locale-dir` is used, predicatably, `json-links` does the following things:

- Populates the latest version of `en.json` so it can be pushed to Transifex (with `tx push -s).
- Populates the translation objects in `dist/workflows/default.json` using the translations available under the given directory. Make sure that the translations are up todate in advance with `tx pull -a`.

Transifex is our localization platform. Workflow messages can be found at https://www.transifex.com/artefactual/archivematica/mcp-workflow/.
