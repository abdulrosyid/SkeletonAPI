REST_MAIN := "$(CURDIR)/cmd/rest"
BIN_REST := "$(CURDIR)/bin/rest"
MIGRATIONS_MAIN := "$(CURDIR)/cmd/migrations"
BIN_MIGRATIONS := "$(CURDIR)/bin/migrations"

install: clean init fetch build migrate

init:
	go mod init SkeletonAPI;

fetch:
	go mod tidy

migrate:
	$(BIN_MIGRATIONS)

run:
	$(BIN_REST)

build:
	go build -i -v -o $(BIN_REST) $(REST_MAIN); \
	go build -i -v -o $(BIN_MIGRATIONS) $(MIGRATIONS_MAIN);

clean:
	rm -f $(CURDIR)/go.mod $(CURDIR)/go.sum; \
	rm -rf $(CURDIR)/bin;