#!/bin/bash
. ./setup.sh

# 把相关的配置文件放到正确的路径
echo "0.Initialize"
mkdir -p ${PWD}/tmp
mkdir -p organizations/fabric-ca/providerOrg
mkdir -p organizations/fabric-ca/subscriberOrg
mkdir -p organizations/fabric-ca/regulatorOrg
mkdir -p organizations/fabric-ca/ordererOrg
cp organizations/fabric-ca/ca.provider.mynetwork.com.yaml organizations/fabric-ca/providerOrg/fabric-ca-server-config.yaml
cp organizations/fabric-ca/ca.subscriber.mynetwork.com.yaml organizations/fabric-ca/subscriberOrg/fabric-ca-server-config.yaml
cp organizations/fabric-ca/ca.regulator.mynetwork.com.yaml organizations/fabric-ca/regulatorOrg/fabric-ca-server-config.yaml
cp organizations/fabric-ca/ca.orderer.mynetwork.com.yaml organizations/fabric-ca/ordererOrg/fabric-ca-server-config.yaml
echo

# 启动了四个 ca 服务，分别是 Orderer_CA, Provider_CA, Subscriber_CA, Regulator_CA
echo "1.Startup CA Services in Network"
CA_IMAGE_TAG=${CA_VERSION} docker-compose -f docker/docker-compose-ca.yaml up -d
echo

sleep 5

# 注册 Peer 和 Orderer 节点
echo "2.Register Peers and Orderer with users"
. organizations/fabric-ca/registerEnroll.sh 
createProviderOrg
createSubscriberOrg
createRegulatorOrg
createOrderer
echo

echo "3.Create orderer.genesis.block"
. scripts/utils.sh
setupCommonENV
# # 因为现在要配置系统通道了所以 FABRIC_CFG_PATH 要指向 configtx.yaml 所在的路径
export FABRIC_CFG_PATH=${PWD}/configtx
configtxgen -profile ThreeOrgsOrdererGenesis -channelID system-channel -outputBlock ./system-genesis-block/genesis.block
configtxgen -profile ThreeOrgsChannel -outputCreateChannelTx ./channel-artifacts/$CHANNEL_NAME.tx -channelID $CHANNEL_NAME
configtxgen -profile ThreeOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/ProviderOrgMSPanchors.tx -channelID $CHANNEL_NAME -asOrg ProviderOrgMSP
configtxgen -profile ThreeOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/SubscriberOrgMSPanchors.tx -channelID $CHANNEL_NAME -asOrg SubscriberOrgMSP
echo

echo "4.Startup Peers and Orderer"
COMPOSE_FILE_BASE=docker/docker-compose-ABC.yaml COMPOSE_FILE_COUCH=docker/docker-compose-couch.yaml IMAGE_TAG=${FABRIC_VERSION} DB_IMAGE_TAG=${DB_VERSION} docker-compose -f ${COMPOSE_FILE_BASE} -f ${COMPOSE_FILE_COUCH} up -d
echo

sleep 5

# echo "5.Create & Join Channel"
# . scripts/setup_channel.sh
# echo

# echo "6.Generate Connection Profiles"
# ./organizations/ccp-generate.sh
# if [ ! -d "${PWD}/app/example01_java/profiles/Org1/tls" ]; then 
#     mkdir -p app/example01_java/profiles/Org1/tls
# fi

# if [ ! -d "${PWD}/app/example01_java/profiles/Org2/tls" ]; then 
#     mkdir -p app/example01_java/profiles/Org2/tls
# fi

# if [ ! -d "${PWD}/app/example02/profiles/Org1/tls" ]; then 
#     mkdir -p app/example02_java/profiles/Org1/tls
# fi

# if [ ! -d "${PWD}/app/example02/profiles/Org2/tls" ]; then 
#     mkdir -p app/example02_java/profiles/Org2/tls
# fi

# cp ./organizations/peerOrganizations/org1.example.com/connection-org1.json app/example01_java/profiles/Org1/connection.json
# cp ./organizations/peerOrganizations/org2.example.com/connection-org2.json app/example01_java/profiles/Org2/connection.json
# cp ./organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem app/example01_java/profiles/Org1/tls/
# cp ./organizations/peerOrganizations/org2.example.com/ca/ca.org2.example.com-cert.pem app/example01_java/profiles/Org2/tls/

# cp ./organizations/peerOrganizations/org1.example.com/connection-org1.json app/example02_java/profiles/Org1/connection.json
# cp ./organizations/peerOrganizations/org2.example.com/connection-org2.json app/example02_java/profiles/Org2/connection.json
# cp ./organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem app/example02_java/profiles/Org1/tls/
# cp ./organizations/peerOrganizations/org2.example.com/ca/ca.org2.example.com-cert.pem app/example02_java/profiles/Org2/tls/

# echo

# echo "Done."
