#!/bin/bash

WORK_PATH=$PWD

export FABRIC_VERSION=2.3.0
export CA_VERSION=1.4.9
export DB_VERSION=3.1.1
DOCKER_NS=hyperledger
if [[ "$1" == "docker" ]]; then
  echo "Pulling Docker Images"
  # Fabric-CA image
  echo "Pulling ${DOCKER_NS}/fabric-ca:${CA_VERSION}"
  docker pull ${DOCKER_NS}/fabric-ca:${CA_VERSION}

  # Fabric images
  FABRIC_IMAGES=(fabric-peer fabric-orderer fabric-tools fabric-ccenv fabric-javaenv fabric-nodeenv fabric-baseos)
  for image in ${FABRIC_IMAGES[@]}; do
    echo "Pulling ${DOCKER_NS}/$image:${FABRIC_VERSION}"
    docker pull ${DOCKER_NS}/$image:${FABRIC_VERSION}
  done

  # Other images
  docker pull couchdb:${DB_VERSION}
else
  echo "ignored."
fi

ARCH=$(echo "$(uname -s|tr '[:upper:]' '[:lower:]'|sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')")

echo "Download Fabric Bianries"
cd ${WORK_PATH}/fabric-bin
FILE_NAME=hyperledger-fabric-${ARCH}-${FABRIC_VERSION}.tar.gz
if [ ! -f "${FILE_NAME}" ]; then
  echo "downloading fabric binaries (${FILE_NAME})..."
  wget https://github.com/hyperledger/fabric/releases/download/v${FABRIC_VERSION}/${FILE_NAME}
else
  echo "fabric binaries existing (${FABRIC_VERSION}), ignored"
fi

CA_FILE_NAME=hyperledger-fabric-ca-${ARCH}-${CA_VERSION}.tar.gz
if [ ! -f $CA_FILE_NAME ]; then
  echo "downloading fabric-ca binaries (${CA_FILE_NAME})..."
  wget https://github.com/hyperledger/fabric-ca/releases/download/v${CA_VERSION}/${CA_FILE_NAME}
else
  echo "fabric-ca binaries existing (${CA_VERSION}), ignored"
fi

if [ -d "${FABRIC_VERSION}" ]; then
  rm -rf ./${FABRIC_VERSION}
fi

mkdir -p ./${FABRIC_VERSION}
cd ./${FABRIC_VERSION}

tar zxf ../${FILE_NAME}
tar zxf ../${CA_FILE_NAME}


sudo cp ${WORK_PATH}/fabric-bin/${FABRIC_VERSION}/bin/* /usr/local/bin/
cd $WORK_PATH
echo