# 拉取需要的 docker 镜像
# image of ca
docker pull hyperledger/fabric-ca:1.4.6
# image of peer
docker pull hyperledger/fabric-peer:2.1.0
# image of orderer
docker pull hyperledger/fabric-orderer:2.1.0
# image of tools & utilities
docker pull hyperledger/fabric-tools:2.1.0
# image of Chaincode deployment for Programming Languages (Go | Java | Node.JS)
docker pull hyperledger/fabric-ccenv:2.1.0
docker pull hyperledger/fabric-javaenv:2.1.0
docker pull hyperledger/fabric-nodeenv:2.1.0
# image of Base-OS of Chaincode runtime
docker pull hyperledger/fabric-baseos:0.4.20
# image of coucddb (one NOSQL DB for ledger state)
docker pull hyperledger/fabric-couchdb:0.4.20

docker images

# 安装 docker-compose
wget https://github.com/docker/compose/releases/download/1.25.3/docker-compose-`uname -s`-`uname -m`
mv docker-compose-`uname -s`-`uname -m` /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
docker-compose -v

# 安装 Go
cd /tmp && wget https://dl.google.com/go/go1.14.12.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.14.12.linux-amd64.tar.gz
cat >> ~/.bashrc <<EOF
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/gopath
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
EOF
source ~/.bashrc
go version

# 安装 Java
apt update
apt install -y openjdk-8-jdk
java -version

# 安装 nvm
cd ~
git clone --branch v0.35.3 https://gitee.com/mirrors/nvm.git .nvm
cd .nvm
. nvm.sh
cat >> ~/.bashrc <<EOF
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
EOF
source ~/.bashrc
nvm --version
nvm install 10
node -v
npm -v

