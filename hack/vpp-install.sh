#!/bin/bash
# Install vpp from source
# TIP: vpp will install dpdk-16.04 automatically
# TIP: uninstall with cmd:
#   dpkg --purge vpp vpp-dbg vpp-dev vpp-dpdk-dev vpp-dpdk-dkms vpp-lib
git clone https://gerrit.fd.io/r/vpp /usr/src/vpp
cd /usr/src/vpp/
# make install-dep
# make build-release
cd build-root
make distclean
./bootstrap.sh
# make V=0 PLATFORM=vpp TAG=vpp install-rpm
# rpm -i *.rpm
make V=0 PLATFORM=vpp TAG=vpp install-deb
dpkg -i *.deb
systemctl start vpp
