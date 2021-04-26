function setGoCC() {
    export CC_LANG=golang
    # export CC_PATH=${PWD}/chaincode/subscriber/${CC_NAME}
    # export CC_PATH=${PWD}/chaincode/${ORGANIZATION_NAME}/${CC_NAME}
    export CC_PATH=${PWD}/chaincode/${CC_NAME}
    # export CC_PATH=${PWD}/chaincode/${CC_NAME}/go
}

function setupVersionENV() {
    export GO_VERSION=1.16
    export DOCKER_VERSION=1.25.3
    export FABRIC_VERSION=2.3.0
    export CA_VERSION=1.4.9
    export DB_VERSION=3.1.1
}

function setupCommonENV() {
    export FABRIC_CFG_PATH=${PWD}/fabric-bin/${FABRIC_VERSION}/config
    export ORDERER_ADDRESS=localhost:6007
    export PEER0_PROVIDER_ADDRESS=localhost:6001
    export PEER0_SUBSCRIBER_ADDRESS=localhost:6003
    export PEER0_REGULATOR_ADDRESS=localhost:6005
    
    export PEER0_PROVIDER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/ca.crt
    export PEER0_SUBSCRIBER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/ca.crt
    export PEER0_REGULATOR_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/ca.crt
    
    export ORDERER_CA=${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/msp/tlscacerts/tlsca.orderer.mynetwork.com-cert.pem
    export CHANNEL_NAME=mychannel
}

function setupSubscriberPeerENV() {
    export CORE_PEER_LOCALMSPID=SubscriberMSP
    export ORGANIZATION_NAME=subscriber
    export CORE_PEER_ADDRESS=$PEER0_SUBSCRIBER_ADDRESS
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_TLS_CERT_FILE=${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/server.crt
    export CORE_PEER_TLS_KEY_FILE=${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/server.key
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/ca.crt
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/users/Admin@subscriber.mynetwork.com/msp
}

function setupProviderPeerENV() {
    export CORE_PEER_LOCALMSPID=ProviderMSP
    export ORGANIZATION_NAME=provider
    export CORE_PEER_ADDRESS=$PEER0_PROVIDER_ADDRESS
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_TLS_CERT_FILE=${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/server.crt
    export CORE_PEER_TLS_KEY_FILE=${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/server.key
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/ca.crt
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/provider.mynetwork.com/users/Admin@provider.mynetwork.com/msp
}

function setupRegulatorPeerENV() {
    export CORE_PEER_LOCALMSPID=RegulatorMSP
    export ORGANIZATION_NAME=regulator
    export CORE_PEER_ADDRESS=$PEER0_REGULATOR_ADDRESS
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_TLS_CERT_FILE=${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/server.crt
    export CORE_PEER_TLS_KEY_FILE=${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/server.key
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/ca.crt
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/users/Admin@regulator.mynetwork.com/msp
}
