set export

PROTO_DIR := "./protos"
GEN_DIR := "./gen"
WORKSPACE_DIR := invocation_directory()

[private]
default:
  @just --list

[private]
clean:
  find $WORKSPACE_DIR/$GEN_DIR -name '*.pb.go' -type f -delete

generate: clean
  @echo "Generating GRPC libraries..."
  buf generate $WORKSPACE_DIR/$PROTO_DIR -o $WORKSPACE_DIR/$GEN_DIR \
    --template $WORKSPACE_DIR/$GEN_DIR/buf.gen.yaml

format:
  @echo "Formatting GRPC libraries..."
  buf format --write $WORKSPACE_DIR/$PROTO_DIR

lint:
  @echo "Linting GRPC libraries..."
  buf lint $WORKSPACE_DIR/$PROTO_DIR

check-breaking:
  @echo "Checking for breaking changes..."
  buf breaking \
    --against "https://github.com/rsturla/platform-contracts.git#branch=main,subdir=$PROTO_DIR" \
    $WORKSPACE_DIR/$PROTO_DIR

fix-deps:
  #!/usr/bin/env bash
  echo "Fixing dependencies..."
  pushd $WORKSPACE_DIR/gen/go
  go mod tidy
  popd
