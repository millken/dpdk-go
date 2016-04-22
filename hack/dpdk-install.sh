#!/bin/bash
set -x
source dpdk.rc

echo "Install dependencies..."
lsb_dist=''
if command_exists lsb_release; then
	lsb_dist="$(lsb_release -si)"
fi
if [ -z "$lsb_dist" ] && [ -r /etc/lsb-release ]; then
	lsb_dist="$(. /etc/lsb-release && echo "$DISTRIB_ID")"
fi
if [ -z "$lsb_dist" ] && [ -r /etc/debian_version ]; then
	lsb_dist='debian'
fi
if [ -z "$lsb_dist" ] && [ -r /etc/fedora-release ]; then
	lsb_dist='fedora'
fi
if [ -z "$lsb_dist" ] && [ -r /etc/oracle-release ]; then
	lsb_dist='oracleserver'
fi
if [ -z "$lsb_dist" ]; then
	if [ -r /etc/centos-release ] || [ -r /etc/redhat-release ]; then
		lsb_dist='centos'
	fi
fi
if [ -z "$lsb_dist" ] && [ -r /etc/os-release ]; then
	lsb_dist="$(. /etc/os-release && echo "$ID")"
fi

lsb_dist="$(echo "$lsb_dist" | tr '[:upper:]' '[:lower:]')"
case "$lsb_dist" in
	ubuntu|debian)
	apt-get install -y vim gcc-multilib libfuse-dev linux-source libssl-dev llvm-dev python autoconf libtool libpciaccess-dev make libcunit1-dev libaio-dev
	;;

	fedora|centos|oraclelinux|redhat)
	yum -y install git cmake gcc autoconf automake device-mapper-devel \
	   sqlite-devel pcre-devel libsepol-devel libselinux-devel \
	   automake autoconf gcc make glibc-devel glibc-devel.i686 kernel-devel \
	   fuse-devel pciutils libtool openssl-devel libpciaccess-devel CUnit-devel libaio-devel
	mkdir -p /lib/modules/$(uname -r)
	ln -sf /usr/src/kernels/$(uname -r) /lib/modules/$(uname -r)/build
	;;

	*)
	echo "Distro $lsb_dist not supported now"
	exit 1
	;;
esac

# download dpdk
curl -sSL http://dpdk.org/browse/dpdk/snapshot/dpdk-${DPDK_VERSION}.tar.gz | tar -xz && mv dpdk-${DPDK_VERSION} ${RTE_SDK}

# install dpdk
cd ${RTE_SDK}
sed -ri 's,(CONFIG_RTE_BUILD_COMBINE_LIBS=).*,\1y,' config/common_linuxapp
sed -ri 's,(CONFIG_RTE_LIBRTE_VHOST=).*,\1y,' config/common_linuxapp
sed -ri 's,(CONFIG_RTE_LIBRTE_VHOST_USER=).*,\1n,' config/common_linuxapp
sed -ri '/CONFIG_RTE_LIBNAME/a CONFIG_RTE_BUILD_FPIC=y' config/common_linuxapp
sed -ri '/EXECENV_CFLAGS  = -pthread -fPIC/{s/$/\nelse ifeq ($(CONFIG_RTE_BUILD_FPIC),y)/;s/$/\nEXECENV_CFLAGS  = -pthread -fPIC/}' mk/exec-env/linuxapp/rte.vars.mk
make config CC=gcc T=$RTE_TARGET
make -j `nproc` RTE_KERNELDIR=/lib/modules/$(uname -r)/build
make install
depmod

# update grub (need reboot)
sed -ri 's,(GRUB_CMDLINE_LINUX_DEFAULT=).*,\1"quiet transparent_hugepage=never default_hugepagesz=2MB hugepagesz=2MB hugepages=512",' /etc/default/grub
update-grub
