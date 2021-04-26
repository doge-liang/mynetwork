#!/bin/bash

. scripts/utils.sh

echo '######## - (COMMON) setup variables - ########'
setupVersionENV
setupCommonENV
sleep 1

export CC_NAME=mycc
INVOKE_PEER="Peer0.Subscriber"
INIT_FUNC="InitLedger"
OPERATE="query"
INVOKE_FUNC=""
INVOKE_FUNC_ARGS=""

export CC_NAME=$1
INVOKE_PEER=$2
OPERATE=$3
INVOKE_FUNC=$4
INVOKE_FUNC_ARGS=$5

echo "'CHAINCODE_NAME' set to '$CC_NAME'"
echo "'CHAINCODE_LANG' set to '$CC_LANG'"
echo "'CHAINCODE_PATH' set to '$CC_PATH'"
echo "'INVOKE_PEER' set to '$INVOKE_PEER'"
echo "'INVOKE_FUNC' set to '$INVOKE_FUNC'"
echo "'INVOKE_FUNC_ARGS' set to '$INVOKE_FUNC_ARGS'"

set -x
if [[ $INVOKE_PEER == "Peer.Subscriber" ]]; then
    setupSubscriberPeerENV
fi
if [[ $INVOKE_PEER == "Peer.Provider" ]]; then
    setupProviderPeerENV
fi
if [[ $INVOKE_PEER == "Peer.Regulator" ]]; then
    setupRegulatorPeerENV
fi
set +x
setGoCC

if [[ ${#INVOKE_FUNC} != 0 ]]; then
    if [[ $OPERATE == 'query' ]]; then
        echo '######## - ('$INVOKE_PEER') query chaincode - ########'
        set -x
        peer chaincode query \
            -C $CHANNEL_NAME \
        -n $CC_NAME \
        -c '{"Function":"'${INVOKE_FUNC}'", "Args":["'$INVOKE_FUNC_ARGS'"]}' | jq
        set +x
    fi

    if [[ $OPERATE == 'invoke' ]]; then
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
            -c '{"Function":"'$INVOKE_FUNC'","Args":["'$INVOKE_FUNC_ARGS'"]}'
        else
            peer chaincode invoke \
            -o ${ORDERER_ADDRESS} \
            -C $CHANNEL_NAME \
            -n ${CC_NAME}  \
            -c '{"Function":"'$INVOKE_FUNC'","Args":["'$INVOKE_FUNC_ARGS'"]}'
        fi
        set +x
    fi
fi