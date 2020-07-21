cd ..
export IMAGE_TAG=1.4

docker-compose -f docker-compose-cli.yaml up -d

docker exec -it cli bash ./scripts/channel/create-channel.sh

docker exec -it cli bash ./scripts/channel/join-peer.sh peer0 rto RTOMSP 8051 1.0
docker exec -it cli bash ./scripts/channel/join-peer.sh peer0 pollution PollutionMSP 9051 1.0
docker exec -it cli bash ./scripts/channel/join-peer.sh peer0 insurance InsuranceMSP 10051 1.0
docker exec -it cli bash ./scripts/channel/join-peer.sh peer0 police PoliceMSP 11051 1.0

CC_NAMES="profilemanager_cc rtoservices_cc vehicle_cc pollution_cc insurance_cc trafficviolation_cc"

for CC in $CC_NAMES; do
    echo "Installing "$CC
    docker exec -it cli bash ./scripts/install-cc/install-onpeer-cc.sh $CC peer0 citizen CitizenMSP 7051 1.0
    docker exec -it cli bash ./scripts/install-cc/install-onpeer-cc.sh $CC peer0 rto RTOMSP 8051 1.0
    docker exec -it cli bash ./scripts/install-cc/install-onpeer-cc.sh $CC peer0 pollution PollutionMSP 9051 1.0
    docker exec -it cli bash ./scripts/install-cc/install-onpeer-cc.sh $CC peer0 insurance InsuranceMSP 10051 1.0
    docker exec -it cli bash ./scripts/install-cc/install-onpeer-cc.sh $CC peer0 police PoliceMSP 11051 1.0
    echo "Instantiating "$CC
    docker exec -it cli bash ./scripts/install-cc/instantiate.sh $CC
done

echo "All Done!"
