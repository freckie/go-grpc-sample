# /internal/cmd/i-proto-ent

- `i-proto-ent` is an implementation of `protoc` plugin for generating **_go ent_** schema files from `*.ent.proto`.
- This tool is used only within the project.
  - the prefix `i-` means "internal" cli tool
- suggestions:
  - read `*.ent.proto` files in `proto/tdoo`
  - generate `*.gen.go` files into `pkg/v1/schema`
- usage: `protoc-gen-i-proto-ent`

- To use proto files as the SSoT, which stands for Single Source of Truth, I restricted manually defining entities in **_go ent_**.