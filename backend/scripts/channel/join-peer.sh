#!/bin/bash
PEER=$1
ORG=$2
MSP=$3
PORT=$4
VERSION=$5

ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/vtan.com/orderers/orderer.vtan.com/msp/tlscacerts/tlsca.vtan.com-cert.pem
CORE_PEER_LOCALMSPID=$MSP
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG.vtan.com/peers/$PEER.$ORG.vtan.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG.vtan.com/users/Admin@$ORG.vtan.com/msp
CORE_PEER_ADDRESS=$PEER.$ORG.vtan.com:$PORT
CHANNEL_NAME=mainchannel
CORE_PEER_TLS_ENABLED=true

sleep 10
peer channel join -b mainchannel.block >&log.txt

cat log.txt
