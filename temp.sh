setupProviderPeerENV

# 链码名
export CC_NAME=strategy
# 链码版本
export CC_VERSION=v1.0
# 链码序列号
export CC_SEQ=1
# 链码策略
# export CC_POLICY="OR('ProviderMSP.peer', 'SubscriberMSP.peer', 'RegulatorMSP.peer')"
export CC_POLICY="OR('ProviderMSP.peer')"
# 可以不设置,自己用来过滤脚本用的
export CC_LIFECYCLE="DEPLOY"
export CC_LABEL=${CC_NAME}_${CC_VERSION}
# 设置 Go 链码的变量
setGoCC

# 检查是否配置了私有数据集合配置文件
if [[ -f ${CC_PATH}/../../collections_config.json ]]; then
	export PRIVATE_COLLECTION_DEF="--collections-config ${CC_PATH}/../../collections_config.json"
fi

pushd $CC_PATH
./build.sh
popd

sleep 5

set -x
rm tmp/${CC_LABEL}.tar.gz

peer lifecycle chaincode package tmp/${CC_LABEL}.tar.gz --path ${CC_PATH} --lang $CC_LANG --label ${CC_LABEL}

set +x

setupProviderPeerENV
peer lifecycle chaincode install tmp/${CC_LABEL}.tar.gz

sleep 5

PACKAGE_ID=$(peer lifecycle chaincode queryinstalled --output json | jq -r '.installed_chaincodes[] | select(.label == env.CC_LABEL) | .package_id')
echo "PACKAGE_ID('$ORGANIZATION_NAME'):" ${PACKAGE_ID}

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

setupSubscriberPeerENV1

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

setupProviderPeerENV

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

peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name ${CC_NAME}

######################################################################

sleep 5

peer chaincode invoke \
	-o ${ORDERER_ADDRESS} \
	--ordererTLSHostnameOverride orderer.mynetwork.com \
	--tls $CORE_PEER_TLS_ENABLED \
	--cafile $ORDERER_CA \
	-C $CHANNEL_NAME \
	-n ${CC_NAME} \
	--isInit -c '{"Function":"","Args":[]}'

sleep 3

peer chaincode invoke \
	-o ${ORDERER_ADDRESS} \
	--ordererTLSHostnameOverride orderer.mynetwork.com \
	--tls $CORE_PEER_TLS_ENABLED \
	--cafile $ORDERER_CA \
	-C $CHANNEL_NAME \
	-n ${CC_NAME} \
	-c '{"Function":"'Init'","Args":[""]}'

# peer chaincode query -C $CHANNEL_NAME -n $CC_NAME -c '{"Function":"GetAllStrategies", "Args":[""]}' | jq

# peer chaincode query -C $CHANNEL_NAME -n $CC_NAME -c '{"Function":"ReadStrategy", "Args":["strategy3"]}' | jq

# peer chaincode query -C $CHANNEL_NAME -n $CC_NAME -c '{"Function":"ReadTrades", "Args":["strategy3"]}' | jq
