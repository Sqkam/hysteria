package udphop

import (
	"net"
	"strings"
)

// UDPHopAddr contains an IP address and a list of ports.
type UDPHopAddrs struct {
	IPs     []net.IP
	Ports   []uint16
	PortStr string
}

func (a *UDPHopAddrs) Network() string {
	return "udphopx"
}

func (a *UDPHopAddrs) String() string {
	var ips []string
	for _, v := range a.IPs {
		ips = append(ips, v.String())
	}
	return net.JoinHostPort(strings.Join(ips, ","), a.PortStr)
}
