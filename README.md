# go-grpc-tmpl

- This repository is a template for gRPC API project in Go.
- he project name is temporarily set as `tdoo`, meaning todo.

## Folder structure
```
./
├── proto
│   └── README.md
├── pkg
│   └── v1
│   │   ├── service
│   │   ├── repository
│   │   ├── models
│   │   └── handler
│   │       ├── http
│   │       └── grpc
├── log
├── internal
│   └── db
├── cmd
│   ├── tdoo
│   │   ├── main.go
│   │   ├── config
│   │   └── cmd
│   └── config
├── README.md
├── Makefile
└── Dockerfile
```