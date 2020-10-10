-include config-new
REST_MAIN := "$(CURDIR)/cmd/rest"
BIN_REST := "$(CURDIR)/bin/rest"

.PHONY: prepare

prepare: clean init fetch vendor

init:
	@go mod init SkeletonAPI; \
	mkdir -p $(CURDIR)/temp

fetch:
	@go mod tidy

vendor:
	@go mod vendor

build-rest:
	@go build -i -v -o $(BIN_REST) $(REST_MAIN)

build-rest-vendor:
	@go build -mod=vendor -ldflags="-w -s" -o $(BIN_REST) $(REST_MAIN)


build: build-rest

build-vendor: build-rest-vendor

run-rest:
	@go run $(CURDIR)/cmd/rest/main.go

deploy: init build-vendor

clean:
	@rm -f $(CURDIR)/go.mod $(CURDIR)/go.sum \
	@rm -rf $(CURDIR)/bin $(CURDIR)/temp $(CURDIR)/vendor