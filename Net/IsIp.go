package Net

import (
	"net"
)

func (conv *Net) IsIp(value string) bool {
	ip := net.ParseIP(value)
	if ip == nil {
		return false
	} else {
		return true
	}
}
