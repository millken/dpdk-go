#!/bin/bash
source hack/dpdk.rc

# download ovs
if [ "${OVS_VERSION}" != "master" ]; then
	curl -sSl http://openvswitch.org/releases/openvswitch-${OVS_VERSION}.tar.gz | tar -xz && mv openvswitch-${OVS_VERSION} ${OVS_DIR}
else
	wget https://github.com/openvswitch/ovs/archive/master.zip
	unzip master.zip
	mv ovs-master ${OVS_DIR}
fi

# Build ovs
cd ${OVS_DIR}
./boot.sh
./configure --with-dpdk=/usr/local/share/dpdk/${RTE_TARGET} --prefix=${OVS_INSTALL_DIR} --localstatedir=/var --enable-ssl --with-linux=/lib/modules/$(uname -r)/build
make -j `nproc`
make install
