#!/bin/sh

lsb_release=$(lsb_release -cs)

# 安装工具包
apt install -y apt-transport-https ca-certificates software-properties-common 
apt install -y unzip git  curl wget vim tree jq
# 安装 gradle 项目构建工具
cd /tmp
wget https://services.gradle.org/distributions/gradle-6.4-bin.zip
unzip gradle-6.4-bin.zip -d ./gradle-6.4
mv gradle-6.4 /usr/local/gradle
cat >> ~/.bashrc <<EOF
# setup gradle environments   配置gradle环境
# =====================
export PATH=$PATH:/usr/local/gradle/bin
# =====================
EOF
source ~/.bashrc

git clone https://gitlab.com/qubing/blockchain_lab_v2.git ~/workspace

# 安装 docker
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $lsb_release stable"
sudo apt-get update & sudo apt-get install -y docker-ce
docker -v
docker images
sudo gpasswd -a ${USER} docker
docker images
curl -sSL https://get.daocloud.io/daotools/set_mirror.sh | sh -s http://f1361db2.m.daocloud.io

sudo systemctl daemon-reload 
sudo systemctl restart docker

# 安装 docker-compose
wget https://github.com/docker/compose/releases/download/1.25.3/docker-compose-`uname -s`-`uname -m`
mv docker-compose-`uname -s`-`uname -m` /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
docker-compose -v

# 安装 Go
cd /tmp && wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
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
