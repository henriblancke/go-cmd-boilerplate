# go-cmd-boilerplate

A boilerplate to create command line utilities

# Technologies

* Docker
* Small lean command line tool (cobra)
    * Compiled binary
* Client compatible instrumentation
    * Prometheus
    * Pushgateway
* Centralized configuration (viper)
* Structured logging
* Easy private package management (golang -> private github)
* Easy & quick tests that run during development
* Sentry releases

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`. You'll have to create a copy of `config.yml` named `config-local.yml` for local development. Make any updates where necessary.

Running it then should be as simple as:

```console
$ make build
$ ./bin/go-cmd-boilerplate
```

### Spin up prometheus pushgateway
```
make start
```

### Testing

``make test``
