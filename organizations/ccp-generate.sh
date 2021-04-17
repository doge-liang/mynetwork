#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.json
}

function json_ccp_2 {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${CAPORT}/$4/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template-2.json
}

function yaml_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

function yaml_ccp_2 {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${P1PORT}/$3/" \
        -e "s/\${CAPORT}/$4/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template-2.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

ORG=Provider
P0PORT=7051
CAPORT=7054
PEERPEM=organizations/peerOrganizations/provider.mynetwork.com/tlsca/tlsca.provider.mynetwork.com-cert.pem
CAPEM=organizations/peerOrganizations/provider.mynetwork.com/ca/ca.provider.mynetwork.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/connection-provider.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/connection-provider.yaml
# echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/provider.mynetwork.com/connection-provider.json
# echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/provider.mynetwork.com/connection-provider.yaml

ORG=Subscriber
P0PORT=9051
P1PORT=9151
CAPORT=8054
PEERPEM=organizations/peerOrganizations/subscriber.mynetwork.com/tlsca/tlsca.subscriber.mynetwork.com-cert.pem
CAPEM=organizations/peerOrganizations/subscriber.mynetwork.com/ca/ca.subscriber.mynetwork.com-cert.pem

echo "$(json_ccp_2 $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/connection-subscriber.json
echo "$(yaml_ccp_2 $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/connection-subscriber.yaml
# echo "$(json_ccp_2 $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/subscriber.mynetwork.com/connection-subscriber.json
# echo "$(yaml_ccp_2 $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/subscriber.mynetwork.com/connection-subscriber.yaml

ORG=Regulator
P0PORT=10051
CAPORT=9054
PEERPEM=organizations/peerOrganizations/regulator.mynetwork.com/tlsca/tlsca.regulator.mynetwork.com-cert.pem
CAPEM=organizations/peerOrganizations/regulator.mynetwork.com/ca/ca.regulator.mynetwork.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/connection-regulator.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/connection-regulator.yaml
# echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/regulator.mynetwork.com/connection-regulator.json
# echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/regulator.mynetwork.com/connection-regulator.yaml
