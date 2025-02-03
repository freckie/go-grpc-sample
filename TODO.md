# todo for the project

### developers' flow
1. Write entity proto files in `proto/tdoo`
   - entity definitions: `.ent.proto`
2. Generate Go Ent schema via `internal/cmd/i-proto-ent`
   - the prefix `i-` means "internal" cli tool
   - `i-proto-ent` generates schema files compatible with **_go ent_** framework
   - suggestions:
     - read `*.ent.proto` files from `proto/tdoo`
     - generate `*.gen.go` files into `pkg/v1/schema`
     - keep folder path of the input file?
3. Generate ORM code via `tdoo gen ent`
   - suggestions:
     - read `*.gen.go` files from `pkg/v1/schema`
     - generate ORM-related files into `pkg/v1/ent`
4. Write proto files for gRPC messages and services in `proto/tdoo`
   - message definitions: `.msg.proto`
   - service definitions: `.svc.proto`
5. Generate gRPC messages and services via `protoc-gen-go` and `protoc-gen-go-grpc`
   - read `*.msg.proto` and `*.svc.proto` files from `proto/tdoo`
   - generate `*.msg.pb.go`, `*.svc.pb.go` files into `pkg/v1/gen`

### todo
- [ ] `i-proto-ent`
  - [ ] some `.tpl` files
  - [ ] read proto file
  - [ ] generate a single entity
    - [ ] fields
    - [ ] enum
  - [ ] generate multiple entities with edges
- [ ] `tdoo`
  - [ ] `tdoo generate ent`
  - [ ] `tdoo generate proto`
  - [ ] `tdoo install`
    - [ ] check if all dependencies installed
      - `protoc-gen-go` & `protoc-gen-go-grpc`
      - et cetera.
