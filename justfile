
# run fmt imports, the run the binary
default *ARGS: fmt imports (sm ARGS)

# clean up, clean up, everybody everywhere
clean: fmt imports tidy

# run the debug binary
s *ARGS:
    go run ./cmd/s {{ ARGS }}

install-s *ARGS:
    go install ./cmd/s {{ ARGS }}

sm *ARGS:
    go run ./cmd/sm {{ ARGS }}

install-sm *ARGS:
    go install ./cmd/sm {{ ARGS }}

# build the release binary
build:
    go build -trimpath -ldflags="-s -w" ./cmd/s
    go build -trimpath -ldflags="-s -w" ./cmd/sm

# build a binary for debugging
build-debug:
    go build ./cmd/s
    go build ./cmd/sm

# run gofmt
fmt:
    go fmt ./cmd/s
    go fmt ./cmd/sm
    go fmt ./internal/*

# run goimports
imports:
    goimports -w ./cmd/s/
    goimports -w ./cmd/sm/
    goimports -w ./internal/*

tidy:
    go mod tidy -e
