#!/bin/bash

WORKSPACE=$(pwd)
if [ "$1" != "" ]; then
    WORKSPACE="${WORKSPACE}/$1"
fi

FUNC2TEST=$2

GOTESTCMD="go test -v ./..."
if [ "${FUNC2TEST}" != "" ]; then
    # test single function
    GOTESTCMD="${GOTESTCMD} -test.run ${FUNC2TEST}"
fi

echo "Running go test under ${WORKSPACE}"
echo "Command: ${GOTESTCMD}"

docker run --rm -it -v $(pwd):$(pwd) -w "${WORKSPACE}" golanger:v1.0.0 bash -c "${GOTESTCMD}"