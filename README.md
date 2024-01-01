# Simple sync application for two dirs
## Sync data from main dir to sub dir

### build your application
```bash
go build ./cmd/app/main.go
```
### run your application
```bash
./main -main-dir='path_to_main_dir' -sub-dir='path_to_sub_dir'
```