export GOPROXY=https://goproxy.cn
rm -rf vendor
rm go.mod
rm go.sum
rm strategy

go mod init mynetwork/chaincode/subscriber/strategy
go mod vendor
go build