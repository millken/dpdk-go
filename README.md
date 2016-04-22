# Go bindings for [dpdk](http://dpdk.org/)

## Install dpdk-go

- Install dpdk by `hack/dpdk-install.sh`
- `go get -u github.com/feiskyer/dpdk-go`

If dpdk is installed inside a virtual machine, (e.g. VMWARE), then patch `hack/vmware.diff` must be applied before compling the dpdk source.

## Sample Applications

### [Helloworld](http://dpdk.org/doc/guides/sample_app_ug/hello_world.html)

```
$ go get -u github.com/feiskyer/dpdk-go/samples/helloworld
$ helloworld -c3 -n1
```
