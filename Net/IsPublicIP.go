package Net

import "net"

// IsPublicIP 是否公网IP.
func (*Net) IsPublicIP(value string) bool {
	ip := net.ParseIP(value)
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsInterfaceLocalMulticast() {
		return false
	}
	if ipV4 := ip.To4(); ipV4 != nil {
		switch true {
		case ipV4[0] == 10:
			return false
		case ipV4[0] == 172 && ipV4[1] >= 16 && ipV4[1] <= 31:
			return false
		case ipV4[0] == 192 && ipV4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}
