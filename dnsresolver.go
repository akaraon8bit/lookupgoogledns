package lookupgoogledns

import (
    "context"
    "net"
    "time"
    "strings"
)

func GoogleResolver() (*net.Resolver) {
    r := &net.Resolver{
        PreferGo: true,
        Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
            d := net.Dialer{
                Timeout: time.Millisecond * time.Duration(10000),
            }
            return d.DialContext(ctx, network, "8.8.8.8:53")
        },
    }
    return  r
}



  func LookupAddr( ip string) (string ){
    r :=  GoogleResolver()
    hostnames, _ := r.LookupAddr(context.Background(),ip)
    hostname := ""
    if hostnames != nil {
        hostname = strings.TrimRight(string(hostnames[0]), ".")
    }
    return hostname
  }