
.PHONY: build build-alpine clean test help default

org := henriblancke
current_dir := $(shell pwd)
project := $(notdir $(current_dir))

gitsha := $(shell git rev-parse HEAD)
gitsha_last := $(shell git rev-parse HEAD~1)
gitsha_dirty=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)

branch_name := $(shell git rev-parse --abbrev-ref HEAD)
branch_clean := $(shell git rev-parse --abbrev-ref HEAD | sed 's/^\///;s/\///g')

image_name := $(shell git remote show origin | grep -e 'Push.*URL.*github.com' | rev | cut -d '/' -f 1 | rev | cut -d '.' -f 1)

build_timestamp := $(shell date '+%Y-%m-%d-%H:%M:%S')

version := $(shell git rev-list --count $(branch_name))


default: test

help:
	@echo 'Management commands for go-cmd-boilerplate:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs dep ensure, mostly used for ci.'
	@echo '    make build-alpine    Compile optimized for alpine linux.'
	@echo '    make package         Build final docker image with just the go binary inside'
	@echo '    make tag             Tag image created by package with latest, git commit and version'
	@echo '    make test            Run tests on a compiled project.'
	@echo '    make push            Push tagged images to registry'
	@echo '    make clean           Clean the directory tree.'
	@echo

define create_volume
$(call delete_volume,$(1))
docker volume create $(1)
endef

define delete_volume
$(call unmount_volume,$(1))
docker volume rm $(1) > /dev/null 2>&1 || true
endef

define mount_volume
docker create --mount source=$(1),target=$(2) --name $(1) alpine:3.8 /bin/true
endef

define unmount_volume
docker rm -f $(1) > /dev/null 2>&1 || true
endef

build:
	@echo "building ${image_name} ${version}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/henriblancke/go-cmd-boilerplate/version.GitCommit=${gitsha}${gitsha_dirty} -X github.com/${org}/${image_name}/version.BuildDate=${build_timestamp}" -o bin/${image_name}

start:
	docker compose up

stop:
	docker compose down

get-deps:
	dep ensure

build-alpine:
	@echo "building ${image_name} ${version}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static" -X github.com/${org}/${image_name}/version.GitCommit=${gitsha}${gitsha_dirty} -X github.com/${org}/${image_name}/version.BuildDate=${build_timestamp}' -o bin/${image_name}

package:
	@echo "building image ${image_name} ${version} $(gitsha)"
	docker build --build-arg VERSION=${version} --build-arg GIT_COMMIT=$(gitscha) -t $(image_name):local .

tag: 
	@echo "Tagging: latest ${version} $(gitsha)"
	docker tag $(org)/$(image_name):local  $(org)/$(image_name):${version}
	docker tag $(org)/$(image_name):local  $(org)/$(image_name):latest

push: tag
	@echo "Pushing docker image to registry: latest ${VERSION} $(GIT_COMMIT)"
	docker push $(org)/$(image_name):${version}
	docker push $(org)/$(image_name):latest

clean:
	@test ! -e bin/${image_name} || rm bin/${image_name}

test:
	go test ./...

sentry:
	$(call create_volume,$(image_name)-sentry)
	$(call mount_volume,$(image_name)-sentry,/work)
	docker cp . $(image_name)-sentry:/work
	docker run --rm \
		--volumes-from $(image_name)-sentry \
		-e SENTRY_AUTH_TOKEN \
		getsentry/sentry-cli releases \
			-o twine-labs \
			-p $(image_name) \
			set-commits \
			-c "twineteam/$(image_name)@$(gitsha_last)..$(gitsha)" \
			$(image_name)@$(version)
	$(call delete_volume,$(image_name)-sentry)
