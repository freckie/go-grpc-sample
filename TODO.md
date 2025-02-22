# todo for the project

### developers' flow
1. Write proto files in `proto/tdoo`
   - entity definitions: `.ent.proto`
   - service definitions: `.svc.proto`
2. Compile proto files and generate go files with `make generate`.
   - `make generate-proto` generates entity codes in go.
   - `make generate-orm` generates ORM code by **_go ent_**.
   - `make generate-common` compiles comon proto files.
   - `make generate-grpc` generates gRPC stubs.
3. (WIP) now implement logics!

### todo
- [ ] `logrus` package for logging
- [v] sample entity
  - [v] task.ent.proto
  - [v] user.ent.proto
- [v] sample svc
  - [v] task.svc.proto
  - [v] user.svc.proto
- [v] orm generation
- [v] grpc generation
- [ ] Makefile
    - [v] install deps
    - [v] generate proto
    - [v] generate orm
    - [v] generate grpc
    - [ ] run
- [ ] gRPC implementation