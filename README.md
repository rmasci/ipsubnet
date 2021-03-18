# IPv4 Subnet Calculator (GO)

Network calculator for subnet mask and other classless (CIDR) network information.

[![Build Status](https://travis-ci.org/brotherpowers/ipsubnet.svg?branch=master)](https://travis-ci.org/brotherpowers/ipsubnet)
[![GoDoc](https://godoc.org/github.com/brotherpowers/ipsubnet?status.svg)](https://github.com/brotherpowers/ipsubnet)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/brotherpowers/ipsubnet)

# This is a fork
This is a fork from [https://github.com/brotherpowers/ipsubnet](https://www.github.com/brotherpowers/ipsubnet)

I would do a pull request but that breaks the API. I converted all int to int64 so that it would compile on 32 bit system. Calls to ipsubnet need to be int64 and not int, and would break code for people using this. So I kept it separate, but wanted to give credit to the person that made it. 

I've not done extensive tests or even changed ip_test.go -- it works for what I need it to do,  and I have to move on with the project I had to make this change for instead of further development. 

### Calculations
 * IP address network subnet masks, network and host portions, and provides aggregated reports.
 * Subnet mask
 * Network portion
 * Host portion
 * Number of IP addresses in the network
 * Number of addressable hosts in the network
 * IP address range
 * Broadcast address

Provides each data in dotted quads, hexadecimal, and binary formats, as well as array of quads.

## Usages

### Install

```go
go get -u github.com/rmasci/ipsubnet
```

### Get Started

```go
import "github.com/rmasci/ipsubnet"

sub := ipsubnet.SubnetCalculator("192.168.112.203", 23)
```


### Get Network Information

```go
sub.GetNumberIPAddresses();      // 512
sub.GetNumberAddressableHosts(); // 510
sub.GetIPAddressRange();         // [ 192.168.112.0, 192.168.113.255 ]
sub.GetNetworkSize();            // 23
sub.GetBroadcastAddress();       // 192.168.113.255
```

### Get IP Address
```go
sub.GetIPAddress();       // 192.168.112.203
sub.GetIPAddressQuads();  // [ 192, 168, 112, 203 ]
sub.GetIPAddressHex();    // C0A870CB
sub.GetIPAddressBinary(); // 11000000101010000111000011001011
```

### Get Subnet Mask
```go
sub.GetSubnetMask();       // 255.255.254.0
sub.GetSubnetMaskQuads();  // [ 255, 255, 254, 0 ]
sub.GetSubnetMaskHex();    // FFFFFE00
sub.GetSubnetMaskBinary(); // 11111111111111111111111000000000
```

### Get Network Portion
```go
sub.GetNetworkPortion();       // 192.168.112.0
sub.GetNetworkPortionQuads();  // [ 192, 168, 112, 0 ]
sub.GetNetworkPortionHex();    // C0A87000
sub.GetNetworkPortionBinary(); // 11000000101010000111000000000000
```

### Get Host Portion
```go
sub.GetHostPortion();       // 0.0.0.203
sub.GetHostPortionQuads();  // [ 0, 0, 0, 203 ]
sub.GetHostPortionHex();    // 000000CB
sub.GetHostPortionBinary(); // 00000000000000000000000011001011
```


## Based On

[IPv4 Subnet Calculator ( PHP )](https://github.com/markrogoyski/ipv4-subnet-calculator-php)

## License

IPv4 Subnet Calculator (GO) is licensed under the MIT License.
