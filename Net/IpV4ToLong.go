package Net

import (
	"net"
)

func (conv *Net) IpV4ToLong(value string) uint32 {
	parsedIP := net.ParseIP(value)
	if parsedIP == nil {
		return 0
	}

	ip := parsedIP.To4()
	if ip != nil {
		var result uint32
		for i := 0; i < 4; i++ {
			result = result<<8 + uint32(ip[i])
		}
		return result
	} else {
		return 0
	}
}
