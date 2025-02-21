.PHONY: all install-deps generate clean

all: install-deps generate

# ----------------------------------------------------------------------------
# Installation targets
# ----------------------------------------------------------------------------
install-deps:
	@echo "installing protoc plugins and internal schema generator..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@master
	@go install github.com/lesomnus/proto-orm/cmd/protoc-gen-orm-ent@latest

# ----------------------------------------------------------------------------
# Codegen
# ----------------------------------------------------------------------------
generate: generate-proto generate-orm

generate-proto:
	@echo "compiling entity-related proto files..."
	@mkdir -p ./pkg/v1/ent
	@protoc -I proto \
		--go_out=. \
		--go_opt=module=frec.kr/tdoo \
		--go_opt=default_api_level=API_OPAQUE \
		--orm-ent_out=./pkg/v1/ent \
		--orm-ent_opt=module=frec.kr/tdoo/pkg/v1/ent \
		--orm-ent_opt=ent=frec.kr/tdoo/pkg/v1/ent/orm \
		proto/**/*.ent.proto

generate-orm:
	@echo "generating ORM codes by go ent..."
	@go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/v1/ent/schema \
		--target ./pkg/v1/ent/orm

# ----------------------------------------------------------------------------
# Cleaning targets
# ----------------------------------------------------------------------------
clean: clean-deps

clean-deps:
	@echo "cleaning installed binaries..."
	@rm -f ${GOPATH}/bin/protoc-gen-orm-ent
