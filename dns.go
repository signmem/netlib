package netlib

import (
	"net"
	"net/url"
)

func DnsCheck(server string) (addrs []string, err error) {

	host, err := HostOnly(server)
        if err != nil {
               return addrs, err
        }

        if net.ParseIP(host) == nil {
                addrs, err := net.LookupHost(host)
                if err != nil {
			return addrs, err
                }
		return addrs, nil
        }
	addrs = append(addrs, host)
	return addrs, nil
}

func HostOnly(hoststr string) (host string, err error) {

        if host, _, err = net.SplitHostPort(hoststr); err != nil {

                if ip := net.ParseIP(hoststr); ip != nil {
                        return hoststr, nil
                }

                if err, ok := err.(*net.AddrError); ok  && err.Err == "missing port in address"  {
                        return hoststr, nil
                }
                return  "",  err
        }
        return host ,nil
}

func UrlDnsCheck(urlstring string) (addr []string, err error) {

	parsed, err := url.Parse(urlstring)
	if err != nil {
		return addr, err
	}

	domain := parsed.Hostname()
	addr, err = DnsCheck(domain)
	return addr, err
}
