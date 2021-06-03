# Build Stage
FROM golang:1.15 AS build-stage

LABEL app="build-go-cmd-boilerplate"
LABEL REPO="https://github.com/henriblancke/go-cmd-boilerplate"

ENV PROJPATH=/go/src/github.com/henriblancke/go-cmd-boilerplate

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/henriblancke/go-cmd-boilerplate
WORKDIR /go/src/github.com/henriblancke/go-cmd-boilerplate

RUN make build-alpine

# Final Stage
FROM alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/henriblancke/go-cmd-boilerplate"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/go-cmd-boilerplate/bin

WORKDIR /opt/go-cmd-boilerplate/bin

COPY --from=build-stage /go/src/github.com/henriblancke/go-cmd-boilerplate/bin/go-cmd-boilerplate /opt/go-cmd-boilerplate/bin/
RUN chmod +x /opt/go-cmd-boilerplate/bin/go-cmd-boilerplate

# Create appuser
RUN adduser -D -g '' go-cmd-boilerplate
USER go-cmd-boilerplate

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/go-cmd-boilerplate/bin/go-cmd-boilerplate"]
