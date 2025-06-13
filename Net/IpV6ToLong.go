package Net

import (
	"math/big"
	"net"
)

func (conv *Net) IpV6ToLong(value string) *big.Int {
	parsedIP := net.ParseIP(value)
	if parsedIP == nil {
		return nil
	}

	ip := parsedIP.To16()
	if ip != nil {
		ipBytes := parsedIP.To16()
		result := new(big.Int)
		result.SetBytes(ipBytes)

		return result
	} else {
		return nil
	}
}
