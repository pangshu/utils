package Net

import (
	"net"
)

func (conv *Net) IsIpV4(value string) bool {
	ip := net.ParseIP(value)
	if ip == nil {
		return false
	}

	if ip.To4() != nil {
		return true
	} else {
		return false
	}
}
