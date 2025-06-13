package Net

import (
	"net"
)

// LocalIps 提取本机所有IP
func (*Net) LocalIps() []string {
	result := make([]string, 0)
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	} else {
		for _, addr := range addr {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if nil != ipNet.IP.To4() {
					result = append(result, ipNet.IP.String())
				}
			}
		}
	}
	return result
}
