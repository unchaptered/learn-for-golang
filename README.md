# Golang

Golang 공부

## Docs

- [Microsoft - Take your first steps with Go](./src/train/README.md)

## Command

- ShellScript

```shell
go run -v $(go list ./...)
```

- Cmd

```cmd
for /f %i in ('go list ./...') do go run -v %i
```