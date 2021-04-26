#!/bin/bash
. ./scripts/utils.sh

echo '######## - (COMMON) setup variables - ########'
setupVersionENV
setupCommonENV
export CC_NAME=mycc
export CC_VERSION=v1.0
export CC_SEQ=1
export CC_POLICY="OR('ProviderMSP.peer', 'SubscriberMSP.peer')"
export CC_LIFECYCLE="DEPLOY"

export CC_LIFECYCLE=$1
export CC_VERSION=$2
export CC_NAME=$3

setGoCC

# 加载私有数据配置文件
echo ${CC_PATH}/../collections_config.json

if [[ -f ${CC_PATH}/../collections_config.json ]]; then
    export PRIVATE_COLLECTION_DEF="--collections-config ${CC_PATH}/../collections_config.json"
fi

export CC_LABEL=${CC_NAME}_${CC_VERSION}

if [[ "$CC_LANG" == "java" ]]; then
    export CC_PATH=$CC_PATH/build/libs
fi

orgList=(Subscriber Provider Regulator)

for orgName in ${orgList[@]}; do

    echo "'CHAINCODE_NAME' set to '$CC_NAME'"
    echo "'CHAINCODE_LANG' set to '$CC_LANG'"
    echo "'CHAINCODE_PATH' set to '$CC_PATH'"
    echo "'CHAINCODE_VERSION' set to '$CC_VERSION'"
    echo "'CHAINCODE_LABEL' set to '$CC_LABEL'"
    echo "'SEQUENCE' set to '$CC_SEQ'"
    echo "'PRIVATE_COLLECTION_DEFINITION' set to '${PRIVATE_COLLECTION_DEF}'"

    if [[ ! -f tmp/${CC_LABEL}.tar.gz ]]; then
        pushd $CC_PATH
        ./build.sh
        popd
    fi
    echo '######## - ('$orgName') install chaincode - ########'
    case $orgName in
    Subscriber)
        setupSubscriberPeerENV
        ;;
    Provider)
        setupProviderPeerENV
        ;;
    Regulator)
        setupRegulatorPeerENV
        ;;
    esac
    setGoCC
    set -x
    if [[ ! -f tmp/${CC_LABEL}.tar.gz ]]; then
        peer lifecycle chaincode package tmp/${CC_LABEL}.tar.gz --path ${CC_PATH} --lang $CC_LANG --label ${CC_LABEL}
    fi
    peer lifecycle chaincode install tmp/${CC_LABEL}.tar.gz
    set +x
done

for orgName in ${orgList[@]}; do
    echo '######## - ('$orgName') install chaincode - ########'
    case $orgName in
    Subscriber)
        setupSubscriberPeerENV
        ;;
    Provider)
        setupProviderPeerENV
        ;;
    Regulator)
        setupRegulatorPeerENV
        ;;
    esac
    PACKAGE_ID=$(peer lifecycle chaincode queryinstalled --output json | jq -r '.installed_chaincodes[] | select(.label == env.CC_LABEL) | .package_id')
    echo "PACKAGE_ID('$orgName'):" ${PACKAGE_ID}
    if [[ "$CORE_PEER_TLS_ENABLED" == "true" ]]; then
        peer lifecycle chaincode approveformyorg \
            -o ${ORDERER_ADDRESS} \
            --ordererTLSHostnameOverride orderer.mynetwork.com \
            --tls $CORE_PEER_TLS_ENABLED \
            --cafile $ORDERER_CA \
            --channelID $CHANNEL_NAME \
            --name ${CC_NAME} \
            --version ${CC_VERSION} \
            --init-required \
            --package-id ${PACKAGE_ID} \
            --sequence $CC_SEQ \
            --waitForEvent \
            --signature-policy "$CC_POLICY" \
            $PRIVATE_COLLECTION_DEF
    else
        peer lifecycle chaincode approveformyorg \
            -o ${ORDERER_ADDRESS} \
            --channelID $CHANNEL_NAME \
            --name ${CC_NAME} \
            --version ${CC_VERSION} \
            --init-required \
            --package-id ${PACKAGE_ID} \
            --sequence $CC_SEQ \
            --waitForEvent \
            --signature-policy "$CC_POLICY" \
            $PRIVATE_COLLECTION_DEF
    fi
    set +x
done

echo '######## - (Provider) check chaincode approvals - ########'
setupProviderPeerENV
set -x
peer lifecycle chaincode checkcommitreadiness \
    --channelID $CHANNEL_NAME \
    --name ${CC_NAME} \
    --version ${CC_VERSION} \
    --sequence $CC_SEQ \
    --output json \
    --init-required \
    --signature-policy "$CC_POLICY" \
    $PRIVATE_COLLECTION_DEF
set +x

echo '######## - (Provider) commit chaincode definition - ########'
setupProviderPeerENV
set -x
if [[ "$CORE_PEER_TLS_ENABLED" == "true" ]]; then
    peer lifecycle chaincode commit \
        -o ${ORDERER_ADDRESS} \
        --ordererTLSHostnameOverride orderer.mynetwork.com \
        --tls $CORE_PEER_TLS_ENABLED \
        --cafile $ORDERER_CA \
        --peerAddresses $PEER0_PROVIDER_ADDRESS \
        --tlsRootCertFiles $PEER0_PROVIDER_TLS_ROOTCERT_FILE \
        --peerAddresses $PEER0_SUBSCRIBER_ADDRESS \
        --tlsRootCertFiles $PEER0_SUBSCRIBER_TLS_ROOTCERT_FILE \
        -C $CHANNEL_NAME \
        --name ${CC_NAME} \
        --version ${CC_VERSION} \
        --sequence $CC_SEQ \
        --init-required \
        --signature-policy "$CC_POLICY" \
        $PRIVATE_COLLECTION_DEF
else
    peer lifecycle chaincode commit -o ${ORDERER_ADDRESS} \
        --peerAddresses $PEER0_PROVIDER_ADDRESS \
        --peerAddresses $PEER0_SUBSCRIBER_ADDRESS \
        -C $CHANNEL_NAME \
        --name ${CC_NAME} \
        --version ${CC_VERSION} \
        --sequence $CC_SEQ \
        --init-required \
        --signature-policy "$CC_POLICY" \
        $PRIVATE_COLLECTION_DEF
fi
set +x
echo '######## - (Provider) check chaincode status - ########'
setupProviderPeerENV
peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name ${CC_NAME}
echo '############# END ###############'
