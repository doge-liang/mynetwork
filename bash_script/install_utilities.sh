#!/bin/sh

lsb_release=$(lsb_release -cs)

# 安装工具包
sudo apt install -y apt-transport-https ca-certificates software-properties-common 
sudo apt install -y unzip git  curl wget vim tree jq
# 安装 gradle 项目构建工具
cd /tmp
sudo wget https://services.gradle.org/distributions/gradle-6.4-bin.zip
unzip gradle-6.4-bin.zip -d ./gradle-6.4
sudo mv gradle-6.4 /usr/local/gradle
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


