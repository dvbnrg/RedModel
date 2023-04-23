# Generate Protocol definitions, Test, and Build
default: fmt gen test
# Format
fmt:
    go fmt
    clang-format pb/*proto -i
# Gen
gen:
    protoc --go-grpc_out=pb --go_out=pb pb/*.proto
# Test
test:
    go test ./... -v
# Build
build:
    go build -o redproto *.go
# Remove build artifacts
clean:
    rm pb/*.pb.go
    rm redproto
# Build docker image
docker: gen
    docker build . -t redproto