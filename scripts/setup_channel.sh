. scripts/utils.sh

echo '######## - (COMMON) setup variables - ########'
setupCommonENV

# echo '######## - (ORG1) create channel - ########'
setupProviderPeerENV
pushd ./channel-artifacts
if [[ "$CORE_PEER_TLS_ENABLED" == "true" ]]; then
    peer channel create -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ${CHANNEL_NAME}.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
else
    peer channel create -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ${CHANNEL_NAME}.tx
fi
popd

echo '######## - (ProviderOrg) join channel - ########'
setupProviderPeerENV
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

echo '######## - (ProviderOrg) update anchor - ########'
setupProviderPeerENV
if [[ "$CORE_PEER_TLS_ENABLED" == "true" ]]; then
    peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ./channel-artifacts/Org1MSPanchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
else
    peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ./channel-artifacts/Org1MSPanchors.tx
fi

echo '######## - (SubscriberOrg) join channel - ########'
setupSubscriberPeerENV2
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

echo '######## - (SubscriberOrg) update anchor - ########'
setupSubscriberPeerENV2
if [[ "$CORE_PEER_TLS_ENABLED" == "true" ]]; then
    peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ./channel-artifacts/Org2MSPanchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
else
    peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ./channel-artifacts/Org2MSPanchors.tx
fi

echo '######## - (RegulatorOrg) join channel - ########'
setupRegulatorPeerENV
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

echo '######## - (RegulatorOrg) update anchor - ########'
setupRegulatorPeerENV
if [[ "$CORE_PEER_TLS_ENABLED" == "true" ]]; then
    peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ./channel-artifacts/Org2MSPanchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
else
    peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ./channel-artifacts/Org2MSPanchors.tx
fi