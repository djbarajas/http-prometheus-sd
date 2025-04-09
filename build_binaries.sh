#!/bin/bash

pushd go_service
make binary
popd

pushd http_service_discovery
make binary
