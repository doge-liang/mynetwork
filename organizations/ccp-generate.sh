#!/bin/bash

function one_line_pem {
    # echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
    echo "$(awk 'NF {sub(/\\n/, ""); printf "%s\\\\n",$0;}' $1)"
}

function json_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG_U}/$1/" \
        -e "s/\${ORG_L}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${CAPORT}/$4/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG_U}/$1/" \
        -e "s/\${ORG_L}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${CAPORT}/$4/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

ORG_U=Provider
ORG_L=provider
P0PORT=6001
CAPORT=9201
PEERPEM=organizations/peerOrganizations/provider.mynetwork.com/tlsca/tlsca.provider.mynetwork.com-cert.pem
CAPEM=organizations/peerOrganizations/provider.mynetwork.com/ca/ca.provider.mynetwork.com-cert.pem

echo "$(json_ccp $ORG_U $ORG_L $P0PORT $CAPORT $PEERPEM $CAPEM)" >organizations/peerOrganizations/provider.mynetwork.com/connection-provider.json
echo "$(yaml_ccp $ORG_U $ORG_L $P0PORT $CAPORT $PEERPEM $CAPEM)" >organizations/peerOrganizations/provider.mynetwork.com/connection-provider.yaml

ORG_U=Subscriber
ORG_L=subscriber
P0PORT=6003
CAPORT=9202
PEERPEM=organizations/peerOrganizations/subscriber.mynetwork.com/tlsca/tlsca.subscriber.mynetwork.com-cert.pem
CAPEM=organizations/peerOrganizations/subscriber.mynetwork.com/ca/ca.subscriber.mynetwork.com-cert.pem

echo "$(json_ccp $ORG_U $ORG_L $P0PORT $CAPORT $PEERPEM $CAPEM)" >organizations/peerOrganizations/subscriber.mynetwork.com/connection-subscriber.json
echo "$(yaml_ccp $ORG_U $ORG_L $P0PORT $CAPORT $PEERPEM $CAPEM)" >organizations/peerOrganizations/subscriber.mynetwork.com/connection-subscriber.yaml

ORG_U=Regulator
ORG_L=regulator
P0PORT=6005
CAPORT=9203
PEERPEM=organizations/peerOrganizations/regulator.mynetwork.com/tlsca/tlsca.regulator.mynetwork.com-cert.pem
CAPEM=organizations/peerOrganizations/regulator.mynetwork.com/ca/ca.regulator.mynetwork.com-cert.pem

echo "$(json_ccp $ORG_U $ORG_L $P0PORT $CAPORT $PEERPEM $CAPEM)" >organizations/peerOrganizations/regulator.mynetwork.com/connection-regulator.json
echo "$(yaml_ccp $ORG_U $ORG_L $P0PORT $CAPORT $PEERPEM $CAPEM)" >organizations/peerOrganizations/regulator.mynetwork.com/connection-regulator.yaml
