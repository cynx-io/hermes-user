package helper

import (
	"net"
	"net/http"
	"strings"
)

func GetClientIP(r *http.Request) string {
	// If the request is behind a reverse proxy, the IP address might be forwarded in the X-Forwarded-For header.
	// First, check for the X-Forwarded-For header.
	ips := r.Header.Get("X-Forwarded-For")
	if ips != "" {
		// The X-Forwarded-For header contains a comma-separated list of IPs
		// The first IP in the list is the original client IP.
		return strings.Split(ips, ",")[0]
	}

	// Otherwise, fallback to the remote address.
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
