#!/bin/bash
set -x
source hack/dpdk.rc
mkdir -p /usr/local/include/
tar zxf hack/${DPDK_VERSION}/dpdk-libs.tar.gz -C /usr/local/lib/
tar zxf hack/${DPDK_VERSION}/dpdk-include.tar.gz -C /usr/local/include/

