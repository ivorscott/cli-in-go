#! /bin/bash

FHASH=`md5 $1`
while true; do
    NHASH=`md5 $1`
    if [ "NHASH" != "$FHASH" ]; then
        ./mdp -file $1
        FHASH=$NHASH
    fi
    sleep 5
done