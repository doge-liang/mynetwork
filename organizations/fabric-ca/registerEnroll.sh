function createProviderOrg {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/provider.mynetwork.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/provider.mynetwork.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-Provider --tls.certfiles ${PWD}/organizations/fabric-ca/providerOrg/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-Provider.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-Provider.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-Provider.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-Provider.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-Provider --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/providerOrg/tls-cert.pem
  set +x

  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-Provider --id.name user1 --id.secret user1pw --id.type client --tls.certfiles ${PWD}/organizations/fabric-ca/providerOrg/tls-cert.pem
  set +x

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-Provider --id.name providerOrgadmin --id.secret providerOrgadminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/providerOrg/tls-cert.pem
  set +x

	mkdir -p organizations/peerOrganizations/provider.mynetwork.com/peers
  mkdir -p organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-Provider -M ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/msp --csr.hosts peer0.provider.mynetwork.com --tls.certfiles ${PWD}/organizations/fabric-ca/providerOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-Provider -M ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls --enrollment.profile tls --csr.hosts peer0.provider.mynetwork.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/providerOrg/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/server.key

  mkdir ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/msp/tlscacerts/ca.crt

  mkdir ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/tlsca/tlsca.provider.mynetwork.com-cert.pem

  mkdir ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/ca
  cp ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/peers/peer0.provider.mynetwork.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/ca/ca.provider.mynetwork.com-cert.pem

  mkdir -p organizations/peerOrganizations/provider.mynetwork.com/users
  mkdir -p organizations/peerOrganizations/provider.mynetwork.com/users/User1@provider.mynetwork.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://user1:user1pw@localhost:7054 --caname ca-Provider -M ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/users/User1@provider.mynetwork.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/providerOrg/tls-cert.pem
  set +x

  mkdir -p organizations/peerOrganizations/provider.mynetwork.com/users/Admin@provider.mynetwork.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://providerOrgadmin:providerOrgadminpw@localhost:7054 --caname ca-Provider -M ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/users/Admin@provider.mynetwork.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/providerOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/provider.mynetwork.com/users/Admin@provider.mynetwork.com/msp/config.yaml

}


function createSubscriberOrg {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/subscriber.mynetwork.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:8054 --caname ca-Subscriber --tls.certfiles ${PWD}/organizations/fabric-ca/subscriberOrg/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-Subscriber.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-Subscriber.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-Subscriber.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-Subscriber.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-Subscriber --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/subscriberOrg/tls-cert.pem
  set +x

  echo
	echo "Register peer1"
  echo
  set -x
	fabric-ca-client register --caname ca-Subscriber --id.name peer1 --id.secret peer1pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/subscriberOrg/tls-cert.pem
  set +x

  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-Subscriber --id.name user1 --id.secret user1pw --id.type client --tls.certfiles ${PWD}/organizations/fabric-ca/subscriberOrg/tls-cert.pem
  set +x

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-Subscriber --id.name subscriberOrgadmin --id.secret subscriberOrgadminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/subscriberOrg/tls-cert.pem
  set +x

	mkdir -p organizations/peerOrganizations/subscriber.mynetwork.com/peers
  mkdir -p organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com
  mkdir -p organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-Subscriber -M ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/msp --csr.hosts peer0.subscriber.mynetwork.com --tls.certfiles ${PWD}/organizations/fabric-ca/subscriberOrg/tls-cert.pem
  set +x
 
  echo
  echo "## Generate the peer1 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer1:peer1pw@localhost:8054 --caname ca-Subscriber -M ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/msp --csr.hosts peer1.subscriber.mynetwork.com --tls.certfiles ${PWD}/organizations/fabric-ca/subscriberOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/msp/config.yaml
  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-Subscriber -M ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls --enrollment.profile tls --csr.hosts peer0.subscriber.mynetwork.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/subscriberOrg/tls-cert.pem
  set +x

  echo
  echo "## Generate the peer1-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer1:peer1pw@localhost:8054 --caname ca-Subscriber -M ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/tls --enrollment.profile tls --csr.hosts peer1.subscriber.mynetwork.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/subscriberOrg/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/server.key

  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/tls/server.key

  mkdir ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/msp/tlscacerts/ca.crt
  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/msp/tlscacerts/ca.crt

  mkdir ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/tlsca/tlsca.subscriber.mynetwork.com-cert.pem
  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/tlsca/tlsca.subscriber.mynetwork.com-cert.pem

  mkdir ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/ca
  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer0.subscriber.mynetwork.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/ca/ca.subscriber.mynetwork.com-cert.pem
  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/peers/peer1.subscriber.mynetwork.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/ca/ca.subscriber.mynetwork.com-cert.pem

  mkdir -p organizations/peerOrganizations/subscriber.mynetwork.com/users
  mkdir -p organizations/peerOrganizations/subscriber.mynetwork.com/users/User1@subscriber.mynetwork.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://user1:user1pw@localhost:8054 --caname ca-Subscriber -M ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/users/User1@subscriber.mynetwork.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/subscriberOrg/tls-cert.pem
  set +x

  mkdir -p organizations/peerOrganizations/subscriber.mynetwork.com/users/Admin@subscriber.mynetwork.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://subscriberOrgadmin:subscriberOrgadminpw@localhost:8054 --caname ca-Subscriber -M ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/users/Admin@subscriber.mynetwork.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/subscriberOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/subscriber.mynetwork.com/users/Admin@subscriber.mynetwork.com/msp/config.yaml

}


