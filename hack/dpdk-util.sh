#!/bin/bash
source hack/dpdk.rc

function stop() {
    # stop everything
    ovs-vsctl del-br ovsbr
    # revert to using kernel driver
    ${RTE_SDK}/tools/dpdk_nic_bind.py -b ixgbe $IFPORT
    ${RTE_SDK}/tools/dpdk_nic_bind.py --status
    pgrep -f openvswitch | xargs kill -9
    /bin/rm -f /var/run/openvswitch/*.pid
}

function init-dpdk() {
    # init hugepages
    # mount -t hugetlbfs -o pagesize=1G none /dev/hugepages
    # echo 512 > /sys/devices/system/node/node0/hugepages/hugepages-2048kB/nr_hugepages
    sysctl vm.nr_hugepages=512
    modprobe uio
    modprobe igb_uio
    
    # bind IFNAME
    ifconfig ${IFNAME} 0
    ${RTE_SDK}/tools/dpdk_nic_bind.py --bind=igb_uio $IFPORT
    ${RTE_SDK}/tools/dpdk_nic_bind.py --status
}

function helloworld() {
    # helloworld demo
    cd ${RTE_SDK}/examples/helloworld
    make RTE_SDK=${RTE_SDK} RTE_TARGET=${RTE_TARGET}
    ./build/helloworld -c 0x03 -n 1
}

function l2fw() {
    # L2 Forward demo
    cd ${RTE_SDK}/examples/l2fwd
    make RTE_SDK=${RTE_SDK} RTE_TARGET=${RTE_TARGET}    
    ./build/l2fwd -c3 -n1 -- -p 0x1 -T 1
}

function get-cpu-layout() {
    ${RTE_SDK}/tools/cpu_layout.py
}

function start-testpmd() {
    # set port-topology=chained for odd number of ports
    ${RTE_SDK}build/app/testpmd -c3 -n1 -- -i --portmask=0x1 -a --port-topology=chained
}

function start-ovs() {
    # start ovs
    mkdir -p /etc/openvswitch /var/log/openvswitch /var/run/openvswitch/
    ovsdb-tool create /etc/openvswitch/conf.db /usr/share/openvswitch/vswitch.ovsschema
    ovs-vsctl --no-wait init

    ovsdb-server /etc/openvswitch/conf.db -vconsole:emer -vsyslog:err -vfile:info --remote=punix:/var/run/openvswitch/db.sock --remote=db:Open_vSwitch,Open_vSwitch,manager_options --private-key=db:Open_vSwitch,SSL,private_key --certificate=db:Open_vSwitch,SSL,certificate --bootstrap-ca-cert=db:Open_vSwitch,SSL,ca_cert --no-chdir --log-file=/var/log/openvswitch/ovsdb-server.log --pidfile=/var/run/openvswitch/ovsdb-server.pid --detach --monitor

    ovs-vswitchd --dpdk -c 0x3 -n 4 --socket-mem 400,400 --proc-type primary -- unix:/var/run/openvswitch/db.sock --log-file=/var/log/openvswitch/ovs-vswitchd.log --pidfile=/var/run/openvswitch/ovs-vswitchd.pid --detach --monitor
}

function setup-bridge() {
    # setup ovs bridge
    ovs-vsctl --no-wait --may-exist add-br ovsbr -- set Bridge ovsbr datapath_type=netdev
    ip link set ovsbr up
    ovs-vsctl --no-wait --may-exist add-port ovsbr dpdk0 -- set Interface dpdk0 type=dpdk
    ovs-vsctl show
}
