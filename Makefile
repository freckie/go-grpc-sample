.PHONY: all install-deps generate clean

all: install-deps generate

# ----------------------------------------------------------------------------
# Installation targets
# ----------------------------------------------------------------------------
install-deps:
	@echo "installing protoc plugins..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@master
	@go install github.com/lesomnus/proto-orm/cmd/protoc-gen-orm-ent@latest
	@go install github.com/yoheimuta/protolint/cmd/protolint@latest
	@go install github.com/yoheimuta/protolint/cmd/protoc-gen-protolint@latest

# ----------------------------------------------------------------------------
# Codegen
# ----------------------------------------------------------------------------
generate: protolint generate-proto generate-orm generate-common generate-grpc

protolint:
	@protoc -I proto \
		--protolint_out=. \
		--protolint_opt=fix,proto_root=proto,config_dir_path=. \
		proto/**/*.proto

generate-proto:
	@echo "compiling entity-related proto files..."
	@mkdir -p ./pkg/v1/gen
	@protoc -I proto \
		--go_out=. \
		--go_opt=module=frec.kr/tdoo \
		--go_opt=default_api_level=API_OPAQUE \
		--orm-ent_out=./pkg/v1/gen \
		--orm-ent_opt=module=frec.kr/tdoo/pkg/v1/gen \
		--orm-ent_opt=ent=frec.kr/tdoo/pkg/v1/gen/tdoo/orm \
		proto/**/*.ent.proto

generate-orm:
	@echo "generating ORM codes by go ent..."
	@go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/v1/gen/tdoo/schema \
		--target ./pkg/v1/gen/tdoo/orm

generate-common:
	@echo "compiling global proto files..."
	@protoc -I proto \
		--go_out=. \
		--go_opt=module=frec.kr/tdoo \
		--go-grpc_out=. \
		--go-grpc_opt=module=frec.kr/tdoo \
		proto/global/common/*.proto

generate-grpc:
	@echo "generating gRPC"
	@protoc -I proto \
		--go_out=. \
		--go_opt=module=frec.kr/tdoo \
		--go-grpc_out=. \
		--go-grpc_opt=module=frec.kr/tdoo \
		proto/**/*.svc.proto

# ----------------------------------------------------------------------------
# Cleaning targets
# ----------------------------------------------------------------------------
clean: clean-generated

clean-generated:
	@echo "cleaning generated files..."
	@rm -r ./pkg/v1/gen

clean-deps:
	@echo "cleaning installed binaries..."
	@rm -f ${GOPATH}/bin/protoc-gen-orm-ent
