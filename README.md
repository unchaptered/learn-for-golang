# Golang

Golang 공부

- ShellScript

```shell
go run -v $(go list ./...)
```

- Cmd

```cmd
for /f %i in ('go list ./...') do go run -v %i
```