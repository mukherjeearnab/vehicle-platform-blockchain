#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${ORGMSP}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${CAPORT}/$4/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        ../../connections/ccp-template.json 
}

ORG=citizen
ORGMSP=Citizen
P0PORT=7051
CAPORT=7054
PEERPEM=../crypto-config/peerOrganizations/citizen.vtan.com/tlsca/tlsca.citizen.vtan.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/citizen.vtan.com/ca/ca.citizen.vtan.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-citizen.json

ORG=rto
ORGMSP=RTO
P0PORT=8051
CAPORT=8054
PEERPEM=../crypto-config/peerOrganizations/rto.vtan.com/tlsca/tlsca.rto.vtan.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/rto.vtan.com/ca/ca.rto.vtan.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-rto.json
ORG=pollution
ORGMSP=Pollution
P0PORT=9051
CAPORT=9054
PEERPEM=../crypto-config/peerOrganizations/pollution.vtan.com/tlsca/tlsca.pollution.vtan.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/pollution.vtan.com/ca/ca.pollution.vtan.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-pollution.json

ORG=insurance
ORGMSP=Insurance
P0PORT=10051
CAPORT=10054
PEERPEM=../crypto-config/peerOrganizations/insurance.vtan.com/tlsca/tlsca.insurance.vtan.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/insurance.vtan.com/ca/ca.insurance.vtan.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-insurance.json

ORG=police
ORGMSP=Police
P0PORT=11051
CAPORT=11054
PEERPEM=../crypto-config/peerOrganizations/police.vtan.com/tlsca/tlsca.police.vtan.com-cert.pem
CAPEM=../crypto-config/peerOrganizations/police.vtan.com/ca/ca.police.vtan.com-cert.pem

echo "$(json_ccp $ORG $ORGMSP $P0PORT $CAPORT $PEERPEM $CAPEM)" >../../connections/connection-police.json
