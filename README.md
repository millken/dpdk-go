# Go bindings for [dpdk](http://dpdk.org/)

## Install dpdk and dpdk-go

- Setup `RTE_TARGET`, `RTE_SDK` and `DPDK_VERSION` in `hack/dpdk.rc`
- Install dpdk by `hack/dpdk-install.sh`
- Initial dpdk:
    - `source hack/dpdk-util.sh`
    - `init-dpdk`
- `go get -u github.com/feiskyer/dpdk-go`

Notes: If dpdk is installed inside a virtual machine, (e.g. VMWARE), then patch `hack/vmware.diff` must be applied before compling the dpdk source.

## Install dpdk-ovs

- First install dpdk
- Install ovs: `hack/ovs-install.sh`
- Start ovs with dpdk:
    - `source hack/dpdk-util.sh`
    - `init-dpdk`
    - `start-ovs`

## Sample Applications

### [Helloworld](http://dpdk.org/doc/guides/sample_app_ug/hello_world.html)

```
$ go get -u github.com/feiskyer/dpdk-go/samples/helloworld
$ helloworld -c3 -n1
```
