# go-cmd-boilerplate

A boilerplate to create command line utilities

# Technologies

* Docker
* Small lean command line tool
    * Compiled binary
* Client compatible instrumentation
    * Prometheus
    * Pushgateway
* Centralized configuration
* Structured logging
* Easy private package management (golang -> private github)
* Easy & quick tests that run during development

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

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
