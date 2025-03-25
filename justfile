
# run fmt imports, the run the binary
default *ARGS: fmt imports (run ARGS)

# clean up, clean up, everybody everywhere
clean: fmt imports tidy

# run the debug binary
run *ARGS:
    go run ./cmd/sm {{ ARGS }}

# build the release binary
build:
    go build -trimpath -ldflags="-s -w" ./cmd/sm

# build a binary for debugging
build-debug:
    go build ./cmd/sm

# run gofmt
fmt:
    go fmt ./cmd/sm
    go fmt ./internal/*

# run goimports
imports:
    goimports -w ./cmd/sm/
    goimports -w ./internal/*

tidy:
    go mod tidy -e
