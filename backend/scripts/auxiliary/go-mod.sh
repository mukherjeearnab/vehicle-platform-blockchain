cd ../../chaincode

CC_DIR=$PWD

CC_NAMES="profilemanager_cc rtoservices_cc vehicle_cc pollution_cc insurance_cc trafficviolation_cc"

for CC in $CC_NAMES; do
    echo "Installing Go dependencies in "$CC
    cd $CC
    go mod vendor
    cd ..
done
echo "Installing Go dependencies complete!"
