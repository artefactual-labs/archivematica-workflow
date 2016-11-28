# archivematica-workflow

**archivematica-workflow** is an effort to pull the Archivematica workflow data out of the database. One of the main goals is to translate its contents into multiple languages.

## Motivation

Ongoing discussion: https://goo.gl/8c4X7G.

Visit our [Transifex project](https://www.transifex.com/artefactual/archivematica/dashboard/) if you want to contribute as a translator.

## Workflow data extraction and translations

The `dist/default/default.json` file contains the JSON-encoded workflow data available in Archivematica (`qa/1.x` branch). It is generated with [`json-links`](https://github.com/artefactual/archivematica-devtools/commit/147141fec86a14dd0585129dfb99f48134142548) which is part of [archivematica-devtools](https://github.com/artefactual/archivematica-devtools). For example, the following command will encode the workflow database existing in your database:

    $ am json-links --locale-dir=dist/workflows/locale/default dist/workflows/default.json

When `--locale-dir` is used, predicatably, `json-links` does the following things:

- Populates the latest version of `en.json` so it can be pushed to Transifex (with `tx push -s).
- Populates the translation objects in `dist/workflows/default.json` using the translations available under the given directory. Make sure that the translations are up todate in advance with `tx pull -a`.

Transifex is our localization platform. Workflow messages can be found at https://www.transifex.com/artefactual/archivematica/mcp-workflow/.
