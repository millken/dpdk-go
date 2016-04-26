#!/bin/bash
set -x
source hack/dpdk.rc
mkdir -p /usr/local/include/
tar zxf hack/2.2.0/dpdk-libs.tar.gz -C /usr/local/lib/
tar zxf hack/2.2.0/dpdk-include.tar.gz -C /usr/local/include/

