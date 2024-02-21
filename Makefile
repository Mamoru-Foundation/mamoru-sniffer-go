BINARY_DIR = ${CURDIR}/packaged/lib/$$(go env GOOS)-$$(go env GOARCH)
LIB_NAME = libmamoru_sniffer_go.so
C-FOR-GO = c-for-go


dev: build-rust generate-go-bindings test
generate-go-bindings:
	$(C-FOR-GO) cforgo.yaml

build-rust:
	cargo build
	mkdir -p $(BINARY_DIR)
	cp target/debug/$(LIB_NAME) $(BINARY_DIR)/

build-rust-release:
	cargo build --release
	mkdir -p $(BINARY_DIR)
	cp target/release/$(LIB_NAME) $(BINARY_DIR)/

build-rust-release-macos-aarch64:
	cargo build --release --target aarch64-apple-darwin
	mkdir -p ./lib/darwin-arm64/
	cp target/aarch64-apple-darwin/release/$(LIB_NAME) ./packaged/lib/darwin-arm64/

test:
	@export LD_LIBRARY_PATH=$(BINARY_DIR):$$LD_LIBRARY_PATH ;  \
	GODEBUG=cgocheck=2 ; \
	go test ./mamoru_sniffer -v

bench:
	@export LD_LIBRARY_PATH=$(BINARY_DIR):$$LD_LIBRARY_PATH ;  \
	GODEBUG=cgocheck=2 ; \
	go test -bench=. ./mamoru_sniffer -v

# Requires manual setup go v1.21 or later
# 	GOEXPERIMENT=cgocheck2  go test ./tests -v
integration-test:
	export LD_LIBRARY_PATH=$(BINARY_DIR):$$LD_LIBRARY_PATH; \
	GODEBUG=cgocheck=2 ; \
	go test ./tests -v

clean:
	cargo clean
	rm -rf generated_bindings/*
