mkdir cotterpin
cd cotterpin
go mod init cotterpin
cobra init --pkg-name cotterpin

## Notes

### Logging

Logging command > https://droctothorpe.github.io/posts/2020/07/leveled-logs-with-cobra-and-logrus/

```shell
➜ go run main.go init -d
INFO[0000] Debug logs enabled
Error: must also specify a resource
```

## Goreleaser

For local snapshots

```shell
goreleaser --snapshot --rm-dist
```
