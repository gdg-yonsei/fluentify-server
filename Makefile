# Define where the *.proto files are located.
PROTO_DIR = idl/proto

# Define the output directory in which to place the build results.
PROTO_GEN_DIR = gen

CURR_DIR = $(PWD)

generate:
	docker run --rm \
	--volume $(CURR_DIR):/workspace \
	--workdir /workspace \
	bufbuild/buf generate

clean:
	rm -rf $(PROTO_GEN_DIR)
