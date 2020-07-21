#!/bin/bash
PEER=$1
ORG=$2
MSP=$3
PORT=$4
VERSION=$5

ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/vehicle.com/orderers/orderer.vehicle.com/msp/tlscacerts/tlsca.vehicle.com-cert.pem
CORE_PEER_LOCALMSPID=$MSP
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG.vehicle.com/peers/$PEER.$ORG.vehicle.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG.vehicle.com/users/Admin@$ORG.vehicle.com/msp
CORE_PEER_ADDRESS=$PEER.$ORG.vehicle.com:$PORT
CHANNEL_NAME=mainchannel
CORE_PEER_TLS_ENABLED=true

peer channel join -b mainchannel.block >&log.txt

cat log.txt
