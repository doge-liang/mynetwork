. scripts/utils.sh

echo '######## - (COMMON) setup variables - ########'
setupCommonENV
export CC_NAME=mycc
INVOKE_PEER="Peer0.Subscriber"
INIT_FUNC="InitLedger"
INVOKE_FUNC=""
INVOKE_FUNC_ARGS=""

if [[ $# -ge 1 ]]; then
    export CC_NAME=$1
fi
if [[ $# -ge 2 ]]; then
    INIT_FUNC=$2
fi
if [[ $# -ge 3 ]]; then
    INVOKE_PEER=$3
fi
if [[ $# -ge 4 ]]; then
    INVOKE_FUNC=$4
fi
if [[ $# -ge 5 ]]; then
    INVOKE_FUNC_ARGS=$5
fi

echo "'CHAINCODE_NAME' set to '$CC_NAME'"
echo "'CHAINCODE_LANG' set to '$CC_LANG'"
echo "'CHAINCODE_PATH' set to '$CC_PATH'"

if [[ $INVOKE_PEER == "Peer0.Subscriber" ]]; then
    setupSubscriberPeerENV0
fi
if [[ $INVOKE_PEER == "Peer1.Subscriber" ]]; then
    setupSubscriberPeerENV1
fi
if [[ $INVOKE_PEER == "Peer.Provider" ]]; then
    setupProviderPeerENV
fi
if [[ $INVOKE_PEER == "Peer.Regulator" ]]; then
    setupRegulatorPeerENV
fi

echo '######## - ('$INVOKE_PEER') init chaincode - ########'
set -x
if [[ "$CORE_PEER_TLS_ENABLED" == "true" ]]; then
    peer chaincode invoke \
    -o ${ORDERER_ADDRESS} \
    --ordererTLSHostnameOverride orderer.mynetwork.com \
    --tls $CORE_PEER_TLS_ENABLED \
    --cafile $ORDERER_CA \
    -C $CHANNEL_NAME \
    -n ${CC_NAME}  \
    --isInit -c '{"Function":"'$INIT_FUNC'","Args":[]}'
else
    peer chaincode invoke \
    -o ${ORDERER_ADDRESS} \
    -C $CHANNEL_NAME \
    -n ${CC_NAME}  \
    --isInit -c '{"Function":"'$INIT_FUNC'","Args":[]}'
fi
set +x
sleep 10

if [[ ${#INVOKE_FUNC} != 0 ]]; then
    echo '######## - ('$INVOKE_PEER') query chaincode - ########'
    set -x
    peer chaincode query -C $CHANNEL_NAME \
    -n $CC_NAME \
    -c '{"Function":"'$INVOKE_FUNC'", "Args":['$INVOKE_FUNC_ARGS']}'
    set +x
fi