function createRegulatorOrg {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/regulator.mynetwork.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:9054 --caname ca-Regulator --tls.certfiles ${PWD}/organizations/fabric-ca/regulatorOrg/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-Regulator.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-Regulator.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-Regulator.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-Regulator.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	fabric-ca-client register --caname ca-Regulator --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/regulatorOrg/tls-cert.pem
  set +x

  echo
  echo "Register user"
  echo
  set -x
  fabric-ca-client register --caname ca-Regulator --id.name user1 --id.secret user1pw --id.type client --tls.certfiles ${PWD}/organizations/fabric-ca/regulatorOrg/tls-cert.pem
  set +x

  echo
  echo "Register the org admin"
  echo
  set -x
  fabric-ca-client register --caname ca-Regulator --id.name regulatorOrgadmin --id.secret regulatorOrgadminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/regulatorOrg/tls-cert.pem
  set +x

	mkdir -p organizations/peerOrganizations/regulator.mynetwork.com/peers
  mkdir -p organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	fabric-ca-client enroll -u https://peer0:peer0pw@localhost:9054 --caname ca-Regulator -M ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/msp --csr.hosts peer0.regulator.mynetwork.com --tls.certfiles ${PWD}/organizations/fabric-ca/regulatorOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:9054 --caname ca-Regulator -M ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls --enrollment.profile tls --csr.hosts peer0.regulator.mynetwork.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/regulatorOrg/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/server.key

  mkdir ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/msp/tlscacerts/ca.crt

  mkdir ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/tlsca/tlsca.regulator.mynetwork.com-cert.pem

  mkdir ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/ca
  cp ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/peers/peer0.regulator.mynetwork.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/ca/ca.regulator.mynetwork.com-cert.pem

  mkdir -p organizations/peerOrganizations/regulator.mynetwork.com/users
  mkdir -p organizations/peerOrganizations/regulator.mynetwork.com/users/User1@regulator.mynetwork.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	fabric-ca-client enroll -u https://user1:user1pw@localhost:9054 --caname ca-Regulator -M ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/users/User1@regulator.mynetwork.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/regulatorOrg/tls-cert.pem
  set +x

  mkdir -p organizations/peerOrganizations/regulator.mynetwork.com/users/Admin@regulator.mynetwork.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://regulatorOrgadmin:regulatorOrgadminpw@localhost:9054 --caname ca-Regulator -M ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/users/Admin@regulator.mynetwork.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/regulatorOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/regulator.mynetwork.com/users/Admin@regulator.mynetwork.com/msp/config.yaml

}


function createOrderer {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/ordererOrganizations/mynetwork.com

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/ordererOrganizations/mynetwork.com
#  rm -rf $FABRIC_CA_CLIENT_HOME/fabric-ca-client-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:10054 --caname ca-Orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-Orderer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-Orderer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-Orderer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-10054-ca-Orderer.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/ordererOrganizations/mynetwork.com/msp/config.yaml


  echo
	echo "Register orderer"
  echo
  set -x
	fabric-ca-client register --caname ca-Orderer --id.name orderer --id.secret ordererpw --id.type orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
    set +x

  echo
  echo "Register the orderer admin"
  echo
  set -x
  fabric-ca-client register --caname ca-Orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  mkdir -p organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com

  echo
  echo "## Generate the orderer msp"
  echo
  set -x
	fabric-ca-client enroll -u https://orderer:ordererpw@localhost:10054 --caname ca-Orderer -M ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/msp --csr.hosts orderer.mynetwork.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/mynetwork.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/msp/config.yaml

  echo
  echo "## Generate the orderer-tls certificates"
  echo
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:10054 --caname ca-Orderer -M ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/tls --enrollment.profile tls --csr.hosts orderer.mynetwork.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/tls/ca.crt
  cp ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/tls/signcerts/* ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/tls/server.crt
  cp ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/tls/keystore/* ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/tls/server.key

  mkdir ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/msp/tlscacerts/tlsca.orderer.mynetwork.com-cert.pem

  mkdir ${PWD}/organizations/ordererOrganizations/mynetwork.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/mynetwork.com/orderers/orderer.mynetwork.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/mynetwork.com/msp/tlscacerts/tlsca.mynetwork.com-cert.pem

  mkdir -p organizations/ordererOrganizations/mynetwork.com/users
  mkdir -p organizations/ordererOrganizations/mynetwork.com/users/Admin@mynetwork.com

  echo
  echo "## Generate the admin msp"
  echo
  set -x
	fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:10054 --caname ca-Orderer -M ${PWD}/organizations/ordererOrganizations/mynetwork.com/users/Admin@mynetwork.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/mynetwork.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/mynetwork.com/users/Admin@mynetwork.com/msp/config.yaml


}
