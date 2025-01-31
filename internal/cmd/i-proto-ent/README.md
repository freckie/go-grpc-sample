# /internal/cmd/i-proto-ent

- `i-proto-ent` is a cli tool for generating **_go ent_** schema files from `*.ent.proto`.
- This tool is used only within the project.
    - the prefix `i-` means "internal" cli tool
- suggestions:
    - read `*.ent.proto` files in `proto/tdoo`
    - generate `*.gen.go` files into `pkg/v1/schema`