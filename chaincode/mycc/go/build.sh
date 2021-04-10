export GOPROXY=https://goproxy.cn
rm -rf vendor
rm go.mod
rm go.sum
rm example01

go mod init mycc
go mod vendor
go build