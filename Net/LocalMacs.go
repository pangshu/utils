package Net

import (
	"net"
)

// LocalMacs 获取本机MAC地址
func (*Net) LocalMacs() []string {
	result := make([]string, 0)
	addr, err := net.Interfaces()
	if err != nil {
		return nil
	} else {
		for _, inf := range addr {
			mac := inf.HardwareAddr
			if mac != nil {
				result = append(result, mac.String())
			}
		}
	}
	return result
}
