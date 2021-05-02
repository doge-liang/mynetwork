. scripts/utils.sh

echo '######## - (COMMON) setup variables - ########'
setupCommonENV

echo '######## - (ProviderOrg) create channel - ########'
setupProviderPeerENV
pushd ./channel-artifacts

set -x
if [[ "$CORE_PEER_TLS_ENABLED" == "true" ]]; then
        peer channel create -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ${CHANNEL_NAME}.tx \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
else
    peer channel create -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ${CHANNEL_NAME}.tx
fi
set +x
popd

echo '######## - (ProviderOrg) join channel - ########'
setupProviderPeerENV
set -x
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block
set +x

echo '######## - (ProviderOrg) update anchor - ########'
setupProviderPeerENV
set -x
if [[ "$CORE_PEER_TLS_ENABLED" == "true" ]]; then
    peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} \
    -f ./channel-artifacts/ProviderMSPanchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
else
    peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ./channel-artifacts/ProviderMSPanchors.tx
fi
set +x

echo '######## - (SubscriberOrg) join channel - ########'
setupSubscriberPeerENV
set -x
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block
set +x

echo '######## - (SubscriberOrg) update anchor - ########'
setupSubscriberPeerENV
set -x
if [[ "$CORE_PEER_TLS_ENABLED" == "true" ]]; then
    peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} \
    -f ./channel-artifacts/SubscriberMSPanchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
else
    peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} -f ./channel-artifacts/SubscriberMSPanchors.tx
fi
set +x

# echo '######## - (RegulatorOrg) join channel - ########'
# setupRegulatorPeerENV
# peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

# echo '######## - (RegulatorOrg) update anchor - ########'
# setupRegulatorPeerENV
# if [[ "$CORE_PEER_TLS_ENABLED" == "true" ]]; then
#     peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} \
#     -f ./channel-artifacts/RegulatorMSPanchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA
# else
#     peer channel update -o ${ORDERER_ADDRESS} -c ${CHANNEL_NAME} \
#     -f ./channel-artifacts/RegulatorMSPanchors.tx
# fi

