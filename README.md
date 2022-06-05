# host-resolver-dns

This is a simple DNS server implementation that uses host system resolver.

A typical use case for this is making VMs to understand host VPN with Split DNS. Typical NAT-ish type of VM networking allows network traffic, but modern VPN clients are also setting up conditional DNS resolvers (Split DNS) to enable host resolution from VPN for certain private DNS zones. Typical guest OS in a VM would not be aware of it, so private zones are not resolvable.

This is not a standalone project - actual DNS server implementation is used from https://github.com/lima-vm/lima (see https://github.com/lima-vm/lima/blob/master/pkg/hostagent/dns/dns.go). Unfortunately, Lima VM does not make it a distributable standalone binary to be used elsewhere, so this repository basically creates a CLI interface to run Lima VM DNS server implementation without Lima VM.

## Install

```bash
go install github.com/dee-kryvenko/host-resolver-dns@latest
```

## Use

By default, it is started as a desktop app. It uses system tray icon and can be stopped from there. For GUI elements to work, it should be installed with CGO: `CGO_ENABLED=1 go install github.com/dee-kryvenko/host-resolver-dns@latest`.

To use it in headless mode - add `--headless` flag. No CGO required to use it in headless mode.

### QEMU

In order to use a DNS server on a typical port 53 - set up port forwarding with iptables:

```bash
sudo iptables -t nat -A PREROUTING -d "10.0.2.2" -p udp --dport 53 -j DNAT --to-destination "10.0.2.2:16237"
sudo iptables -t nat -A OUTPUT -d "10.0.2.2" -p udp --dport 53 -j DNAT --to-destination "10.0.2.2:16237"
```
