#!/bin/bash

export GOPATH="$GOPATH:`pwd`"
echo "Building toukei" 
pushd bin && go build toukei && popd
