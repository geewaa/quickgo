# 项目结构

  appname
    ├── cmd
    │   └── appnameapp
    │       ├── main.go
    │       └── wire
    │          ├── wire.go
    │          └── wire_gen.go
    ├── config
    │   ├── local.yml
    │   └── prod.yml
    ├── internal
    │   ├── app
    │   │   ├── app.go
    │   │   └── appname_app.go
    │   ├── model
    │   │   └── appname.go
    │   └── repository
    │       ├── repository.go
    │       └── appname.go
    ├── pkg
    │   ├── config
    │   │   └── config.go
    │   └── log
    │       └── log.go
    ├── README.md
    ├── go.mod
    └── go.sum

