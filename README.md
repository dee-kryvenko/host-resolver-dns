# host-resolver-dns

This is a simple DNS server implementation that uses host system resolver.

A typical use case for this is making VMs to understand host VPN with Split DNS. Typical NAT-ish type of VM networking allows network traffic, but modern VPN clients are also setting up conditional DNS resolvers (Split DNS) to enable host resolution from VPN for certain private DNS zones. Typical guest OS in a VM would not be aware of it, so private zones are not resolvable.

This is not a standalone project - actual DNS server implementation is used from https://github.com/lima-vm/lima (see https://github.com/lima-vm/lima/blob/master/pkg/hostagent/dns/dns.go). Unfortunately, Lima VM does not make it a distributable standalone binary to be used elsewhere, so this repository basically creates a CLI interface to run Lima VM DNS server implementation without Lima VM.

## Install

```bash
go install github.com/dee-kryvenko/host-resolver-dns@latest
```
