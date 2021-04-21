export GOPROXY=https://goproxy.cn
rm -rf vendor
rm go.mod
rm go.sum
rm strategy
rm ~/mynetwork/tmp/${CC_LABEL}.tar.gz

go mod init mynetwork/chaincode/strategy
go mod vendor
go build
