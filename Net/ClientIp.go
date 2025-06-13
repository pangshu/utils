package Net

import (
	"net"
	"net/http"
	"strings"
)

// ClientIp 返回远程客户端的 IP，如 192.168.1.1
func (*Net) ClientIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		host, port, err := net.SplitHostPort(strings.TrimSpace(remoteAddr))
		if err == nil && host != "" && port != "" {
			remoteAddr = host
		} else {
			remoteAddr = ""
		}
	}

	// 统一本地地址表示
	if remoteAddr == "::1" || remoteAddr == "0:0:0:0:0:0:0:1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}
