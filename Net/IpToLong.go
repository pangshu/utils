package Net

import (
	"math/big"
	"net"
	"utils/Var"
)

func (conv *Net) IpToLong(value string) *Var.Var {
	parsedIP := net.ParseIP(value)
	if parsedIP == nil {
		return nil
	}

	ipV4 := parsedIP.To4()
	if ipV4 != nil {
		var result uint32
		for i := 0; i < 4; i++ {
			result = result<<8 + uint32(ipV4[i])
		}
		return Var.New(result)
	}

	ip := parsedIP.To16()
	if ip != nil {
		ipBytes := parsedIP.To16()
		result := new(big.Int)
		result.SetBytes(ipBytes)

		return Var.New(result.String())
	} else {
		return nil
	}

}
