#!/bin/bash
echo "Creating channel..."
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/vehicle.com/orderers/orderer.vehicle.com/msp/tlscacerts/tlsca.vehicle.com-cert.pem
CORE_PEER_LOCALMSPID=CitizenMSP
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/citizen.vehicle.com/peers/peer0.citizen.vehicle.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/citizen.vehicle.com/users/Admin@citizen.vehicle.com/msp
CORE_PEER_ADDRESS=peer0.citizen.vehicle.com:7051
CHANNEL_NAME=mainchannel
CORE_PEER_TLS_ENABLED=true
ORDERER_SYSCHAN_ID=syschain

peer channel create -o orderer.vehicle.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt

cat log.txt

#peer channel create -o orderer.vehicle.com:7050 /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/vehicle.com/orderers/orderer.vehicle.com/msp/tlscacerts/tlsca.vehicle.com-cert.pem -c mainchannel -f ./channel-artifacts/channel.tx

echo
echo "Channel created, joining Citizen..."
peer channel join -b mainchannel.block
