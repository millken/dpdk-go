## Exception Path Sample Application

The Exception Path sample application is a simple example that demonstrates the use of the DPDK to set up an exception path for packets to go through the Linux* kernel. This is done by using virtual TAP network interfaces. These can be read from and written to by the DPDK application and appear to the kernel as a standard network interface.

The application creates two threads for each NIC port being used. One thread reads from the port and writes the data unmodified to a thread-specific TAP interface. The second thread reads from a TAP interface and writes the data unmodified to the NIC port.

![](http://dpdk.org/doc/guides/_images/exception_path_example.svg)

## Usage

Reads packets from `dpdk port 0`, and then write the data to `tap_dpdk_00`:

```sh
$ go get -u github.com/feiskyer/dpdk-go/samples/exception-path
$ exception-path
```

Open another termial, and run 

```sh
$ tcpdump -nn -i tap_dpdk_00
```

## c verison exception_path

```sh
# prepare bridge
apt-get install -y openvpn bridge-utils
openvpn --mktun --dev tap_dpdk_00
openvpn --mktun --dev tap_dpdk_01
ip link set tap_dpdk_00 up
ip link set tap_dpdk_01 up
brctl addbr br0
brctl addif br0 tap_dpdk_00
brctl addif br0 tap_dpdk_01
ifconfig br0 up

# start application
./exception_path -c3 -n2 -- -p1 -i1 -o2

# Getting Statistics
# Open Another termial, and run
$ killall -USR1 exception_path
```

**Clearup**

```sh
ifconfig br0 down
brctl delbr br0
openvpn --rmtun --dev tap_dpdk_00
openvpn --rmtun --dev tap_dpdk_01
```