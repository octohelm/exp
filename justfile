# 查看可用入口
[group('meta')]
default:
    @just --list --unsorted

# 查看可用入口
[group('meta')]
list:
    @just --list --unsorted

# 整理 Go 依赖
[group('env')]
[no-cd]
dep:
    go mod tidy

# 升级当前模块依赖
[group('env')]
[no-cd]
update:
    go get -u ./...

# 运行测试
[group('test')]
[no-cd]
test path='./...' *args:
    CGO_ENABLED=0 go test -count=1 -failfast {{ args }} {{ path }}

# 执行 go fix
[group('fmt')]
[no-cd]
fix path='./...':
    go fix {{ path }}

# 执行 gofumpt
[group('fmt')]
[no-cd]
fmt dir='.':
    go tool gofumpt -l -w {{ dir }}
