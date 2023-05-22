![logo](https://github.com/infratographer/website/blob/main/source/theme/assets/pictures/logo.jpg?raw=true)
# Example API

This repo is a little bit different than others found in the Infratpgarpher ecosystem. It''s meant to help understand through a simple example how to use [ent](https://entgo.io) and [gqlgen](https://gqlgen.com) to build a [GraphQL](https://graphql.org) API.

There are two main components to this repo:

1. The [ent](https://entgo.io) schema and code generation for the database
1. The [gqlgen](https://gqlgen.com) schema and code generation for the API

You can find a more detailed explanation of how to use `ent` and `gqlgen`  at their  examples;

* [`ent` tutorial](https://entgo.io/docs/tutorial-setup)
* [`gqlgen` getting started](https://gqlgen.com/getting-started/)

What you'll find here is a more practical example based on the `ent` tutorial. The `ent` tutorial is great, however; this just puts it all together for those who learn better from poking around an example. This example also starts to show how the Infratographer ecosystem can be used to build an application.


## Prerequisites

This example expects a working knowledge of `go`. If you are new to `go` you can find a great tutorial [here](https://tour.golang.org/welcome/1). This also assumes a rudimentary understanding of `graphql`. If you are new to `graphql` you can find a great tutorial [here](https://graphql.org/learn/).

You will also need to install the following:

* [Docker](https://docs.docker.com/get-docker/)
* [Go](https://golang.org/doc/install)

If you use devcontainers, you can use the `.devcontainer` directory to get started. You will need to install the following:

* [Docker](https://docs.docker.com/get-docker/)
* [VSCode](https://code.visualstudio.com/download)
* [VSCode Remote Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

While these are not required, they are recommended as the `devcontainer ` captures a working development environment that will make it easier to get started.


## Getting started

You can start by copying the following files from the root of the repo:

1. `.devcontainer/` (optional)
1. `generate.go`
1. `.golangci.yaml` (optional)
1. `gqlgen.yml`
1. `gqlgenc.yml`
1. `go.mod`
1. `main.go`
1. `cmd/`

You will then provide a few files to get started by copying the following:

```bash
generate.go
internal/api/
└── doc.go    # This will ensure that gqlgen can write to this directory
internal/ent/
├── entc.go
├── schema
│   ├── doc.go
│   ├── id_prefixes.go  # this is specific to the Infratographer ecosystem, it contains prefixed ID prefiex strings owned by this API
│   └── todo.go
schema
└── .gitkeep   # This will ensure ent can write to this directory
```

Run `go generate ./...`. This will generate the `ent` schema and code. (The generate will exit in failure at this point, but that's ok.)

```bash
internal/ent/generated/
internal/api/
├── doc.go
├── entity.resolvers.go
├── ent.resolvers.go
├── gen_federation.go   # (generated, do not modify)
├── gen_models.go       # (generated, do not modify)
├── gen_server.go       # (generated, do not modify)
├── resolver.go         # (generated, but copy paste this to your own resolver.go)
├── tenant.resolvers.go
└── todo.resolvers.go
schema/
└──  ent.graphql        # (generated, do not modify)
```

The `gen_*` files are never modified by hand. The `*.resolver.go` files are where you will add your business logic. The `schema/ent.graphql` file is where ent will generate the schema for the API. You then add your own schema files to the `schema/` directory. For example:

```bash
schema
├── tenant.graphql
└── todo.graphql
```

Once you rerun `go generate ./...` you will see the following files generated:

```bash
schema.graphql
internal/testclient
├── gen_client.go   # (generated, do not modify)
└── gen_models.go   # (generated, do not modify)
```

The `schema.graphql` file is the schema for the API that globs `schema/*graphql` and adds bits for federation. The `internal/testclient/*.gql` directory contains the code to test the API. You will add the `queries` and `mutations` you want to use in your tests.
At this point, you can start taking a look at the `api_test` package to see how the `testclient` and `api` are used. It's important to remember that with GraphQL the client isn't hitting predefined endpoints, but making requests for connected nodes on the graph. This means that the `testclient` is a bit different than what you might be used to with REST APIs.

## Development and Contributing

* [Development Guide](docs/development.md)
* [Contributing](https://infratographer.com/community/contributing/)

## Code of Conduct

[Contributor Code of Conduct](https://infratographer.com/community/code-of-conduct/). By participating in this project you agree to abide by its terms.

## Contact

To contact the maintainers, please open a [GitHub Issue](https://github.com/infratographer/example-api/issues/new)

## License

[Apache 2.0](LICENSE)
